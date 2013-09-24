// this file is generated, do not edit it!

package main

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"time"
)

// force import even if unused to avoid writing detection code
var (
	_ = time.Monday
	_ = mysql.NullTime{}
	_ = sql.NullString{}
)

type Bindable interface {
	Bind() (cols []interface{})
	Copy() Bindable
	Query(location ...string) (sql string, args []string)
}

type CharacterSet struct {
	Charset          string `mysqlname:"Charset" mysqltype:"VARCHAR(255) NOT NULL"`
	Description      string `mysqlname:"Description" mysqltype:"VARCHAR(255) NOT NULL"`
	Defaultcollation string `mysqlname:"Default collation" mysqltype:"VARCHAR(255) NOT NULL"`
	Maxlen           int64  `mysqlname:"Maxlen" mysqltype:"BIGINT NOT NULL"`
}

// Bind returns a slice of all columns for a CharacterSet to use it with Scan.
// Bind only one CharacterSet and copy it every row after calling Scan.
func (v *CharacterSet) Bind() (cols []interface{}) {
	return []interface{}{
		&v.Charset,
		&v.Description,
		&v.Defaultcollation,
		&v.Maxlen,
	}
}

// copy returns a shallow copy of the CharacterSet.
func (v *CharacterSet) copy() *CharacterSet {
	if v == nil {
		return nil
	}
	w := *v
	return &w
}

// Copy returns a shallow copy of the CharacterSet.
func (v *CharacterSet) Copy() Bindable {
	return v.copy()
}

// Query provides sql metadata used to query for CharacterSet.
// It returns the SQL string and the arguments needed for prepared statements.
// Query panics if arguments are given.
func (v *CharacterSet) Query(location ...string) (sql string, args []string) {
	if len(location) != 0 {
		panic("Query must be called without arguments")
	}
	return "SHOW CHARACTER SET", nil
}

type Collation struct {
	Collation string `mysqlname:"Collation" mysqltype:"VARCHAR(255) NOT NULL"`
	Charset   string `mysqlname:"Charset" mysqltype:"VARCHAR(255) NOT NULL"`
	Id        int64  `mysqlname:"Id" mysqltype:"BIGINT NOT NULL"`
	Default   string `mysqlname:"Default" mysqltype:"VARCHAR(255) NOT NULL"`
	Compiled  string `mysqlname:"Compiled" mysqltype:"VARCHAR(255) NOT NULL"`
	Sortlen   int64  `mysqlname:"Sortlen" mysqltype:"BIGINT NOT NULL"`
}

// Bind returns a slice of all columns for a Collation to use it with Scan.
// Bind only one Collation and copy it every row after calling Scan.
func (v *Collation) Bind() (cols []interface{}) {
	return []interface{}{
		&v.Collation,
		&v.Charset,
		&v.Id,
		&v.Default,
		&v.Compiled,
		&v.Sortlen,
	}
}

// copy returns a shallow copy of the Collation.
func (v *Collation) copy() *Collation {
	if v == nil {
		return nil
	}
	w := *v
	return &w
}

// Copy returns a shallow copy of the Collation.
func (v *Collation) Copy() Bindable {
	return v.copy()
}

// Query provides sql metadata used to query for Collation.
// It returns the SQL string and the arguments needed for prepared statements.
// Query panics if arguments are given.
func (v *Collation) Query(location ...string) (sql string, args []string) {
	if len(location) != 0 {
		panic("Query must be called without arguments")
	}
	return "SHOW COLLATION", nil
}

type Engine struct {
	Engine       string         `mysqlname:"Engine" mysqltype:"VARCHAR(255) NOT NULL"`
	Support      string         `mysqlname:"Support" mysqltype:"VARCHAR(255) NOT NULL"`
	Comment      string         `mysqlname:"Comment" mysqltype:"VARCHAR(255) NOT NULL"`
	Transactions sql.NullString `mysqlname:"Transactions" mysqltype:"VARCHAR(255)"`
	XA           sql.NullString `mysqlname:"XA" mysqltype:"VARCHAR(255)"`
	Savepoints   sql.NullString `mysqlname:"Savepoints" mysqltype:"VARCHAR(255)"`
}

// Bind returns a slice of all columns for a Engine to use it with Scan.
// Bind only one Engine and copy it every row after calling Scan.
func (v *Engine) Bind() (cols []interface{}) {
	return []interface{}{
		&v.Engine,
		&v.Support,
		&v.Comment,
		&v.Transactions,
		&v.XA,
		&v.Savepoints,
	}
}

// copy returns a shallow copy of the Engine.
func (v *Engine) copy() *Engine {
	if v == nil {
		return nil
	}
	w := *v
	return &w
}

// Copy returns a shallow copy of the Engine.
func (v *Engine) Copy() Bindable {
	return v.copy()
}

// Query provides sql metadata used to query for Engine.
// It returns the SQL string and the arguments needed for prepared statements.
// Query panics if arguments are given.
func (v *Engine) Query(location ...string) (sql string, args []string) {
	if len(location) != 0 {
		panic("Query must be called without arguments")
	}
	return "SHOW ENGINES", nil
}

type Privilege struct {
	Privilege string `mysqlname:"Privilege" mysqltype:"VARCHAR(255) NOT NULL"`
	Context   string `mysqlname:"Context" mysqltype:"VARCHAR(255) NOT NULL"`
	Comment   string `mysqlname:"Comment" mysqltype:"VARCHAR(255) NOT NULL"`
}

