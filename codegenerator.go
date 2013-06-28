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

func SQL(structName string, parts ...interface{}) Query {
	return Query(append([]interface{}{structName}, parts...))
}

type TemplateColumn struct {
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

func (q Query) StructName() string {
	if name, ok := q[0].(string); ok {
		return name
	}
	return ""
}

func (q Query) Sample() (string, error) {
	sql := ""
	for i, part := range q[1:] {
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
	var lastType reflect.Type = stringType
	list := ""
	for _, arg := range a {
		atype := arg.Type
		if atype == nil {
			atype = stringType
		}
		switch atype {
		case lastType:
			// skip
		default:
			t := atype
			if t == nil {
				t = stringType
			}
			list += " " + atype.String()
		}
		list += ", " + arg.Name
		lastType = atype
	}
	if lastType == nil {
		lastType = stringType
	}
	return list[2:] + " " + lastType.String()
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
	argMap := make(map[string]int)
	for _, part := range q[1:] {
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
				idx, ok := argMap[v.Name]
				if !ok {
					idx = len(argMap)
					argMap[v.Name] = idx
				}
				query += fmt.Sprintf(`" + location[%d] + "`, idx)
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

func (c TemplateColumn) Goname() string {
	name := gonameregexp.ReplaceAllString(c.Name(), "")
	return strings.Title(name)
}

var gotagger = strings.NewReplacer(` `, "_", `:`, "_", `"`, "'")

func (c TemplateColumn) Declaration() (string, error) {
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
	return fmt.Sprintf("%s %s // %s", c.Goname(), goType, sqlType), nil
}

func generate(writer io.Writer, db *sql.DB, query Query) error {
	const structTemplate = `~$args := .Query.Args~
type ~.Name~ struct {~range $i, $e := .Cols~
	~$e.Declaration~~end~
}

func (v *~.Name~) bindTo(t []interface{}) (*~.Name~, []interface{}) {
	if v == nil {
		v = &~.Name~{}
	}
	if len(t) != ~.Cols|len~ {
		t = make([]interface{}, ~.Cols|len~)
	}~range $i, $e := .Cols~
	t[~$i~] = &v.~$e.Goname~~end~
	return v, t
}

func (v *~.Name~) BindTo(t []interface{}) (Bindable, []interface{}) {
	return v.bindTo(t)
}

func (v *~.Name~) Columns() []string {
	return []string{~range $i, $e := .Cols~
		~printf "%q" $e.Name~,~end~
	}
}

func (v *~.Name~) copy() *~.Name~ {
	if v == nil {
		return nil
	}
	return &(*v)
}

func (v *~.Name~) Copy() Bindable {
	return v.copy()
}

~$statics := $args.Static~~if $statics~// location is: ~$statics.Declaration~
~end~func (v *~.Name~) Query(location...string) (string, []string) {
	return ~.Query~, []string{~$args.Dynamic.QuotedNames~}
}

`
	name := query.StructName()
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
	cols := make([]TemplateColumn, len(resp.columns))
	for i, c := range resp.columns {
		cols[i] = TemplateColumn{c}
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

type Bindable interface {
	BindTo(t []interface{}) (Bindable, []interface{})
	Columns() []string
	Copy() Bindable
	Query(location...string) (string, []string)
}

`
	_, err := fmt.Fprintf(writer, baseTemplate, pkg)
	if err != nil {
		return err
	}
	for _, query := range queries {
		// consider arguments
		err = generate(writer, db, query)
		if err != nil {
			return err
		}
	}
	return nil
}
