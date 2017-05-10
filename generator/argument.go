package generator

import (
	"reflect"
	"strconv"
)

// Map provides a name for a field in a scope ("go", "mysql", "json", "xml").
type Map func(name string, scope string) string

// Arg represents an argument in a MySQL query
type Arg struct {
	// argument name
	Name string `json:"name"`
	// argument value for sample query
	Sample string `json:"example"`
	// Dynamic arguments become "?" in the query and are provided at runtime.
	// Static arguments are identifiers and have to be known at query construction time
	// (e.g. table and schema name)
	Dynamic bool `json:"dynamic,omitempty"`
	// argument should be quoted
	Quoted bool `json:"quote,omitempty"`
	// argument type (string by default)
	Type reflect.Type `json:"-"`
	// alternatives
	ValueType interface{}            `json:"-"`
	Values    map[string]interface{} `json:"-"`
	// description
	Comment string `json:"-"`
}

// Arglist represents a sequence of arguments
type Arglist []*Arg

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
		names += ", " + strconv.Quote(arg.Name)
	}
	return names[2:]
}

func (a *Arg) String() string {
	return a.Name
}