// Bind returns a slice of all columns for a Privilege to use it with Scan.
// Bind only one Privilege and copy it every row after calling Scan.
func (v *Privilege) Bind() (cols []interface{}) {
	return []interface{}{
		&v.Privilege,
		&v.Context,
		&v.Comment,
	}
}

// copy returns a shallow copy of the Privilege.
func (v *Privilege) copy() *Privilege {
	if v == nil {
		return nil
	}
	w := *v
	return &w
}

// Copy returns a shallow copy of the Privilege.
func (v *Privilege) Copy() Bindable {
	return v.copy()
}

// Query provides sql metadata used to query for Privilege.
// It returns the SQL string and the arguments needed for prepared statements.
// Query panics if arguments are given.
func (v *Privilege) Query(location ...string) (sql string, args []string) {
	if len(location) != 0 {
		panic("Query must be called without arguments")
	}
	return "SHOW PRIVILEGES", nil
}

type Grant struct {
	Grant string `mysqlname:"#Grant#" mysqltype:"VARCHAR(255) NOT NULL"`
}

// Bind returns a slice of all columns for a Grant to use it with Scan.
// Bind only one Grant and copy it every row after calling Scan.
func (v *Grant) Bind() (cols []interface{}) {
	return []interface{}{
		&v.Grant,
	}
}

// copy returns a shallow copy of the Grant.
func (v *Grant) copy() *Grant {
	if v == nil {
		return nil
	}
	w := *v
	return &w
}

// Copy returns a shallow copy of the Grant.
func (v *Grant) Copy() Bindable {
	return v.copy()
}

// Query provides sql metadata used to query for Grant.
// It returns the SQL string and the arguments needed for prepared statements.
// Query panics if arguments are given.
func (v *Grant) Query(location ...string) (sql string, args []string) {
	if len(location) != 0 {
		panic("Query must be called without arguments")
	}
	return "SHOW GRANTS FOR '?'@'?'", []string{"user", "host"}
}

type Schema struct {
	Database string `mysqlname:"Database" mysqltype:"VARCHAR(255) NOT NULL"`
}

// Bind returns a slice of all columns for a Schema to use it with Scan.
// Bind only one Schema and copy it every row after calling Scan.
func (v *Schema) Bind() (cols []interface{}) {
	return []interface{}{
		&v.Database,
	}
}

// copy returns a shallow copy of the Schema.
func (v *Schema) copy() *Schema {
	if v == nil {
		return nil
	}
	w := *v
	return &w
}

// Copy returns a shallow copy of the Schema.
func (v *Schema) Copy() Bindable {
	return v.copy()
}

// Query provides sql metadata used to query for Schema.
// It returns the SQL string and the arguments needed for prepared statements.
// Query panics if arguments are given.
func (v *Schema) Query(location ...string) (sql string, args []string) {
	if len(location) != 0 {
		panic("Query must be called without arguments")
	}
	return "SHOW SCHEMAS", nil
}

type TableStatus struct {
	Name          string         `mysqlname:"Name" mysqltype:"VARCHAR(255) NOT NULL"`
	Engine        sql.NullString `mysqlname:"Engine" mysqltype:"VARCHAR(255)"`
	Version       sql.NullInt64  `mysqlname:"Version" mysqltype:"BIGINT UNSIGNED"`
	Rowformat     sql.NullString `mysqlname:"Row_format" mysqltype:"VARCHAR(255)"`
	Rows          sql.NullInt64  `mysqlname:"Rows" mysqltype:"BIGINT UNSIGNED"`
	Avgrowlength  sql.NullInt64  `mysqlname:"Avg_row_length" mysqltype:"BIGINT UNSIGNED"`
	Datalength    sql.NullInt64  `mysqlname:"Data_length" mysqltype:"BIGINT UNSIGNED"`
	Maxdatalength sql.NullInt64  `mysqlname:"Max_data_length" mysqltype:"BIGINT UNSIGNED"`
	Indexlength   sql.NullInt64  `mysqlname:"Index_length" mysqltype:"BIGINT UNSIGNED"`
	Datafree      sql.NullInt64  `mysqlname:"Data_free" mysqltype:"BIGINT UNSIGNED"`
	Autoincrement sql.NullInt64  `mysqlname:"Auto_increment" mysqltype:"BIGINT UNSIGNED"`
	Createtime    mysql.NullTime `mysqlname:"Create_time" mysqltype:"DATETIME"`
	Updatetime    mysql.NullTime `mysqlname:"Update_time" mysqltype:"DATETIME"`
	Checktime     mysql.NullTime `mysqlname:"Check_time" mysqltype:"DATETIME"`
	Collation     sql.NullString `mysqlname:"Collation" mysqltype:"VARCHAR(255)"`
	Checksum      sql.NullInt64  `mysqlname:"Checksum" mysqltype:"BIGINT UNSIGNED"`
	Createoptions sql.NullString `mysqlname:"Create_options" mysqltype:"VARCHAR(255)"`
	Comment       string         `mysqlname:"Comment" mysqltype:"VARCHAR(255) NOT NULL"`
}

// Bind returns a slice of all columns for a TableStatus to use it with Scan.
// Bind only one TableStatus and copy it every row after calling Scan.
func (v *TableStatus) Bind() (cols []interface{}) {
	return []interface{}{
		&v.Name,
		&v.Engine,
		&v.Version,
		&v.Rowformat,
		&v.Rows,
		&v.Avgrowlength,
		&v.Datalength,
		&v.Maxdatalength,
		&v.Indexlength,
		&v.Datafree,
		&v.Autoincrement,
		&v.Createtime,
		&v.Updatetime,
		&v.Checktime,
		&v.Collation,
		&v.Checksum,
		&v.Createoptions,
		&v.Comment,
	}
}

