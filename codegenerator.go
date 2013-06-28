package dbexplorer

import (
	"database/sql"
	"fmt"
	"github.com/arnehormann/sqlinternals/mysqlinternals"
	"io"
	"reflect"
	"regexp"
	"strings"
	"text/template"
)

type Query []interface{}

func SQL(parts ...interface{}) Query {
	return Query(parts)
}

type Column struct {
	mysqlinternals.Column
}

type Arg struct {
	// argument name
	Name string
	// argument value for sample query
	Sample string
	// argument to prepared statement (true) or static identifier (false)?
	Dynamic bool
	// argument should be quoted
	Quoted bool
	// argument type (string by default)
	Type reflect.Type
	// alternatives
	ValueType string
	Values    map[string]interface{}
	// description
	Comment string
}

type Arglist []*Arg

func (q Query) Sample() (string, error) {
	sql := ""
	for i, part := range q {
		switch v := part.(type) {
		case string:
			sql += v
		case *Arg:
			if v.Sample == "" {
				return sql, fmt.Errorf("argument %d has no Sample set", i)
			}
			if v.Quoted {
				sql += fmt.Sprintf("%q", v.Sample)
			} else {
				sql += v.Sample
			}
		default:
			return sql, fmt.Errorf("'%v' is not string or *Arg", part)
		}
	}
	return sql, nil
}

func (q Query) Args() Arglist {
	args := Arglist{}
	for _, part := range q {
		switch v := part.(type) {
		case *Arg:
			args = append(args, v)
		}
	}
	return args
}

func (a Arglist) Dynamic() Arglist {
	args := Arglist{}
	for _, arg := range a {
		if arg.Dynamic {
			args = append(args, arg)
		}
	}
	return args
}

func (a Arglist) Static() Arglist {
	args := Arglist{}
	for _, arg := range a {
		if !arg.Dynamic {
			args = append(args, arg)
		}
	}
	return args
}

func (a Arglist) Declaration() string {
	if len(a) == 0 {
		return ""
	}
	var stringType = reflect.TypeOf("")
	var lastType reflect.Type = nil
	list := ""
	for _, arg := range a {
		switch {
		case arg.Type == lastType,
			arg.Type == nil && lastType == stringType,
			arg.Type == stringType && lastType == nil:
			// skip
		default:
			t := arg.Type
			if t == nil {
				t = stringType
			}
			list += " " + arg.Type.String()
		}
		list += ", " + arg.Name
		lastType = arg.Type
	}
	if lastType == nil {
		lastType = stringType
	}
	return list[2:] + lastType.String()
}

func (a Arglist) QuotedNames() string {
	if len(a) == 0 {
		return ""
	}
	names := ""
	for _, arg := range a {
		names += ", " + fmt.Sprintf("%q", arg.Name)
	}
	return names[2:]
}

// this may break with bad quoting, but the user can influence the breakage
// and it's not at runtime so there's little risk
func (q Query) String() string {
	type staticArg struct {
		pos  int
		name string
	}
	query := `"`
	for _, part := range q {
		switch v := part.(type) {
		case string:
			query += v
		case *Arg:
			switch {
			case v.Dynamic && v.Quoted:
				query += `'?'`
			case v.Dynamic && !v.Quoted:
				query += `?`
			default:
				query += `" + ` + v.Name + ` + "`
			}
		}
	}
	query += `"`
	return query
}

func (a *Arg) String() string {
	return a.Name
}

var gonameregexp *regexp.Regexp

func init() {
	r, err := regexp.Compile("[^A-Za-z0-9]*")
	if err != nil {
		panic(err)
	}
	gonameregexp = r
}

func (c Column) Goname() string {
	name := gonameregexp.ReplaceAllString(c.Name(), "")
	return strings.Title(name)
}

var gotagger = strings.NewReplacer(` `, "_", `:`, "_", `"`, "'")

func (c Column) Declaration() (string, error) {
	name := c.Name()
	goType, err := c.ReflectGoType()
	if err != nil {
		return "", err
	}
	var sqlType string
	if c.MysqlParameters() == mysqlinternals.ParamMustLength {
		sqlType, err = c.MysqlDeclaration(255) // should be varchar mostly
	} else {
		sqlType, err = c.MysqlDeclaration()
	}
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s %s `mysql:\"name=%s,type=%s\"`",
		c.Goname(), goType, gotagger.Replace(name), gotagger.Replace(sqlType)), nil
}

func generate(writer io.Writer, db *sql.DB, name string, query Query) error {
	const structTemplate = `~$args := .Query.Args~
type ~.Name~ struct {~range $i, $e := .Cols~
	~$e.Declaration~~end~
}

func (v *~.Name~) bindTo(t []interface{}) (*~.Name~, []interface{}) {
	if v == nil {
		v := &~.Name~{}
	}
	if len(t) != ~.Cols|len~ {
		t = make([]interface{}, ~.Cols|len~)
	}~range $i, $e := .Cols~
	t[~$i~] = &v.~$e.Goname~~end~
	return &v
}

func (v *~.Name~) BindTo(t []interface{}) (interface{}, []interface{}) {
	return v.bindTo(t)
}

func (v *~.Name~) Query(~$args.Static.Declaration~) (string, []string) {
	return ~.Query~, []string{~$args.Dynamic.QuotedNames~}
}

`
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
	cols := make([]Column, len(resp.columns))
	for i, c := range resp.columns {
		cols[i] = Column{c}
	}
	return templ.Execute(writer, struct {
		Name  string
		Query Query
		Cols  []Column
	}{
		Name:  name,
		Query: query,
		Cols:  cols,
	})
}

func Generate(writer io.Writer, db *sql.DB, pkg string, queries map[string]Query) error {
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
	for name, query := range queries {
		// consider arguments
		err = generate(writer, db, name, query)
		if err != nil {
			return err
		}
	}
	return nil
}
