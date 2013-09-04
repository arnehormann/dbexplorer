package generator

import (
	"database/sql"
	"fmt"
	"io"
	"text/template"
)

// store names of generated structs (Arg.Name, Arg.ValueType)
type marker map[string]struct{}

func (mark *marker) Generate(writer io.Writer, db *sql.DB, query Query) error {
	name, ok := query.StructName()
	if !ok {
		return fmt.Errorf("No struct name in query %#v", query)
	}
	_, known := (*mark)[name]
	if known {
		return fmt.Errorf("Type %s was already declared", name)
	}
	(*mark)[name] = struct{}{}
	templ, err := template.New(name).Delims("~", "~").Parse(structTemplate)
	if err != nil {
		return err
	}
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
	overrides := query.Overrides()
	cols := make([]TemplateColumn, len(resp.columns))
	for i, c := range resp.columns {
		// TODO: consider overrides here
		cols[i] = TemplateColumn{Column: c}
		if v, ok := overrides[i]; ok {
			cols[i].nameOverride = v
		}
	}
	return templ.Execute(writer, struct {
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
