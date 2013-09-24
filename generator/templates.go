package generator

const (
	// top of the generated file
	baseTemplate = `// this file is generated, do not edit it!

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
	// per query struct
	structTemplate = `~$args := .Query.Args~~$statics := $args.Static~~$dynamics := $args.Dynamic~~$declaration := $statics.Declaration~
type ~.Name~ struct {~range $i, $e := .Cols~
	~$e.Declaration~~end~
}

// Bind returns a slice of all columns for a ~.Name~ to use it with Scan.
// Bind only one ~.Name~ and copy it every row after calling Scan.
func (v *~.Name~) Bind() (cols []interface{}) {
	return []interface{}{
	~range $i, $e := .Cols~	&v.~$e.Goname~,
	~end~}
}

// copy returns a shallow copy of the ~.Name~.
func (v *~.Name~) copy() *~.Name~ {
	if v == nil {
		return nil
	}
	w := *v
	return &w
}

// Copy returns a shallow copy of the ~.Name~.
func (v *~.Name~) Copy() Bindable {
	return v.copy()
}
~if $statics~
// query provides sql metadata used to query for ~.Name~.
// It returns the SQL string and the arguments needed for prepared statements.
func (v *~.Name~) query(~$statics.Declaration~) (sql string, args []string) {
	return ~.Query~, ~if $dynamics~[]string{~$dynamics.QuotedNames~}~else~nil~end~
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
	return ~.Query~, ~if $dynamics~[]string{~$dynamics.QuotedNames~}~else~nil~end~~end~
}

`
)
