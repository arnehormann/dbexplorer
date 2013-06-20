// +build ignore

package main

import (
	"database/sql"
	"github.com/arnehormann/dbexplorer"
	"os"
	"strings"
)

var args = map[string]string{
	"schema": "mysql",
	"table":  "user",
	"user":   "root",
	"host":   "localhost",
}

var queries = map[string]string{
	"CharacterSet":    "SHOW CHARACTER SET",
	"Collation":       "SHOW COLLATION",
	"Engine":          "SHOW ENGINES",
	"Privilege":       "SHOW PRIVILEGES",
	"BinaryLog":       "SHOW BINARY LOGS",
	"InnodbStatus":    "SHOW ENGINE INNODB STATUS",
	"InnodbMutex":     "SHOW ENGINE INNODB MUTEX",
	"Process":         "SHOW PROCESSLIST",
	"MasterStatus":    "SHOW MASTER STATUS",
	"SlaveHost":       "SHOW SLAVE HOSTS",
	"SlaveStatus":     "SHOW SLAVE STATUS",
	"GlobalStatus":    "SHOW GLOBAL STATUS",
	"GlobalVariable":  "SHOW GLOBAL VARIABLES",
	"SessionStatus":   "SHOW SESSION STATUS",
	"SessionVariable": "SHOW SESSION VARIABLES",
	"Warning":         "SHOW WARNINGS",
	"Error":           "SHOW ERRORS",
	"Schema":          "SHOW SCHEMAS",
	"OpenTable":       "SHOW OPEN TABLES",
	"Trigger":         "SHOW TRIGGERS FROM `{{.schema}}`",
	"TableStatus":     "SHOW TABLE STATUS FROM `{{.schema}}`",
	"Column":          "SHOW COLUMNS FROM `{{.schema}}`.`{{.table}}`",
	"Index":           "SHOW INDEXES FROM `{{.schema}}`.`{{.table}}`",
	"Grant":           `SHOW GRANTS FOR "{{.user}}"@"{{.host}}"`,
	// these don't exist in a vanilla db, would have to create them first
	//"View":        "SHOW CREATE VIEW `{{.schema}}`.`{{.view}}`",
	//"CreateFunction":  "SHOW CREATE FUNCTION `{{.schema}}.{{.function}}`",
	//"CreateProcedure": "SHOW CREATE PROCEDURE `{{.schema}}.{{.procedure}}`",
}

func main() {
	dsn := os.Getenv("MYSQL_DSN")
	if dsn == "" {
		dsn = "root@tcp(localhost:3306)/?charset=utf8mb4,utf8&parseTime=false"
	}
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	args := map[string]string{
		"schema": "mysql",
		"table":  "user",
		"user":   "root",
		"host":   "localhost",
	}
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	currentDir := strings.Split(dir, string(os.PathSeparator))
	err = dbexplorer.GenerateAll(os.Stdout, db, currentDir[len(currentDir)-1], queries, args)
	if err != nil {
		panic(err)
	}
}