// copy returns a shallow copy of the TableStatus.
func (v *TableStatus) copy() *TableStatus {
	if v == nil {
		return nil
	}
	w := *v
	return &w
}

// Copy returns a shallow copy of the TableStatus.
func (v *TableStatus) Copy() Bindable {
	return v.copy()
}

// query provides sql metadata used to query for TableStatus.
// It returns the SQL string and the arguments needed for prepared statements.
func (v *TableStatus) query(schema string) (sql string, args []string) {
	return "SHOW TABLE STATUS FROM `" + schema + "`", nil
}

// Query provides sql metadata used to query for TableStatus.
// It returns the SQL string and the arguments needed for prepared statements.
// Query panics if location is not schema string.
func (v *TableStatus) Query(location ...string) (sql string, args []string) {
	if len(location) != 1 {
		panic("Query must be called with arguments (schema string)")
	}
	return v.query(location[0])
}

type Column struct {
	Field   string  `mysqlname:"Field" mysqltype:"VARCHAR(255) NOT NULL"`
	Type    []uint8 `mysqlname:"Type" mysqltype:"BLOB NOT NULL"`
	Null    string  `mysqlname:"Null" mysqltype:"VARCHAR(255) NOT NULL"`
	Key     string  `mysqlname:"Key" mysqltype:"VARCHAR(255) NOT NULL"`
	Default []uint8 `mysqlname:"Default" mysqltype:"BLOB"`
	Extra   string  `mysqlname:"Extra" mysqltype:"VARCHAR(255) NOT NULL"`
}

// Bind returns a slice of all columns for a Column to use it with Scan.
// Bind only one Column and copy it every row after calling Scan.
func (v *Column) Bind() (cols []interface{}) {
	return []interface{}{
		&v.Field,
		&v.Type,
		&v.Null,
		&v.Key,
		&v.Default,
		&v.Extra,
	}
}

// copy returns a shallow copy of the Column.
func (v *Column) copy() *Column {
	if v == nil {
		return nil
	}
	w := *v
	return &w
}

// Copy returns a shallow copy of the Column.
func (v *Column) Copy() Bindable {
	return v.copy()
}

// query provides sql metadata used to query for Column.
// It returns the SQL string and the arguments needed for prepared statements.
func (v *Column) query(schema, table string) (sql string, args []string) {
	return "SHOW COLUMNS FROM `" + schema + "`.`" + table + "`", nil
}

// Query provides sql metadata used to query for Column.
// It returns the SQL string and the arguments needed for prepared statements.
// Query panics if location is not schema, table string.
func (v *Column) Query(location ...string) (sql string, args []string) {
	if len(location) != 2 {
		panic("Query must be called with arguments (schema, table string)")
	}
	return v.query(location[0], location[1])
}

type Index struct {
	Table        string         `mysqlname:"Table" mysqltype:"VARCHAR(255) NOT NULL"`
	Nonunique    int64          `mysqlname:"Non_unique" mysqltype:"BIGINT NOT NULL"`
	Keyname      string         `mysqlname:"Key_name" mysqltype:"VARCHAR(255) NOT NULL"`
	Seqinindex   int64          `mysqlname:"Seq_in_index" mysqltype:"BIGINT NOT NULL"`
	Columnname   string         `mysqlname:"Column_name" mysqltype:"VARCHAR(255) NOT NULL"`
	Collation    sql.NullString `mysqlname:"Collation" mysqltype:"VARCHAR(255)"`
	Cardinality  sql.NullInt64  `mysqlname:"Cardinality" mysqltype:"BIGINT"`
	Subpart      sql.NullInt64  `mysqlname:"Sub_part" mysqltype:"BIGINT"`
	Packed       sql.NullString `mysqlname:"Packed" mysqltype:"VARCHAR(255)"`
	Null         string         `mysqlname:"Null" mysqltype:"VARCHAR(255) NOT NULL"`
	Indextype    string         `mysqlname:"Index_type" mysqltype:"VARCHAR(255) NOT NULL"`
	Comment      sql.NullString `mysqlname:"Comment" mysqltype:"VARCHAR(255)"`
	Indexcomment string         `mysqlname:"Index_comment" mysqltype:"VARCHAR(255) NOT NULL"`
}

// Bind returns a slice of all columns for a Index to use it with Scan.
// Bind only one Index and copy it every row after calling Scan.
func (v *Index) Bind() (cols []interface{}) {
	return []interface{}{
		&v.Table,
		&v.Nonunique,
		&v.Keyname,
		&v.Seqinindex,
		&v.Columnname,
		&v.Collation,
		&v.Cardinality,
		&v.Subpart,
		&v.Packed,
		&v.Null,
		&v.Indextype,
		&v.Comment,
		&v.Indexcomment,
	}
}

// copy returns a shallow copy of the Index.
func (v *Index) copy() *Index {
	if v == nil {
		return nil
	}
	w := *v
	return &w
}

// Copy returns a shallow copy of the Index.
func (v *Index) Copy() Bindable {
	return v.copy()
}

// query provides sql metadata used to query for Index.
// It returns the SQL string and the arguments needed for prepared statements.
func (v *Index) query(schema, table string) (sql string, args []string) {
	return "SHOW INDEXES FROM `" + schema + "`.`" + table + "`", nil
}

