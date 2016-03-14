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

var gonameregexp = regexp.MustCompile("[^A-Za-z0-9]*")

func (c TemplateColumn) Name() string {
	name := c.Column.Name()
	if c.nameOverride != "" {
		name = c.nameOverride
	}
	return name
}

func (c TemplateColumn) Bindable() bool {
	return !c.skipBind
}

func (c TemplateColumn) Goname() string {
	return strings.Title(gonameregexp.ReplaceAllString(c.Name(), ""))
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
		name = "#" + c.nameOverride + "#"
	}
	goname := c.Goname()
	dataname := strings.ToLower(goname) + ",omitempty"
	if c.skipData {
		dataname = "-"
	}
	postfix := fmt.Sprintf(
		"`mysql:\"%[1]s,%[2]s\" json:%[3]q xml:%[3]q`",
		name,
		sqlType,
		dataname, //TODO differentiate between json / xml
	)
	goSafeType, err := c.ReflectSqlType(false)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s %s %s", goname, goSafeType, postfix), nil
}
