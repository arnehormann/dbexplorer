package generator

import (
	"database/sql"
	"fmt"
	"io"
	"text/template"
)

// store names of generated structs (Arg.Name, Arg.ValueType)
type marker map[string]struct{}

var structTmpl = template.Must(template.New("mysqlStruct").Delims("~", "~").Parse(structTemplate))

func (mark *marker) Generate(writer io.Writer, db *sql.DB, query Query) error {
	var exists struct{}
	name, ok := query.StructName()
	if !ok {
		return fmt.Errorf("No struct name in query %#v", query)
	}
	if _, known := (*mark)[name]; known {
		return fmt.Errorf("Type %s was already declared", name)
	}
	(*mark)[name] = exists
	sql, err := query.Sample()
	if err != nil {
		return err
	}
	req := &request{
		feeder: func(t []interface{}) interface{} { return nil },
		query:  sql,
	}
	resp, err := req.scan(db)
	if err != nil {
		return err
	}
	overrides := query.Overrides(len(resp.columns))
	cols := make([]TemplateColumn, len(resp.columns))
	for i, c := range resp.columns {
		cols[i] = TemplateColumn{Column: c}
		cols[i].nameOverride = overrides[i]
	}
	return structTmpl.Execute(writer, struct {
		Name  string
		Query Query
		Cols  []TemplateColumn
	}{
		Name:  name,
		Query: query,
		Cols:  cols,
	})
}

func Generate(writer io.Writer, db *sql.DB, pkg string, queries ...Query) error {
	mark := make(marker)
	_, err := fmt.Fprintf(writer, baseTemplate, pkg)
	if err != nil {
		return err
	}
	for _, query := range queries {
		// consider arguments
		err = (&mark).Generate(writer, db, query)
		if err != nil {
			return err
		}
	}
	return nil
}
