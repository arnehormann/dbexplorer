// +build ignore

// TODO maybe: provide overrides for nullability if MySQL rows vary?

package main

import (
	"database/sql"
	dbg "local/github.com/arnehormann/dbexplorer/generator"
	"os"
	"strings"
)

var (
	argSchema = &dbg.Arg{
		Name:   "schema",
		Sample: "mysql",
	}
	argTable = &dbg.Arg{
		Name:   "table",
		Sample: "user",
	}
	argUser = &dbg.Arg{
		Name:    "user",
		Sample:  "root",
		Dynamic: true,
		Quoted:  true,
	}
	argHost = &dbg.Arg{
		Name:    "host",
		Sample:  "localhost",
		Dynamic: true,
		Quoted:  true,
	}
	queries = []dbg.Query{
		// static information
		dbg.SQL("CharacterSet", "SHOW CHARACTER SET"),
		dbg.SQL("Collation", "SHOW COLLATION"),
		dbg.SQL("Engine", "SHOW ENGINES"),
		dbg.SQL("Privilege", "SHOW PRIVILEGES"),

		// database structure
		dbg.SQL("Grant", map[int]string{0: "Grant"}, "SHOW GRANTS FOR ", argUser, "@", argHost),
		dbg.SQL("Schema", "SHOW SCHEMAS"),
		dbg.SQL("TableStatus", "SHOW TABLE STATUS FROM `", argSchema, "`"),
		dbg.SQL("Column", "SHOW COLUMNS FROM `", argSchema, "`.`", argTable, "`"),
		dbg.SQL("Index", "SHOW INDEXES FROM `", argSchema, "`.`", argTable, "`"),
		//dbg.SQL("Trigger", "SHOW TRIGGERS FROM `", argSchema, "`"),
		//dbg.SQL("CreateView", "SHOW CREATE VIEW `{{.schema}}`.`{{.view}}`"),
		//dbg.SQL("CreateFunction", "SHOW CREATE FUNCTION `{{.schema}}.{{.function}}`"),
		//dbg.SQL("CreateProcedure", "SHOW CREATE PROCEDURE `{{.schema}}.{{.procedure}}`"),

		// status and variables
		dbg.SQL("SessionStatus", "SHOW SESSION STATUS"),
		dbg.SQL("SessionVariables", "SHOW SESSION VARIABLES"),
		dbg.SQL("GlobalStatus", "SHOW GLOBAL STATUS"),
		dbg.SQL("GlobalVariables", "SHOW GLOBAL VARIABLES"),

		// last connection status
		dbg.SQL("Error", "SHOW ERRORS"),
		dbg.SQL("Warning", "SHOW WARNINGS"),

		// tuning and monitoring
		dbg.SQL("InnodbMutex", "SHOW ENGINE INNODB MUTEX"),
		dbg.SQL("InnodbStatus", "SHOW ENGINE INNODB STATUS"),
		dbg.SQL("OpenTable", "SHOW OPEN TABLES"),
		dbg.SQL("Process", "SHOW PROCESSLIST"),
		//dbg.SQL("BinaryLog", "SHOW BINARY LOGS"),

		// replication
		dbg.SQL("MasterStatus", "SHOW MASTER STATUS"),
		dbg.SQL("SlaveHost", "SHOW SLAVE HOSTS"),
		dbg.SQL("SlaveStatus", "SHOW SLAVE STATUS"),

		// alternative, already and better provided by TABLE STATUS
		//dbg.SQL("CreateTable", "SHOW CREATE TABLE `", argSchema, "`.`", argTable, "`"),
	}
)

func main() {
	dsn := os.Getenv("MYSQL_DSN")
	if dsn == "" {
		dsn = "root@tcp(localhost:3306)/?parseTime=false"
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
	err = dbg.Generate(os.Stdout, db, currentDir[len(currentDir)-1], queries...)
	if err != nil {
		panic(err)
	}
}