// Query provides sql metadata used to query for Index.
// It returns the SQL string and the arguments needed for prepared statements.
// Query panics if location is not schema, table string.
func (v *Index) Query(location ...string) (sql string, args []string) {
	if len(location) != 2 {
		panic("Query must be called with arguments (schema, table string)")
	}
	return v.query(location[0], location[1])
}

type SessionStatus struct {
	Variablename string         `mysqlname:"Variable_name" mysqltype:"VARCHAR(255) NOT NULL"`
	Value        sql.NullString `mysqlname:"Value" mysqltype:"VARCHAR(255)"`
}

// Bind returns a slice of all columns for a SessionStatus to use it with Scan.
// Bind only one SessionStatus and copy it every row after calling Scan.
func (v *SessionStatus) Bind() (cols []interface{}) {
	return []interface{}{
		&v.Variablename,
		&v.Value,
	}
}

// copy returns a shallow copy of the SessionStatus.
func (v *SessionStatus) copy() *SessionStatus {
	if v == nil {
		return nil
	}
	w := *v
	return &w
}

// Copy returns a shallow copy of the SessionStatus.
func (v *SessionStatus) Copy() Bindable {
	return v.copy()
}

// Query provides sql metadata used to query for SessionStatus.
// It returns the SQL string and the arguments needed for prepared statements.
// Query panics if arguments are given.
func (v *SessionStatus) Query(location ...string) (sql string, args []string) {
	if len(location) != 0 {
		panic("Query must be called without arguments")
	}
	return "SHOW SESSION STATUS", nil
}

type SessionVariables struct {
	Variablename string         `mysqlname:"Variable_name" mysqltype:"VARCHAR(255) NOT NULL"`
	Value        sql.NullString `mysqlname:"Value" mysqltype:"VARCHAR(255)"`
}

// Bind returns a slice of all columns for a SessionVariables to use it with Scan.
// Bind only one SessionVariables and copy it every row after calling Scan.
func (v *SessionVariables) Bind() (cols []interface{}) {
	return []interface{}{
		&v.Variablename,
		&v.Value,
	}
}

// copy returns a shallow copy of the SessionVariables.
func (v *SessionVariables) copy() *SessionVariables {
	if v == nil {
		return nil
	}
	w := *v
	return &w
}

// Copy returns a shallow copy of the SessionVariables.
func (v *SessionVariables) Copy() Bindable {
	return v.copy()
}

// Query provides sql metadata used to query for SessionVariables.
// It returns the SQL string and the arguments needed for prepared statements.
// Query panics if arguments are given.
func (v *SessionVariables) Query(location ...string) (sql string, args []string) {
	if len(location) != 0 {
		panic("Query must be called without arguments")
	}
	return "SHOW SESSION VARIABLES", nil
}

type GlobalStatus struct {
	Variablename string         `mysqlname:"Variable_name" mysqltype:"VARCHAR(255) NOT NULL"`
	Value        sql.NullString `mysqlname:"Value" mysqltype:"VARCHAR(255)"`
}

// Bind returns a slice of all columns for a GlobalStatus to use it with Scan.
// Bind only one GlobalStatus and copy it every row after calling Scan.
func (v *GlobalStatus) Bind() (cols []interface{}) {
	return []interface{}{
		&v.Variablename,
		&v.Value,
	}
}

// copy returns a shallow copy of the GlobalStatus.
func (v *GlobalStatus) copy() *GlobalStatus {
	if v == nil {
		return nil
	}
	w := *v
	return &w
}

// Copy returns a shallow copy of the GlobalStatus.
func (v *GlobalStatus) Copy() Bindable {
	return v.copy()
}

// Query provides sql metadata used to query for GlobalStatus.
// It returns the SQL string and the arguments needed for prepared statements.
// Query panics if arguments are given.
func (v *GlobalStatus) Query(location ...string) (sql string, args []string) {
	if len(location) != 0 {
		panic("Query must be called without arguments")
	}
	return "SHOW GLOBAL STATUS", nil
}

type GlobalVariables struct {
	Variablename string         `mysqlname:"Variable_name" mysqltype:"VARCHAR(255) NOT NULL"`
	Value        sql.NullString `mysqlname:"Value" mysqltype:"VARCHAR(255)"`
}

// Bind returns a slice of all columns for a GlobalVariables to use it with Scan.
// Bind only one GlobalVariables and copy it every row after calling Scan.
func (v *GlobalVariables) Bind() (cols []interface{}) {
	return []interface{}{
		&v.Variablename,
		&v.Value,
	}
}

// copy returns a shallow copy of the GlobalVariables.
func (v *GlobalVariables) copy() *GlobalVariables {
	if v == nil {
		return nil
	}
	w := *v
	return &w
}

// Copy returns a shallow copy of the GlobalVariables.
func (v *GlobalVariables) Copy() Bindable {
	return v.copy()
}

// Query provides sql metadata used to query for GlobalVariables.
// It returns the SQL string and the arguments needed for prepared statements.
// Query panics if arguments are given.
func (v *GlobalVariables) Query(location ...string) (sql string, args []string) {
	if len(location) != 0 {
		panic("Query must be called without arguments")
	}
	return "SHOW GLOBAL VARIABLES", nil
}

type Error struct {
	Level   string `mysqlname:"Level" mysqltype:"VARCHAR(255) NOT NULL"`
	Code    uint32 `mysqlname:"Code" mysqltype:"INT UNSIGNED NOT NULL"`
	Message string `mysqlname:"Message" mysqltype:"VARCHAR(255) NOT NULL"`
}

