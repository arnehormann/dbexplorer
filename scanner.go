package dbexplorer

import (
	"database/sql"
	"fmt"
	"github.com/arnehormann/sqlinternals/mysqlinternals"
	"reflect"
	"strings"
)

type structFeeder func(t []interface{}) interface{}

type request struct {
	parentIdx    int
	responseType reflect.Type
	feeder       structFeeder
	name         []string // schema, {table, column / trigger / proc / func}
	query        string
	queryArgs    []interface{}
}

func (r *request) path() string {
	if r == nil {
		return ""
	}
	return r.responseType.String() + ":" + strings.Join(r.name, "/")
}

type response struct {
	request
	columns []mysqlinternals.Column
	values  []interface{}
}

func (r *request) scan(db *sql.DB) (*response, error) {
	if db == nil || r == nil {
		return nil, nil
	}
	stmt, err := db.Prepare(r.query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(r.queryArgs...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	columns, err := mysqlinternals.Columns(rows)
	if err != nil {
		return nil, err
	}
	numCols := len(columns)
	if r.responseType != nil {
		numFields := r.responseType.NumField()
		if numCols != numFields {
			return nil, fmt.Errorf("[%v] cannot be written to [%v]", columns, r.responseType)
		}
	}
	receivers := make([]interface{}, numCols*2)
	for i := 0; i < numCols; i++ {
		receivers[i] = &receivers[numCols+i]
	}
	values := []interface{}{}
	for rows.Next() {
		data := r.feeder(receivers)
		err = rows.Scan(receivers[:numCols]...)
		if err != nil {
			return nil, err
		}
		values = append(values, data)
	}
	return &response{
		request: *r,
		columns: columns,
		values:  values,
	}, nil
}

type requestError struct {
	req request
	err error
}

func (r *requestError) Error() string {
	return r.err.Error()
}
