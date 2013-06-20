package dbexplorer

import (
	"bytes"
	"database/sql"
	"fmt"
	"github.com/arnehormann/sqlinternals/mysqlinternals"
	"io"
	"text/template"
)

func Generate(writer io.Writer, db *sql.DB, name, query string, args ...interface{}) error {
	type column struct {
		Name    string
		Type    string
		Tags    string
		Comment string
	}
	const structTemplate = `
type ~.Name~ struct {~range $i, $e := .Cols~
	~$e.Name~ ~$e.Type~~if $e.Tags~ ~$e.Tags~~end~ // ~$e.Comment~~end~
}

func (q *~.Name~) ScanInto(t []interface{}) interface{} {
	v := ~.Name~{}~range $i, $e := .Cols~
	t[~$i~] = &v.~$e.Name~~end~
	return &v
}

func (q *~.Name~) Query() (string, []string) {
	return ~.Query | printf "%q"~, []string{~range $i, $e := .QueryArgs~~$e | printf "%q,"~~end~}
}

`
	req := &request{
		feeder:    func(t []interface{}) interface{} { return nil },
		query:     query,
		queryArgs: args,
	}
	resp, err := req.scan(db)
	if err != nil {
		return err
	}
	templ, err := template.New(name).Delims("~", "~").Parse(structTemplate)
	if err != nil {
		return err
	}
	cols := make([]column, len(resp.columns))
	queryArgsFmt := ""
	for i, c := range resp.columns {
		goType, err := c.ReflectGoType()
		if err != nil {
			return err
		}
		var sqlType string
		if c.MysqlParameters() == mysqlinternals.ParamMustLength {
			sqlType, err = c.MysqlDeclaration(255) // should be varchar mostly
		} else {
			sqlType, err = c.MysqlDeclaration()
		}
		if err != nil {
			return err
		}
		cols[i] = column{
			Name:    c.Name(),
			Type:    goType.String(),
			Comment: sqlType,
		}
		if i == 0 {
			queryArgsFmt += "%q"
		} else {
			queryArgsFmt += ",%q"
		}
	}
	return templ.Execute(writer, struct {
		Name      string
		Query     string
		QueryArgs []interface{}
		Cols      []column
	}{
		Name:      name,
		Query:     query,
		QueryArgs: args,
		Cols:      cols,
	})
}

func GenerateAll(writer io.Writer, db *sql.DB, pkg string, queries, args map[string]string) error {
	const baseTemplate = `// this file is generated, do not edit it!

package %s

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"time"
)

// force import even if unused to avoid writing detection code
var (
	_ = time.Monday
	_ = mysql.NullTime{}
	_ = sql.NullString{}
)
`
	_, err := fmt.Fprintf(writer, baseTemplate, pkg)
	if err != nil {
		return err
	}
	buffer := bytes.NewBuffer(nil)
	for name, query := range queries {
		buffer.Reset()
		tmpl, err := template.New(name).Parse(query)
		if err != nil {
			return err
		}
		err = tmpl.Execute(buffer, args)
		if err != nil {
			return err
		}
		// consider arguments
		err = Generate(writer, db, name, buffer.String())
		if err != nil {
			return err
		}
	}
	return nil
}
