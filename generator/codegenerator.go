package generator

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

// store names of generated structs (Arg.Name, Arg.ValueType)
type marker map[string]struct{}

// Query represents a Query a struct should be generated for.
// It contains the name of the struct to be created as its first argument.
// The name may be followed by an optional map[int]string to override
// field names of the respective column indices.
// Following that are parts of the SQL (string) and Arg or *Arg (parameters).
type Query []interface{}

// SQL provides a simple routine to create a Query
func SQL(structName string, parts ...interface{}) Query {
	return Query(append([]interface{}{structName}, parts...))
}

// TemplateColumn contains all available information for a
// column retrievable from a mysql response.
// nameOverride is filled with the name to use if the one provided by MySQL
// should not be used
type TemplateColumn struct {
	mysqlinternals.Column
	nameOverride string
}

// Arg represents an argument in a MySQL query
type Arg struct {
	// argument name
	Name string
	// argument value for sample query
	Sample string
	// Dynamic arguments become "?" in the query and are provided at runtime.
	// Static arguments are identifiers and have to be known at query construction time
	// (e.g. table and schema name)
	Dynamic bool
	// argument should be quoted
	Quoted bool
	// argument type (string by default)
	Type reflect.Type
	// alternatives
	ValueType interface{}
	Values    map[string]interface{}
	// description
	Comment string
}

type Tags string

// Arglist represents a sequence of arguments
type Arglist []*Arg

func (q Query) StructName() (name string, ok bool) {
	name, ok = q[0].(string)
	return
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
		case map[interface{}]string, map[int]string:
			// ignore, it's an override map for column names (from int or string)
		default:
			return sql, fmt.Errorf("type %T of argument %d is not supported", part, i)
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
				query += fmt.Sprintf(`" + %s + "`, v.Name)
			}
		}
	}
	query += `"`
	return query
}

func (q Query) Overrides() map[int]string {
	for _, v := range q[1:] {
		if m, ok := v.(map[int]string); ok {
			return m
		}
	}
	return map[int]string{}
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

func (c TemplateColumn) Name() string {
	name := c.Column.Name()
	if c.nameOverride != "" {
		name = c.nameOverride
	}
	return name
}

func (c TemplateColumn) Goname() string {
	return strings.Title(gonameregexp.ReplaceAllString(c.Name(), ""))
}

var gotagger = strings.NewReplacer(` `, "_", `:`, "_", `"`, "'")

func (c TemplateColumn) Declaration() (decl string, err error) {
	var sqlType string
	switch c.MysqlParameters() {
	case mysqlinternals.ParamOneOrMore:
		sqlType = "#ENUM/SET#"
	case mysqlinternals.ParamMustLength:
		// should be varchar mostly
		sqlType, err = c.MysqlDeclaration(255)
	default:
		sqlType, err = c.MysqlDeclaration()
	}
	if err != nil {
		return "", err
	}
	name := c.Name()
	if c.nameOverride != "" {
		name = "#" + c.nameOverride + "#"
	}
	postfix := fmt.Sprintf("`mysqlname:%q mysqltype:%q`", name, sqlType)
	goSafeType, err := c.ReflectSqlType(false)
	if err != nil {
		return "", err
	}
	var comment string
	if goType, err := c.ReflectGoType(); err == nil && goType != goSafeType {
		comment = " // should fit into " + goType.String()
	}
	return fmt.Sprintf("%s %s %s %s", c.Goname(), goSafeType, postfix, comment), nil
}

func (mark *marker) generate(writer io.Writer, db *sql.DB, query Query) error {
	const structTemplate = `~$args := .Query.Args~~$statics := $args.Static~~$dynamics := $args.Dynamic~~$declaration := $statics.Declaration~
type ~.Name~ struct {~range $i, $e := .Cols~
	~$e.Declaration~~end~
}

// bind returns a slice of all columns for a ~.Name~ to use it with Scan.
// Bind only one ~.Name~ and copy it every row after calling Scan.
func (v *~.Name~) Bind() (cols []interface{}) {
	return []interface{}{
	~range $i, $e := .Cols~&v.~$e.Goname~,
	~end~}
}

// copy returns a shallow copy of the ~.Name~.
func (v *~.Name~) copy() *~.Name~ {
	if v == nil {
		return nil
	}
	return &(*v)
}

// Copy returns a shallow copy of the ~.Name~.
func (v *~.Name~) Copy() Bindable {
	return v.copy()
}
~if $statics~
// query provides sql metadata used to query for ~.Name~.
// It returns the SQL string and the arguments needed for prepared statements.
func (v *~.Name~) query(~$statics.Declaration~) (sql string, args []string) {
	return ~.Query~, []string{~$dynamics.QuotedNames~}
}
~end~
// Query provides sql metadata used to query for ~.Name~.
// It returns the SQL string and the arguments needed for prepared statements.
~if $statics~// Query panics if location is not ~$declaration~.~else~// Query panics if arguments are given.~end~
func (v *~.Name~) Query(location...string) (sql string, args []string) {
	if len(location) != ~$statics|len~ {
		panic("Query must be called ~if $declaration~with arguments (~$declaration~)~else~without arguments~end~")
	}~if $statics~
	return v.query(~range $i, $e := $statics~~if $i~, ~end~location[~$i~]~end~)~else~
	return ~.Query~, []string{~$dynamics.QuotedNames~}~end~
}

`
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
	Bind() (cols []interface{})
	Copy() Bindable
	Query(location...string) (sql string, args []string)
}

`
	mark := make(marker)
	_, err := fmt.Fprintf(writer, baseTemplate, pkg)
	if err != nil {
		return err
	}
	for _, query := range queries {
		// consider arguments
		err = (&mark).generate(writer, db, query)
		if err != nil {
			return err
		}
	}
	return nil
}