// Bind returns a slice of all columns for a Error to use it with Scan.
// Bind only one Error and copy it every row after calling Scan.
func (v *Error) Bind() (cols []interface{}) {
	return []interface{}{
		&v.Level,
		&v.Code,
		&v.Message,
	}
}

// copy returns a shallow copy of the Error.
func (v *Error) copy() *Error {
	if v == nil {
		return nil
	}
	w := *v
	return &w
}

// Copy returns a shallow copy of the Error.
func (v *Error) Copy() Bindable {
	return v.copy()
}

// Query provides sql metadata used to query for Error.
// It returns the SQL string and the arguments needed for prepared statements.
// Query panics if arguments are given.
func (v *Error) Query(location ...string) (sql string, args []string) {
	if len(location) != 0 {
		panic("Query must be called without arguments")
	}
	return "SHOW ERRORS", nil
}

type Warning struct {
	Level   string `mysqlname:"Level" mysqltype:"VARCHAR(255) NOT NULL"`
	Code    uint32 `mysqlname:"Code" mysqltype:"INT UNSIGNED NOT NULL"`
	Message string `mysqlname:"Message" mysqltype:"VARCHAR(255) NOT NULL"`
}

// Bind returns a slice of all columns for a Warning to use it with Scan.
// Bind only one Warning and copy it every row after calling Scan.
func (v *Warning) Bind() (cols []interface{}) {
	return []interface{}{
		&v.Level,
		&v.Code,
		&v.Message,
	}
}

// copy returns a shallow copy of the Warning.
func (v *Warning) copy() *Warning {
	if v == nil {
		return nil
	}
	w := *v
	return &w
}

// Copy returns a shallow copy of the Warning.
func (v *Warning) Copy() Bindable {
	return v.copy()
}

// Query provides sql metadata used to query for Warning.
// It returns the SQL string and the arguments needed for prepared statements.
// Query panics if arguments are given.
func (v *Warning) Query(location ...string) (sql string, args []string) {
	if len(location) != 0 {
		panic("Query must be called without arguments")
	}
	return "SHOW WARNINGS", nil
}

type InnodbMutex struct {
	Type   string `mysqlname:"Type" mysqltype:"VARCHAR(255) NOT NULL"`
	Name   string `mysqlname:"Name" mysqltype:"VARCHAR(255) NOT NULL"`
	Status string `mysqlname:"Status" mysqltype:"VARCHAR(255) NOT NULL"`
}

// Bind returns a slice of all columns for a InnodbMutex to use it with Scan.
// Bind only one InnodbMutex and copy it every row after calling Scan.
func (v *InnodbMutex) Bind() (cols []interface{}) {
	return []interface{}{
		&v.Type,
		&v.Name,
		&v.Status,
	}
}

// copy returns a shallow copy of the InnodbMutex.
func (v *InnodbMutex) copy() *InnodbMutex {
	if v == nil {
		return nil
	}
	w := *v
	return &w
}

// Copy returns a shallow copy of the InnodbMutex.
func (v *InnodbMutex) Copy() Bindable {
	return v.copy()
}

// Query provides sql metadata used to query for InnodbMutex.
// It returns the SQL string and the arguments needed for prepared statements.
// Query panics if arguments are given.
func (v *InnodbMutex) Query(location ...string) (sql string, args []string) {
	if len(location) != 0 {
		panic("Query must be called without arguments")
	}
	return "SHOW ENGINE INNODB MUTEX", nil
}

type InnodbStatus struct {
	Type   string `mysqlname:"Type" mysqltype:"VARCHAR(255) NOT NULL"`
	Name   string `mysqlname:"Name" mysqltype:"VARCHAR(255) NOT NULL"`
	Status string `mysqlname:"Status" mysqltype:"VARCHAR(255) NOT NULL"`
}

// Bind returns a slice of all columns for a InnodbStatus to use it with Scan.
// Bind only one InnodbStatus and copy it every row after calling Scan.
func (v *InnodbStatus) Bind() (cols []interface{}) {
	return []interface{}{
		&v.Type,
		&v.Name,
		&v.Status,
	}
}

// copy returns a shallow copy of the InnodbStatus.
func (v *InnodbStatus) copy() *InnodbStatus {
	if v == nil {
		return nil
	}
	w := *v
	return &w
}

// Copy returns a shallow copy of the InnodbStatus.
func (v *InnodbStatus) Copy() Bindable {
	return v.copy()
}

// Query provides sql metadata used to query for InnodbStatus.
// It returns the SQL string and the arguments needed for prepared statements.
// Query panics if arguments are given.
func (v *InnodbStatus) Query(location ...string) (sql string, args []string) {
	if len(location) != 0 {
		panic("Query must be called without arguments")
	}
	return "SHOW ENGINE INNODB STATUS", nil
}

type OpenTable struct {
	Database   string `mysqlname:"Database" mysqltype:"VARCHAR(255) NOT NULL"`
	Table      string `mysqlname:"Table" mysqltype:"VARCHAR(255) NOT NULL"`
	Inuse      int64  `mysqlname:"In_use" mysqltype:"BIGINT NOT NULL"`
	Namelocked int64  `mysqlname:"Name_locked" mysqltype:"BIGINT NOT NULL"`
}

// Bind returns a slice of all columns for a OpenTable to use it with Scan.
// Bind only one OpenTable and copy it every row after calling Scan.
func (v *OpenTable) Bind() (cols []interface{}) {
	return []interface{}{
		&v.Database,
		&v.Table,
		&v.Inuse,
		&v.Namelocked,
	}
}

