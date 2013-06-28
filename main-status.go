// +build ignore

package main

import (
	"database/sql"
	dbe "github.com/arnehormann/dbexplorer"
	"os"
	"strings"
)

var (
	argSchema = &dbe.Arg{
		Name:   "schema",
		Sample: "mysql",
	}
	argTable = &dbe.Arg{
		Name:   "table",
		Sample: "user",
	}
	argUser = &dbe.Arg{
		Name:    "user",
		Sample:  "root",
		Dynamic: true,
		Quoted:  true,
	}
	argHost = &dbe.Arg{
		Name:    "host",
		Sample:  "localhost",
		Dynamic: true,
		Quoted:  true,
	}
	argStatus = &dbe.Arg{
		Name:      "scope",
		Sample:    "SESSION STATUS",
		ValueType: "scope",
		Values: map[string]interface{}{
			"SessionStatus":    "SESSION STATUS",
			"SessionVariables": "SESSION VARIABLES",
			"GlobalStatus":     "GLOBAL STATUS",
			"GlobalVariables":  "GLOBAL VARIABLES",
		},
	}
	queries = map[string]dbe.Query{
		//"BinaryLog":    dbe.SQL("SHOW BINARY LOGS"),
		"CharacterSet": dbe.SQL("SHOW CHARACTER SET"),
		"Collation":    dbe.SQL("SHOW COLLATION"),
		"Column":       dbe.SQL("SHOW COLUMNS FROM `", argSchema, "`.`", argTable, "`"),
		"Engine":       dbe.SQL("SHOW ENGINES"),
		"Error":        dbe.SQL("SHOW ERRORS"),
		"Grant":        dbe.SQL("SHOW GRANTS FOR ", argUser, "@", argHost),
		"Index":        dbe.SQL("SHOW INDEXES FROM `", argSchema, "`.`", argTable, "`"),
		"InnodbMutex":  dbe.SQL("SHOW ENGINE INNODB MUTEX"),
		"InnodbStatus": dbe.SQL("SHOW ENGINE INNODB STATUS"),
		"MasterStatus": dbe.SQL("SHOW MASTER STATUS"),
		"OpenTable":    dbe.SQL("SHOW OPEN TABLES"),
		"Privilege":    dbe.SQL("SHOW PRIVILEGES"),
		"Process":      dbe.SQL("SHOW PROCESSLIST"),
		"Schema":       dbe.SQL("SHOW SCHEMAS"),
		"SlaveHost":    dbe.SQL("SHOW SLAVE HOSTS"),
		"SlaveStatus":  dbe.SQL("SHOW SLAVE STATUS"),
		"TableStatus":  dbe.SQL("SHOW TABLE STATUS FROM `", argSchema, "`"),
		"Trigger":      dbe.SQL("SHOW TRIGGERS FROM `", argSchema, "`"),
		"Variables":    dbe.SQL("SHOW ", argStatus),
		"Warning":      dbe.SQL("SHOW WARNINGS"),
		// these don't exist in a vanilla db, would have to create them first
		//"View":        dbe.SQL("SHOW CREATE VIEW `{{.schema}}`.`{{.view}}`"),
		//"CreateFunction":  dbe.SQL("SHOW CREATE FUNCTION `{{.schema}}.{{.function}}`"),
		//"CreateProcedure": dbe.SQL("SHOW CREATE PROCEDURE `{{.schema}}.{{.procedure}}`"),
	}
)

func main() {
	dsn := os.Getenv("MYSQL_DSN")
	if dsn == "" {
		dsn = "root@tcp(localhost:3306)/?charset=utf8mb4,utf8&parseTime=false"
	}
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	currentDir := strings.Split(dir, string(os.PathSeparator))
	err = dbe.Generate(os.Stdout, db, currentDir[len(currentDir)-1], queries)
	if err != nil {
		panic(err)
	}
}
