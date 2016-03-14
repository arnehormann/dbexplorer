// +build ignore

// TODO maybe: provide overrides for nullability if MySQL rows vary?

package main

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	gen "github.com/arnehormann/dbexplorer/generator"
)

var (
	// default arguments for variable identifiers that can not be transported
	// by `?` wildcards:

	// schema
	argSchema = &gen.Arg{
		Name:   "schema",
		Sample: "mysql",
	}

	// table
	argTable = &gen.Arg{
		Name:   "table",
		Sample: "user",
	}

	// user
	argUser = &gen.Arg{
		Name:    "user",
		Sample:  "root",
		Dynamic: true,
		Quoted:  true,
	}

	// host
	argHost = &gen.Arg{
		Name:    "host",
		Sample:  "localhost",
		Dynamic: true,
		Quoted:  true,
	}

	// struct name followed by query (which may be assembled from comma separated parts)
	queries = []gen.Query{
		// static information
		gen.SQL("CharacterSet", "SHOW CHARACTER SET"),
		gen.SQL("Collation", "SHOW COLLATION"),
		gen.SQL("Engine", "SHOW ENGINES"),
		gen.SQL("Privilege", "SHOW PRIVILEGES"),

		// database structure
		// the map argument indicates a renamed field - it is dynamic in Columns and we want to force a name
		gen.SQL("Grant", map[int]string{0: "Grant"}, "SHOW GRANTS FOR ", argUser, "@", argHost),
		gen.SQL("Schema", "SHOW SCHEMAS"),
		gen.SQL("TableStatus", "SHOW TABLE STATUS FROM `", argSchema, "`"),
		gen.SQL("Column", "SHOW COLUMNS FROM `", argSchema, "`.`", argTable, "`"),
		gen.SQL("Index", "SHOW INDEXES FROM `", argSchema, "`.`", argTable, "`"),
		//gen.SQL("Trigger", "SHOW TRIGGERS FROM `", argSchema, "`"),
		//gen.SQL("CreateView", "SHOW CREATE VIEW `{{.schema}}`.`{{.view}}`"),
		//gen.SQL("CreateFunction", "SHOW CREATE FUNCTION `{{.schema}}.{{.function}}`"),
		//gen.SQL("CreateProcedure", "SHOW CREATE PROCEDURE `{{.schema}}.{{.procedure}}`"),

		// status and variables
		gen.SQL("SessionStatus", "SHOW SESSION STATUS"),
		gen.SQL("SessionVariables", "SHOW SESSION VARIABLES"),
		gen.SQL("GlobalStatus", "SHOW GLOBAL STATUS"),
		gen.SQL("GlobalVariables", "SHOW GLOBAL VARIABLES"),

		// last connection status
		gen.SQL("Error", "SHOW ERRORS"),
		gen.SQL("Warning", "SHOW WARNINGS"),

		// tuning and monitoring
		gen.SQL("InnodbMutex", "SHOW ENGINE INNODB MUTEX"),
		gen.SQL("InnodbStatus", "SHOW ENGINE INNODB STATUS"),
		gen.SQL("OpenTable", "SHOW OPEN TABLES"),
		gen.SQL("Process", "SHOW PROCESSLIST"),
		//gen.SQL("BinaryLog", "SHOW BINARY LOGS"),

		// replication
		gen.SQL("MasterStatus", "SHOW MASTER STATUS"),
		gen.SQL("SlaveHost", "SHOW SLAVE HOSTS"),
		gen.SQL("SlaveStatus", "SHOW SLAVE STATUS"),

		// alternative, already and better provided by TABLE STATUS
		//gen.SQL("CreateTable", "SHOW CREATE TABLE `", argSchema, "`.`", argTable, "`"),

		gen.SQL("User", "SELECT user, host FROM mysql.user"),
	}
)

func main() {
	var errmsg = "Please set the environment variable MYSQL_DSN (e.g. 'root@tcp(localhost:3306)/?parseTime=false')"
	dsn := os.Getenv("MYSQL_DSN")
	if dsn == "" {
		fmt.Fprintln(os.Stderr, errmsg)
		os.Exit(-1)
	}
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Fprintln(os.Stderr, errmsg)
		panic(err)
	}
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	currentDir := strings.Split(dir, string(os.PathSeparator))
	packageName := currentDir[len(currentDir)-1]
	if len(os.Args) > 1 {
		packageName = os.Args[1]
	}
	err = gen.Generate(os.Stdout, db, packageName, queries...)
	if err != nil {
		panic(err)
	}
}
