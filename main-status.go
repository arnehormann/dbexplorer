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
	queries = []dbe.Query{
		dbe.SQL("CharacterSet", "SHOW CHARACTER SET"),
		dbe.SQL("Collation", "SHOW COLLATION"),
		dbe.SQL("Column", "SHOW COLUMNS FROM `", argSchema, "`.`", argTable, "`"),
		dbe.SQL("Engine", "SHOW ENGINES"),
		dbe.SQL("Error", "SHOW ERRORS"),
		dbe.SQL("Grant", "SHOW GRANTS FOR ", argUser, "@", argHost),
		dbe.SQL("Index", "SHOW INDEXES FROM `", argSchema, "`.`", argTable, "`"),
		dbe.SQL("InnodbMutex", "SHOW ENGINE INNODB MUTEX"),
		dbe.SQL("InnodbStatus", "SHOW ENGINE INNODB STATUS"),
		dbe.SQL("MasterStatus", "SHOW MASTER STATUS"),
		dbe.SQL("OpenTable", "SHOW OPEN TABLES"),
		dbe.SQL("Privilege", "SHOW PRIVILEGES"),
		dbe.SQL("Process", "SHOW PROCESSLIST"),
		dbe.SQL("Schema", "SHOW SCHEMAS"),
		dbe.SQL("SlaveHost", "SHOW SLAVE HOSTS"),
		dbe.SQL("SlaveStatus", "SHOW SLAVE STATUS"),
		dbe.SQL("TableStatus", "SHOW TABLE STATUS FROM `", argSchema, "`"),
		dbe.SQL("Trigger", "SHOW TRIGGERS FROM `", argSchema, "`"),
		dbe.SQL("Variables", "SHOW ", argStatus),
		dbe.SQL("Warning", "SHOW WARNINGS"),

		// disabled in my system
		//dbe.SQL("BinaryLog", "SHOW BINARY LOGS"),

		// provided by TABLE STATUS
		//dbe.SQL("CreateTable", "SHOW CREATE TABLE `", argSchema, "`.`", argTable, "`"),

		// these don't exist in a vanilla db, would have to create them first
		//dbe.SQL("CreateFunction", "SHOW CREATE FUNCTION `{{.schema}}.{{.function}}`"),
		//dbe.SQL("CreateProcedure", "SHOW CREATE PROCEDURE `{{.schema}}.{{.procedure}}`"),
		//dbe.SQL("CreateView", "SHOW CREATE VIEW `{{.schema}}`.`{{.view}}`"),
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
	err = dbe.Generate(os.Stdout, db, currentDir[len(currentDir)-1], queries...)
	if err != nil {
		panic(err)
	}
}
