package generator

import (
	"fmt"
	"reflect"
	"regexp"
)

// Query represents a Query a struct should be generated for.
// It contains the name of the struct to be created as its first argument.
// The name may be followed by an optional map[int]string to override
// field names of the respective column indices.
// Following that are parts of the SQL (string) and Arg or *Arg (parameters).
type Query []interface{}

func init() {
	r, err := regexp.Compile("[^A-Za-z0-9]*")
	if err != nil {
		panic(err)
	}
	gonameregexp = r
}

// SQL provides a simple routine to create a Query
func SQL(structName string, parts ...interface{}) Query {
	return Query(append([]interface{}{structName}, parts...))
}

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
		case reflect.StructField:
			// ignore, it's an DB-invisible field
		case Map, map[string]string, map[int]string:
			// ignore, it's an override map for field names (maps: from index or column name)
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

func (q Query) String() string {
	// code below may break with bad quoting, but the user can influence the breakage
	// and it's not at runtime so there is little risk
	query := `"`
	argMap := make(map[string]int)
	for _, part := range q[1:] {
		switch v := part.(type) {
		case string:
			query += v
		case *Arg:
			if v.Dynamic {
				if v.Quoted {
					query += `'?'`
				} else {
					query += `?`
				}
			} else {
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
	return nil
}