// copy returns a shallow copy of the OpenTable.
func (v *OpenTable) copy() *OpenTable {
	if v == nil {
		return nil
	}
	w := *v
	return &w
}

// Copy returns a shallow copy of the OpenTable.
func (v *OpenTable) Copy() Bindable {
	return v.copy()
}

// Query provides sql metadata used to query for OpenTable.
// It returns the SQL string and the arguments needed for prepared statements.
// Query panics if arguments are given.
func (v *OpenTable) Query(location ...string) (sql string, args []string) {
	if len(location) != 0 {
		panic("Query must be called without arguments")
	}
	return "SHOW OPEN TABLES", nil
}

type Process struct {
	Id      int64          `mysqlname:"Id" mysqltype:"BIGINT NOT NULL"`
	User    string         `mysqlname:"User" mysqltype:"VARCHAR(255) NOT NULL"`
	Host    string         `mysqlname:"Host" mysqltype:"VARCHAR(255) NOT NULL"`
	Db      sql.NullString `mysqlname:"db" mysqltype:"VARCHAR(255)"`
	Command string         `mysqlname:"Command" mysqltype:"VARCHAR(255) NOT NULL"`
	Time    int32          `mysqlname:"Time" mysqltype:"INT NOT NULL"`
	State   sql.NullString `mysqlname:"State" mysqltype:"VARCHAR(255)"`
	Info    sql.NullString `mysqlname:"Info" mysqltype:"VARCHAR(255)"`
}

// Bind returns a slice of all columns for a Process to use it with Scan.
// Bind only one Process and copy it every row after calling Scan.
func (v *Process) Bind() (cols []interface{}) {
	return []interface{}{
		&v.Id,
		&v.User,
		&v.Host,
		&v.Db,
		&v.Command,
		&v.Time,
		&v.State,
		&v.Info,
	}
}

// copy returns a shallow copy of the Process.
func (v *Process) copy() *Process {
	if v == nil {
		return nil
	}
	w := *v
	return &w
}

// Copy returns a shallow copy of the Process.
func (v *Process) Copy() Bindable {
	return v.copy()
}

// Query provides sql metadata used to query for Process.
// It returns the SQL string and the arguments needed for prepared statements.
// Query panics if arguments are given.
func (v *Process) Query(location ...string) (sql string, args []string) {
	if len(location) != 0 {
		panic("Query must be called without arguments")
	}
	return "SHOW PROCESSLIST", nil
}

type MasterStatus struct {
	File            string `mysqlname:"File" mysqltype:"VARCHAR(255) NOT NULL"`
	Position        uint64 `mysqlname:"Position" mysqltype:"BIGINT UNSIGNED NOT NULL"`
	BinlogDoDB      string `mysqlname:"Binlog_Do_DB" mysqltype:"VARCHAR(255) NOT NULL"`
	BinlogIgnoreDB  string `mysqlname:"Binlog_Ignore_DB" mysqltype:"VARCHAR(255) NOT NULL"`
	ExecutedGtidSet string `mysqlname:"Executed_Gtid_Set" mysqltype:"VARCHAR(255) NOT NULL"`
}

// Bind returns a slice of all columns for a MasterStatus to use it with Scan.
// Bind only one MasterStatus and copy it every row after calling Scan.
func (v *MasterStatus) Bind() (cols []interface{}) {
	return []interface{}{
		&v.File,
		&v.Position,
		&v.BinlogDoDB,
		&v.BinlogIgnoreDB,
		&v.ExecutedGtidSet,
	}
}

// copy returns a shallow copy of the MasterStatus.
func (v *MasterStatus) copy() *MasterStatus {
	if v == nil {
		return nil
	}
	w := *v
	return &w
}

// Copy returns a shallow copy of the MasterStatus.
func (v *MasterStatus) Copy() Bindable {
	return v.copy()
}

// Query provides sql metadata used to query for MasterStatus.
// It returns the SQL string and the arguments needed for prepared statements.
// Query panics if arguments are given.
func (v *MasterStatus) Query(location ...string) (sql string, args []string) {
	if len(location) != 0 {
		panic("Query must be called without arguments")
	}
	return "SHOW MASTER STATUS", nil
}

type SlaveHost struct {
	Serverid  uint32 `mysqlname:"Server_id" mysqltype:"INT UNSIGNED NOT NULL"`
	Host      string `mysqlname:"Host" mysqltype:"VARCHAR(255) NOT NULL"`
	Port      uint32 `mysqlname:"Port" mysqltype:"INT UNSIGNED NOT NULL"`
	Masterid  uint32 `mysqlname:"Master_id" mysqltype:"INT UNSIGNED NOT NULL"`
	SlaveUUID string `mysqlname:"Slave_UUID" mysqltype:"VARCHAR(255) NOT NULL"`
}

// Bind returns a slice of all columns for a SlaveHost to use it with Scan.
// Bind only one SlaveHost and copy it every row after calling Scan.
func (v *SlaveHost) Bind() (cols []interface{}) {
	return []interface{}{
		&v.Serverid,
		&v.Host,
		&v.Port,
		&v.Masterid,
		&v.SlaveUUID,
	}
}

// copy returns a shallow copy of the SlaveHost.
func (v *SlaveHost) copy() *SlaveHost {
	if v == nil {
		return nil
	}
	w := *v
	return &w
}

