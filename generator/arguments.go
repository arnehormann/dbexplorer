package generator

import (
	"fmt"
	"reflect"
)

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
		names += ", " + fmt.Sprintf("%q", arg.Name)
	}
	return names[2:]
}

func (a *Arg) String() string {
	return a.Name
}
