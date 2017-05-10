package generator

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/arnehormann/sqlinternals/mysqlinternals"
)

// TemplateColumn contains all available information for a
// column retrievable from a mysql response.
// nameOverride is filled with the name to use if the one provided by MySQL
// should not be used
type TemplateColumn struct {
	// TODO use StructField

	mysqlinternals.Column
	nameOverride string
	skipData     bool
	skipBind     bool
}

var goID = regexp.MustCompile("[^A-Za-z0-9]*")

func (c TemplateColumn) Name() string {
	if c.nameOverride != "" {
		return c.nameOverride
	}
	return c.Column.Name()
}

func (c TemplateColumn) Bindable() bool {
	return !c.skipBind
}

func (c TemplateColumn) Goname() string {
	parts := strings.Split(c.Name(), "_")
	for i := range parts {
		parts[i] = strings.Title(parts[i])
	}
	return goID.ReplaceAllString(strings.Join(parts, ""), "")
}

func (c TemplateColumn) Declaration() (decl string, err error) {
	var sqlType string
	switch c.MysqlParameters() {
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
		name = c.nameOverride
	}
	goname := c.Goname()
	dataname := strings.ToLower(goname) + ",omitempty"
	if c.skipData {
		dataname = "-"
	}
	goSafeType, err := c.ReflectSqlType(false)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf(
		"%[1]s %[2]s `mysql:\"%[3]s,%[4]s\" json:%[5]q xml:%[5]q`",
		goname, goSafeType, name, sqlType, dataname,
	), nil
}