// Copy returns a shallow copy of the SlaveHost.
func (v *SlaveHost) Copy() Bindable {
	return v.copy()
}

// Query provides sql metadata used to query for SlaveHost.
// It returns the SQL string and the arguments needed for prepared statements.
// Query panics if arguments are given.
func (v *SlaveHost) Query(location ...string) (sql string, args []string) {
	if len(location) != 0 {
		panic("Query must be called without arguments")
	}
	return "SHOW SLAVE HOSTS", nil
}

type SlaveStatus struct {
	SlaveIOState              string `mysqlname:"Slave_IO_State" mysqltype:"VARCHAR(255) NOT NULL"`
	MasterHost                string `mysqlname:"Master_Host" mysqltype:"VARCHAR(255) NOT NULL"`
	MasterUser                string `mysqlname:"Master_User" mysqltype:"VARCHAR(255) NOT NULL"`
	MasterPort                uint32 `mysqlname:"Master_Port" mysqltype:"INT UNSIGNED NOT NULL"`
	ConnectRetry              uint32 `mysqlname:"Connect_Retry" mysqltype:"INT UNSIGNED NOT NULL"`
	MasterLogFile             string `mysqlname:"Master_Log_File" mysqltype:"VARCHAR(255) NOT NULL"`
	ReadMasterLogPos          uint64 `mysqlname:"Read_Master_Log_Pos" mysqltype:"BIGINT UNSIGNED NOT NULL"`
	RelayLogFile              string `mysqlname:"Relay_Log_File" mysqltype:"VARCHAR(255) NOT NULL"`
	RelayLogPos               uint64 `mysqlname:"Relay_Log_Pos" mysqltype:"BIGINT UNSIGNED NOT NULL"`
	RelayMasterLogFile        string `mysqlname:"Relay_Master_Log_File" mysqltype:"VARCHAR(255) NOT NULL"`
	SlaveIORunning            string `mysqlname:"Slave_IO_Running" mysqltype:"VARCHAR(255) NOT NULL"`
	SlaveSQLRunning           string `mysqlname:"Slave_SQL_Running" mysqltype:"VARCHAR(255) NOT NULL"`
	ReplicateDoDB             string `mysqlname:"Replicate_Do_DB" mysqltype:"VARCHAR(255) NOT NULL"`
	ReplicateIgnoreDB         string `mysqlname:"Replicate_Ignore_DB" mysqltype:"VARCHAR(255) NOT NULL"`
	ReplicateDoTable          string `mysqlname:"Replicate_Do_Table" mysqltype:"VARCHAR(255) NOT NULL"`
	ReplicateIgnoreTable      string `mysqlname:"Replicate_Ignore_Table" mysqltype:"VARCHAR(255) NOT NULL"`
	ReplicateWildDoTable      string `mysqlname:"Replicate_Wild_Do_Table" mysqltype:"VARCHAR(255) NOT NULL"`
	ReplicateWildIgnoreTable  string `mysqlname:"Replicate_Wild_Ignore_Table" mysqltype:"VARCHAR(255) NOT NULL"`
	LastErrno                 uint32 `mysqlname:"Last_Errno" mysqltype:"INT UNSIGNED NOT NULL"`
	LastError                 string `mysqlname:"Last_Error" mysqltype:"VARCHAR(255) NOT NULL"`
	SkipCounter               uint32 `mysqlname:"Skip_Counter" mysqltype:"INT UNSIGNED NOT NULL"`
	ExecMasterLogPos          uint64 `mysqlname:"Exec_Master_Log_Pos" mysqltype:"BIGINT UNSIGNED NOT NULL"`
	RelayLogSpace             uint64 `mysqlname:"Relay_Log_Space" mysqltype:"BIGINT UNSIGNED NOT NULL"`
	UntilCondition            string `mysqlname:"Until_Condition" mysqltype:"VARCHAR(255) NOT NULL"`
	UntilLogFile              string `mysqlname:"Until_Log_File" mysqltype:"VARCHAR(255) NOT NULL"`
	UntilLogPos               uint64 `mysqlname:"Until_Log_Pos" mysqltype:"BIGINT UNSIGNED NOT NULL"`
	MasterSSLAllowed          string `mysqlname:"Master_SSL_Allowed" mysqltype:"VARCHAR(255) NOT NULL"`
	MasterSSLCAFile           string `mysqlname:"Master_SSL_CA_File" mysqltype:"VARCHAR(255) NOT NULL"`
	MasterSSLCAPath           string `mysqlname:"Master_SSL_CA_Path" mysqltype:"VARCHAR(255) NOT NULL"`
	MasterSSLCert             string `mysqlname:"Master_SSL_Cert" mysqltype:"VARCHAR(255) NOT NULL"`
	MasterSSLCipher           string `mysqlname:"Master_SSL_Cipher" mysqltype:"VARCHAR(255) NOT NULL"`
	MasterSSLKey              string `mysqlname:"Master_SSL_Key" mysqltype:"VARCHAR(255) NOT NULL"`
	SecondsBehindMaster       uint64 `mysqlname:"Seconds_Behind_Master" mysqltype:"BIGINT UNSIGNED NOT NULL"`
	MasterSSLVerifyServerCert string `mysqlname:"Master_SSL_Verify_Server_Cert" mysqltype:"VARCHAR(255) NOT NULL"`
	LastIOErrno               uint32 `mysqlname:"Last_IO_Errno" mysqltype:"INT UNSIGNED NOT NULL"`
	LastIOError               string `mysqlname:"Last_IO_Error" mysqltype:"VARCHAR(255) NOT NULL"`
	LastSQLErrno              uint32 `mysqlname:"Last_SQL_Errno" mysqltype:"INT UNSIGNED NOT NULL"`
	LastSQLError              string `mysqlname:"Last_SQL_Error" mysqltype:"VARCHAR(255) NOT NULL"`
	ReplicateIgnoreServerIds  string `mysqlname:"Replicate_Ignore_Server_Ids" mysqltype:"VARCHAR(255) NOT NULL"`
	MasterServerId            uint32 `mysqlname:"Master_Server_Id" mysqltype:"INT UNSIGNED NOT NULL"`
	MasterUUID                string `mysqlname:"Master_UUID" mysqltype:"VARCHAR(255) NOT NULL"`
	MasterInfoFile            string `mysqlname:"Master_Info_File" mysqltype:"VARCHAR(255) NOT NULL"`
	SQLDelay                  uint32 `mysqlname:"SQL_Delay" mysqltype:"INT UNSIGNED NOT NULL"`
	SQLRemainingDelay         uint32 `mysqlname:"SQL_Remaining_Delay" mysqltype:"INT UNSIGNED NOT NULL"`
	SlaveSQLRunningState      string `mysqlname:"Slave_SQL_Running_State" mysqltype:"VARCHAR(255) NOT NULL"`
	MasterRetryCount          uint64 `mysqlname:"Master_Retry_Count" mysqltype:"BIGINT UNSIGNED NOT NULL"`
	MasterBind                string `mysqlname:"Master_Bind" mysqltype:"VARCHAR(255) NOT NULL"`
	LastIOErrorTimestamp      string `mysqlname:"Last_IO_Error_Timestamp" mysqltype:"VARCHAR(255) NOT NULL"`
	LastSQLErrorTimestamp     string `mysqlname:"Last_SQL_Error_Timestamp" mysqltype:"VARCHAR(255) NOT NULL"`
	MasterSSLCrl              string `mysqlname:"Master_SSL_Crl" mysqltype:"VARCHAR(255) NOT NULL"`
	MasterSSLCrlpath          string `mysqlname:"Master_SSL_Crlpath" mysqltype:"VARCHAR(255) NOT NULL"`
	RetrievedGtidSet          string `mysqlname:"Retrieved_Gtid_Set" mysqltype:"VARCHAR(255) NOT NULL"`
	ExecutedGtidSet           string `mysqlname:"Executed_Gtid_Set" mysqltype:"VARCHAR(255) NOT NULL"`
	AutoPosition              uint32 `mysqlname:"Auto_Position" mysqltype:"INT UNSIGNED NOT NULL"`
}

