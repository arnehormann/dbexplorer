package generator

import (
	"fmt"
	"github.com/arnehormann/sqlinternals/mysqlinternals"
	"regexp"
	"strings"
)

// TemplateColumn contains all available information for a
// column retrievable from a mysql response.
// nameOverride is filled with the name to use if the one provided by MySQL
// should not be used
type TemplateColumn struct {
	mysqlinternals.Column
	nameOverride string
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
	return fmt.Sprintf("%s %s %s", c.Goname(), goSafeType, postfix), nil
}