// Bind returns a slice of all columns for a SlaveStatus to use it with Scan.
// Bind only one SlaveStatus and copy it every row after calling Scan.
func (v *SlaveStatus) Bind() (cols []interface{}) {
	return []interface{}{
		&v.SlaveIOState,
		&v.MasterHost,
		&v.MasterUser,
		&v.MasterPort,
		&v.ConnectRetry,
		&v.MasterLogFile,
		&v.ReadMasterLogPos,
		&v.RelayLogFile,
		&v.RelayLogPos,
		&v.RelayMasterLogFile,
		&v.SlaveIORunning,
		&v.SlaveSQLRunning,
		&v.ReplicateDoDB,
		&v.ReplicateIgnoreDB,
		&v.ReplicateDoTable,
		&v.ReplicateIgnoreTable,
		&v.ReplicateWildDoTable,
		&v.ReplicateWildIgnoreTable,
		&v.LastErrno,
		&v.LastError,
		&v.SkipCounter,
		&v.ExecMasterLogPos,
		&v.RelayLogSpace,
		&v.UntilCondition,
		&v.UntilLogFile,
		&v.UntilLogPos,
		&v.MasterSSLAllowed,
		&v.MasterSSLCAFile,
		&v.MasterSSLCAPath,
		&v.MasterSSLCert,
		&v.MasterSSLCipher,
		&v.MasterSSLKey,
		&v.SecondsBehindMaster,
		&v.MasterSSLVerifyServerCert,
		&v.LastIOErrno,
		&v.LastIOError,
		&v.LastSQLErrno,
		&v.LastSQLError,
		&v.ReplicateIgnoreServerIds,
		&v.MasterServerId,
		&v.MasterUUID,
		&v.MasterInfoFile,
		&v.SQLDelay,
		&v.SQLRemainingDelay,
		&v.SlaveSQLRunningState,
		&v.MasterRetryCount,
		&v.MasterBind,
		&v.LastIOErrorTimestamp,
		&v.LastSQLErrorTimestamp,
		&v.MasterSSLCrl,
		&v.MasterSSLCrlpath,
		&v.RetrievedGtidSet,
		&v.ExecutedGtidSet,
		&v.AutoPosition,
	}
}

// copy returns a shallow copy of the SlaveStatus.
func (v *SlaveStatus) copy() *SlaveStatus {
	if v == nil {
		return nil
	}
	w := *v
	return &w
}

// Copy returns a shallow copy of the SlaveStatus.
func (v *SlaveStatus) Copy() Bindable {
	return v.copy()
}

// Query provides sql metadata used to query for SlaveStatus.
// It returns the SQL string and the arguments needed for prepared statements.
// Query panics if arguments are given.
func (v *SlaveStatus) Query(location ...string) (sql string, args []string) {
	if len(location) != 0 {
		panic("Query must be called without arguments")
	}
	return "SHOW SLAVE STATUS", nil
}
