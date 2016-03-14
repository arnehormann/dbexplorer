// this file is generated, do not edit it!

package dbexplorer

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
	Query(location...string) (sql string, args []string)
}


type CharacterSet struct {
	Charset string `mysql:"Charset,VARCHAR(255) NOT NULL" json:"charset,omitempty" xml:"charset,omitempty"`
	Description string `mysql:"Description,VARCHAR(255) NOT NULL" json:"description,omitempty" xml:"description,omitempty"`
	Defaultcollation string `mysql:"Default collation,VARCHAR(255) NOT NULL" json:"defaultcollation,omitempty" xml:"defaultcollation,omitempty"`
	Maxlen int64 `mysql:"Maxlen,BIGINT NOT NULL" json:"maxlen,omitempty" xml:"maxlen,omitempty"`
}

// Bind returns a slice of all columns for a CharacterSet to use it with Scan.
// Bind only one CharacterSet, reuse the []interface{} and store a copy after calling Scan.
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
func (v *CharacterSet) Query(location...string) (sql string, args []string) {
	if len(location) != 0 {
		panic("Query must be called without arguments")
	}
	return "SHOW CHARACTER SET", nil
}


type Collation struct {
	Collation string `mysql:"Collation,VARCHAR(255) NOT NULL" json:"collation,omitempty" xml:"collation,omitempty"`
	Charset string `mysql:"Charset,VARCHAR(255) NOT NULL" json:"charset,omitempty" xml:"charset,omitempty"`
	Id int64 `mysql:"Id,BIGINT NOT NULL" json:"id,omitempty" xml:"id,omitempty"`
	Default string `mysql:"Default,VARCHAR(255) NOT NULL" json:"default,omitempty" xml:"default,omitempty"`
	Compiled string `mysql:"Compiled,VARCHAR(255) NOT NULL" json:"compiled,omitempty" xml:"compiled,omitempty"`
	Sortlen int64 `mysql:"Sortlen,BIGINT NOT NULL" json:"sortlen,omitempty" xml:"sortlen,omitempty"`
}

// Bind returns a slice of all columns for a Collation to use it with Scan.
// Bind only one Collation, reuse the []interface{} and store a copy after calling Scan.
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
func (v *Collation) Query(location...string) (sql string, args []string) {
	if len(location) != 0 {
		panic("Query must be called without arguments")
	}
	return "SHOW COLLATION", nil
}


type Engine struct {
	Engine string `mysql:"Engine,VARCHAR(255) NOT NULL" json:"engine,omitempty" xml:"engine,omitempty"`
	Support string `mysql:"Support,VARCHAR(255) NOT NULL" json:"support,omitempty" xml:"support,omitempty"`
	Comment string `mysql:"Comment,VARCHAR(255) NOT NULL" json:"comment,omitempty" xml:"comment,omitempty"`
	Transactions sql.NullString `mysql:"Transactions,VARCHAR(255)" json:"transactions,omitempty" xml:"transactions,omitempty"`
	XA sql.NullString `mysql:"XA,VARCHAR(255)" json:"xa,omitempty" xml:"xa,omitempty"`
	Savepoints sql.NullString `mysql:"Savepoints,VARCHAR(255)" json:"savepoints,omitempty" xml:"savepoints,omitempty"`
}

// Bind returns a slice of all columns for a Engine to use it with Scan.
// Bind only one Engine, reuse the []interface{} and store a copy after calling Scan.
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
func (v *Engine) Query(location...string) (sql string, args []string) {
	if len(location) != 0 {
		panic("Query must be called without arguments")
	}
	return "SHOW ENGINES", nil
}


type Privilege struct {
	Privilege string `mysql:"Privilege,VARCHAR(255) NOT NULL" json:"privilege,omitempty" xml:"privilege,omitempty"`
	Context string `mysql:"Context,VARCHAR(255) NOT NULL" json:"context,omitempty" xml:"context,omitempty"`
	Comment string `mysql:"Comment,VARCHAR(255) NOT NULL" json:"comment,omitempty" xml:"comment,omitempty"`
}

// Bind returns a slice of all columns for a Privilege to use it with Scan.
// Bind only one Privilege, reuse the []interface{} and store a copy after calling Scan.
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
func (v *Privilege) Query(location...string) (sql string, args []string) {
	if len(location) != 0 {
		panic("Query must be called without arguments")
	}
	return "SHOW PRIVILEGES", nil
}


type Grant struct {
	Grant string `mysql:"#Grant#,VARCHAR(255) NOT NULL" json:"grant,omitempty" xml:"grant,omitempty"`
}

// Bind returns a slice of all columns for a Grant to use it with Scan.
// Bind only one Grant, reuse the []interface{} and store a copy after calling Scan.
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
func (v *Grant) Query(location...string) (sql string, args []string) {
	if len(location) != 0 {
		panic("Query must be called without arguments")
	}
	return "SHOW GRANTS FOR '?'@'?'", []string{"user", "host"}
}


type Schema struct {
	Database string `mysql:"Database,VARCHAR(255) NOT NULL" json:"database,omitempty" xml:"database,omitempty"`
}

// Bind returns a slice of all columns for a Schema to use it with Scan.
// Bind only one Schema, reuse the []interface{} and store a copy after calling Scan.
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
func (v *Schema) Query(location...string) (sql string, args []string) {
	if len(location) != 0 {
		panic("Query must be called without arguments")
	}
	return "SHOW SCHEMAS", nil
}


type TableStatus struct {
	Name string `mysql:"Name,VARCHAR(255) NOT NULL" json:"name,omitempty" xml:"name,omitempty"`
	Engine sql.NullString `mysql:"Engine,VARCHAR(255)" json:"engine,omitempty" xml:"engine,omitempty"`
	Version sql.NullInt64 `mysql:"Version,BIGINT UNSIGNED" json:"version,omitempty" xml:"version,omitempty"`
	Rowformat sql.NullString `mysql:"Row_format,VARCHAR(255)" json:"rowformat,omitempty" xml:"rowformat,omitempty"`
	Rows sql.NullInt64 `mysql:"Rows,BIGINT UNSIGNED" json:"rows,omitempty" xml:"rows,omitempty"`
	Avgrowlength sql.NullInt64 `mysql:"Avg_row_length,BIGINT UNSIGNED" json:"avgrowlength,omitempty" xml:"avgrowlength,omitempty"`
	Datalength sql.NullInt64 `mysql:"Data_length,BIGINT UNSIGNED" json:"datalength,omitempty" xml:"datalength,omitempty"`
	Maxdatalength sql.NullInt64 `mysql:"Max_data_length,BIGINT UNSIGNED" json:"maxdatalength,omitempty" xml:"maxdatalength,omitempty"`
	Indexlength sql.NullInt64 `mysql:"Index_length,BIGINT UNSIGNED" json:"indexlength,omitempty" xml:"indexlength,omitempty"`
	Datafree sql.NullInt64 `mysql:"Data_free,BIGINT UNSIGNED" json:"datafree,omitempty" xml:"datafree,omitempty"`
	Autoincrement sql.NullInt64 `mysql:"Auto_increment,BIGINT UNSIGNED" json:"autoincrement,omitempty" xml:"autoincrement,omitempty"`
	Createtime mysql.NullTime `mysql:"Create_time,DATETIME" json:"createtime,omitempty" xml:"createtime,omitempty"`
	Updatetime mysql.NullTime `mysql:"Update_time,DATETIME" json:"updatetime,omitempty" xml:"updatetime,omitempty"`
	Checktime mysql.NullTime `mysql:"Check_time,DATETIME" json:"checktime,omitempty" xml:"checktime,omitempty"`
	Collation sql.NullString `mysql:"Collation,VARCHAR(255)" json:"collation,omitempty" xml:"collation,omitempty"`
	Checksum sql.NullInt64 `mysql:"Checksum,BIGINT UNSIGNED" json:"checksum,omitempty" xml:"checksum,omitempty"`
	Createoptions sql.NullString `mysql:"Create_options,VARCHAR(255)" json:"createoptions,omitempty" xml:"createoptions,omitempty"`
	Comment string `mysql:"Comment,VARCHAR(255) NOT NULL" json:"comment,omitempty" xml:"comment,omitempty"`
}

// Bind returns a slice of all columns for a TableStatus to use it with Scan.
// Bind only one TableStatus, reuse the []interface{} and store a copy after calling Scan.
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
func (v *TableStatus) Query(location...string) (sql string, args []string) {
	if len(location) != 1 {
		panic("Query must be called with arguments (schema string)")
	}
	return v.query(location[0])
}


type Column struct {
	Field string `mysql:"Field,VARCHAR(255) NOT NULL" json:"field,omitempty" xml:"field,omitempty"`
	Type []uint8 `mysql:"Type,BLOB NOT NULL" json:"type,omitempty" xml:"type,omitempty"`
	Null string `mysql:"Null,VARCHAR(255) NOT NULL" json:"null,omitempty" xml:"null,omitempty"`
	Key string `mysql:"Key,VARCHAR(255) NOT NULL" json:"key,omitempty" xml:"key,omitempty"`
	Default []uint8 `mysql:"Default,BLOB" json:"default,omitempty" xml:"default,omitempty"`
	Extra string `mysql:"Extra,VARCHAR(255) NOT NULL" json:"extra,omitempty" xml:"extra,omitempty"`
}

// Bind returns a slice of all columns for a Column to use it with Scan.
// Bind only one Column, reuse the []interface{} and store a copy after calling Scan.
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
func (v *Column) Query(location...string) (sql string, args []string) {
	if len(location) != 2 {
		panic("Query must be called with arguments (schema, table string)")
	}
	return v.query(location[0], location[1])
}


type Index struct {
	Table string `mysql:"Table,VARCHAR(255) NOT NULL" json:"table,omitempty" xml:"table,omitempty"`
	Nonunique int64 `mysql:"Non_unique,BIGINT NOT NULL" json:"nonunique,omitempty" xml:"nonunique,omitempty"`
	Keyname string `mysql:"Key_name,VARCHAR(255) NOT NULL" json:"keyname,omitempty" xml:"keyname,omitempty"`
	Seqinindex int64 `mysql:"Seq_in_index,BIGINT NOT NULL" json:"seqinindex,omitempty" xml:"seqinindex,omitempty"`
	Columnname string `mysql:"Column_name,VARCHAR(255) NOT NULL" json:"columnname,omitempty" xml:"columnname,omitempty"`
	Collation sql.NullString `mysql:"Collation,VARCHAR(255)" json:"collation,omitempty" xml:"collation,omitempty"`
	Cardinality sql.NullInt64 `mysql:"Cardinality,BIGINT" json:"cardinality,omitempty" xml:"cardinality,omitempty"`
	Subpart sql.NullInt64 `mysql:"Sub_part,BIGINT" json:"subpart,omitempty" xml:"subpart,omitempty"`
	Packed sql.NullString `mysql:"Packed,VARCHAR(255)" json:"packed,omitempty" xml:"packed,omitempty"`
	Null string `mysql:"Null,VARCHAR(255) NOT NULL" json:"null,omitempty" xml:"null,omitempty"`
	Indextype string `mysql:"Index_type,VARCHAR(255) NOT NULL" json:"indextype,omitempty" xml:"indextype,omitempty"`
	Comment sql.NullString `mysql:"Comment,VARCHAR(255)" json:"comment,omitempty" xml:"comment,omitempty"`
	Indexcomment string `mysql:"Index_comment,VARCHAR(255) NOT NULL" json:"indexcomment,omitempty" xml:"indexcomment,omitempty"`
}

// Bind returns a slice of all columns for a Index to use it with Scan.
// Bind only one Index, reuse the []interface{} and store a copy after calling Scan.
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
func (v *Index) Query(location...string) (sql string, args []string) {
	if len(location) != 2 {
		panic("Query must be called with arguments (schema, table string)")
	}
	return v.query(location[0], location[1])
}


type SessionStatus struct {
	Variablename string `mysql:"Variable_name,VARCHAR(255) NOT NULL" json:"variablename,omitempty" xml:"variablename,omitempty"`
	Value sql.NullString `mysql:"Value,VARCHAR(255)" json:"value,omitempty" xml:"value,omitempty"`
}

// Bind returns a slice of all columns for a SessionStatus to use it with Scan.
// Bind only one SessionStatus, reuse the []interface{} and store a copy after calling Scan.
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
func (v *SessionStatus) Query(location...string) (sql string, args []string) {
	if len(location) != 0 {
		panic("Query must be called without arguments")
	}
	return "SHOW SESSION STATUS", nil
}


type SessionVariables struct {
	Variablename string `mysql:"Variable_name,VARCHAR(255) NOT NULL" json:"variablename,omitempty" xml:"variablename,omitempty"`
	Value sql.NullString `mysql:"Value,VARCHAR(255)" json:"value,omitempty" xml:"value,omitempty"`
}

// Bind returns a slice of all columns for a SessionVariables to use it with Scan.
// Bind only one SessionVariables, reuse the []interface{} and store a copy after calling Scan.
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
func (v *SessionVariables) Query(location...string) (sql string, args []string) {
	if len(location) != 0 {
		panic("Query must be called without arguments")
	}
	return "SHOW SESSION VARIABLES", nil
}


type GlobalStatus struct {
	Variablename string `mysql:"Variable_name,VARCHAR(255) NOT NULL" json:"variablename,omitempty" xml:"variablename,omitempty"`
	Value sql.NullString `mysql:"Value,VARCHAR(255)" json:"value,omitempty" xml:"value,omitempty"`
}

// Bind returns a slice of all columns for a GlobalStatus to use it with Scan.
// Bind only one GlobalStatus, reuse the []interface{} and store a copy after calling Scan.
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
func (v *GlobalStatus) Query(location...string) (sql string, args []string) {
	if len(location) != 0 {
		panic("Query must be called without arguments")
	}
	return "SHOW GLOBAL STATUS", nil
}


type GlobalVariables struct {
	Variablename string `mysql:"Variable_name,VARCHAR(255) NOT NULL" json:"variablename,omitempty" xml:"variablename,omitempty"`
	Value sql.NullString `mysql:"Value,VARCHAR(255)" json:"value,omitempty" xml:"value,omitempty"`
}

// Bind returns a slice of all columns for a GlobalVariables to use it with Scan.
// Bind only one GlobalVariables, reuse the []interface{} and store a copy after calling Scan.
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
func (v *GlobalVariables) Query(location...string) (sql string, args []string) {
	if len(location) != 0 {
		panic("Query must be called without arguments")
	}
	return "SHOW GLOBAL VARIABLES", nil
}


type Error struct {
	Level string `mysql:"Level,VARCHAR(255) NOT NULL" json:"level,omitempty" xml:"level,omitempty"`
	Code uint32 `mysql:"Code,INT UNSIGNED NOT NULL" json:"code,omitempty" xml:"code,omitempty"`
	Message string `mysql:"Message,VARCHAR(255) NOT NULL" json:"message,omitempty" xml:"message,omitempty"`
}

// Bind returns a slice of all columns for a Error to use it with Scan.
// Bind only one Error, reuse the []interface{} and store a copy after calling Scan.
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
func (v *Error) Query(location...string) (sql string, args []string) {
	if len(location) != 0 {
		panic("Query must be called without arguments")
	}
	return "SHOW ERRORS", nil
}


type Warning struct {
	Level string `mysql:"Level,VARCHAR(255) NOT NULL" json:"level,omitempty" xml:"level,omitempty"`
	Code uint32 `mysql:"Code,INT UNSIGNED NOT NULL" json:"code,omitempty" xml:"code,omitempty"`
	Message string `mysql:"Message,VARCHAR(255) NOT NULL" json:"message,omitempty" xml:"message,omitempty"`
}

// Bind returns a slice of all columns for a Warning to use it with Scan.
// Bind only one Warning, reuse the []interface{} and store a copy after calling Scan.
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
func (v *Warning) Query(location...string) (sql string, args []string) {
	if len(location) != 0 {
		panic("Query must be called without arguments")
	}
	return "SHOW WARNINGS", nil
}


type InnodbMutex struct {
	Type string `mysql:"Type,VARCHAR(255) NOT NULL" json:"type,omitempty" xml:"type,omitempty"`
	Name string `mysql:"Name,VARCHAR(255) NOT NULL" json:"name,omitempty" xml:"name,omitempty"`
	Status string `mysql:"Status,VARCHAR(255) NOT NULL" json:"status,omitempty" xml:"status,omitempty"`
}

// Bind returns a slice of all columns for a InnodbMutex to use it with Scan.
// Bind only one InnodbMutex, reuse the []interface{} and store a copy after calling Scan.
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
func (v *InnodbMutex) Query(location...string) (sql string, args []string) {
	if len(location) != 0 {
		panic("Query must be called without arguments")
	}
	return "SHOW ENGINE INNODB MUTEX", nil
}


type InnodbStatus struct {
	Type string `mysql:"Type,VARCHAR(255) NOT NULL" json:"type,omitempty" xml:"type,omitempty"`
	Name string `mysql:"Name,VARCHAR(255) NOT NULL" json:"name,omitempty" xml:"name,omitempty"`
	Status string `mysql:"Status,VARCHAR(255) NOT NULL" json:"status,omitempty" xml:"status,omitempty"`
}

// Bind returns a slice of all columns for a InnodbStatus to use it with Scan.
// Bind only one InnodbStatus, reuse the []interface{} and store a copy after calling Scan.
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
func (v *InnodbStatus) Query(location...string) (sql string, args []string) {
	if len(location) != 0 {
		panic("Query must be called without arguments")
	}
	return "SHOW ENGINE INNODB STATUS", nil
}


type OpenTable struct {
	Database string `mysql:"Database,VARCHAR(255) NOT NULL" json:"database,omitempty" xml:"database,omitempty"`
	Table string `mysql:"Table,VARCHAR(255) NOT NULL" json:"table,omitempty" xml:"table,omitempty"`
	Inuse int64 `mysql:"In_use,BIGINT NOT NULL" json:"inuse,omitempty" xml:"inuse,omitempty"`
	Namelocked int64 `mysql:"Name_locked,BIGINT NOT NULL" json:"namelocked,omitempty" xml:"namelocked,omitempty"`
}

// Bind returns a slice of all columns for a OpenTable to use it with Scan.
// Bind only one OpenTable, reuse the []interface{} and store a copy after calling Scan.
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
func (v *OpenTable) Query(location...string) (sql string, args []string) {
	if len(location) != 0 {
		panic("Query must be called without arguments")
	}
	return "SHOW OPEN TABLES", nil
}


type Process struct {
	Id int64 `mysql:"Id,BIGINT NOT NULL" json:"id,omitempty" xml:"id,omitempty"`
	User string `mysql:"User,VARCHAR(255) NOT NULL" json:"user,omitempty" xml:"user,omitempty"`
	Host string `mysql:"Host,VARCHAR(255) NOT NULL" json:"host,omitempty" xml:"host,omitempty"`
	Db sql.NullString `mysql:"db,VARCHAR(255)" json:"db,omitempty" xml:"db,omitempty"`
	Command string `mysql:"Command,VARCHAR(255) NOT NULL" json:"command,omitempty" xml:"command,omitempty"`
	Time int32 `mysql:"Time,INT NOT NULL" json:"time,omitempty" xml:"time,omitempty"`
	State sql.NullString `mysql:"State,VARCHAR(255)" json:"state,omitempty" xml:"state,omitempty"`
	Info sql.NullString `mysql:"Info,VARCHAR(255)" json:"info,omitempty" xml:"info,omitempty"`
	Progress float64 `mysql:"Progress,DOUBLE NOT NULL" json:"progress,omitempty" xml:"progress,omitempty"`
}

// Bind returns a slice of all columns for a Process to use it with Scan.
// Bind only one Process, reuse the []interface{} and store a copy after calling Scan.
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
		&v.Progress,
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
func (v *Process) Query(location...string) (sql string, args []string) {
	if len(location) != 0 {
		panic("Query must be called without arguments")
	}
	return "SHOW PROCESSLIST", nil
}


type MasterStatus struct {
	File string `mysql:"File,VARCHAR(255) NOT NULL" json:"file,omitempty" xml:"file,omitempty"`
	Position uint64 `mysql:"Position,BIGINT UNSIGNED NOT NULL" json:"position,omitempty" xml:"position,omitempty"`
	BinlogDoDB string `mysql:"Binlog_Do_DB,VARCHAR(255) NOT NULL" json:"binlogdodb,omitempty" xml:"binlogdodb,omitempty"`
	BinlogIgnoreDB string `mysql:"Binlog_Ignore_DB,VARCHAR(255) NOT NULL" json:"binlogignoredb,omitempty" xml:"binlogignoredb,omitempty"`
}

// Bind returns a slice of all columns for a MasterStatus to use it with Scan.
// Bind only one MasterStatus, reuse the []interface{} and store a copy after calling Scan.
func (v *MasterStatus) Bind() (cols []interface{}) {
	return []interface{}{
		&v.File,
		&v.Position,
		&v.BinlogDoDB,
		&v.BinlogIgnoreDB,
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
func (v *MasterStatus) Query(location...string) (sql string, args []string) {
	if len(location) != 0 {
		panic("Query must be called without arguments")
	}
	return "SHOW MASTER STATUS", nil
}


type SlaveHost struct {
	Serverid uint32 `mysql:"Server_id,INT UNSIGNED NOT NULL" json:"serverid,omitempty" xml:"serverid,omitempty"`
	Host string `mysql:"Host,VARCHAR(255) NOT NULL" json:"host,omitempty" xml:"host,omitempty"`
	Port uint32 `mysql:"Port,INT UNSIGNED NOT NULL" json:"port,omitempty" xml:"port,omitempty"`
	Masterid uint32 `mysql:"Master_id,INT UNSIGNED NOT NULL" json:"masterid,omitempty" xml:"masterid,omitempty"`
}

// Bind returns a slice of all columns for a SlaveHost to use it with Scan.
// Bind only one SlaveHost, reuse the []interface{} and store a copy after calling Scan.
func (v *SlaveHost) Bind() (cols []interface{}) {
	return []interface{}{
		&v.Serverid,
		&v.Host,
		&v.Port,
		&v.Masterid,
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
func (v *SlaveHost) Query(location...string) (sql string, args []string) {
	if len(location) != 0 {
		panic("Query must be called without arguments")
	}
	return "SHOW SLAVE HOSTS", nil
}


type SlaveStatus struct {
	SlaveIOState string `mysql:"Slave_IO_State,VARCHAR(255) NOT NULL" json:"slaveiostate,omitempty" xml:"slaveiostate,omitempty"`
	MasterHost string `mysql:"Master_Host,VARCHAR(255) NOT NULL" json:"masterhost,omitempty" xml:"masterhost,omitempty"`
	MasterUser string `mysql:"Master_User,VARCHAR(255) NOT NULL" json:"masteruser,omitempty" xml:"masteruser,omitempty"`
	MasterPort uint32 `mysql:"Master_Port,INT UNSIGNED NOT NULL" json:"masterport,omitempty" xml:"masterport,omitempty"`
	ConnectRetry uint32 `mysql:"Connect_Retry,INT UNSIGNED NOT NULL" json:"connectretry,omitempty" xml:"connectretry,omitempty"`
	MasterLogFile string `mysql:"Master_Log_File,VARCHAR(255) NOT NULL" json:"masterlogfile,omitempty" xml:"masterlogfile,omitempty"`
	ReadMasterLogPos uint64 `mysql:"Read_Master_Log_Pos,BIGINT UNSIGNED NOT NULL" json:"readmasterlogpos,omitempty" xml:"readmasterlogpos,omitempty"`
	RelayLogFile string `mysql:"Relay_Log_File,VARCHAR(255) NOT NULL" json:"relaylogfile,omitempty" xml:"relaylogfile,omitempty"`
	RelayLogPos uint64 `mysql:"Relay_Log_Pos,BIGINT UNSIGNED NOT NULL" json:"relaylogpos,omitempty" xml:"relaylogpos,omitempty"`
	RelayMasterLogFile string `mysql:"Relay_Master_Log_File,VARCHAR(255) NOT NULL" json:"relaymasterlogfile,omitempty" xml:"relaymasterlogfile,omitempty"`
	SlaveIORunning string `mysql:"Slave_IO_Running,VARCHAR(255) NOT NULL" json:"slaveiorunning,omitempty" xml:"slaveiorunning,omitempty"`
	SlaveSQLRunning string `mysql:"Slave_SQL_Running,VARCHAR(255) NOT NULL" json:"slavesqlrunning,omitempty" xml:"slavesqlrunning,omitempty"`
	ReplicateDoDB string `mysql:"Replicate_Do_DB,VARCHAR(255) NOT NULL" json:"replicatedodb,omitempty" xml:"replicatedodb,omitempty"`
	ReplicateIgnoreDB string `mysql:"Replicate_Ignore_DB,VARCHAR(255) NOT NULL" json:"replicateignoredb,omitempty" xml:"replicateignoredb,omitempty"`
	ReplicateDoTable string `mysql:"Replicate_Do_Table,VARCHAR(255) NOT NULL" json:"replicatedotable,omitempty" xml:"replicatedotable,omitempty"`
	ReplicateIgnoreTable string `mysql:"Replicate_Ignore_Table,VARCHAR(255) NOT NULL" json:"replicateignoretable,omitempty" xml:"replicateignoretable,omitempty"`
	ReplicateWildDoTable string `mysql:"Replicate_Wild_Do_Table,VARCHAR(255) NOT NULL" json:"replicatewilddotable,omitempty" xml:"replicatewilddotable,omitempty"`
	ReplicateWildIgnoreTable string `mysql:"Replicate_Wild_Ignore_Table,VARCHAR(255) NOT NULL" json:"replicatewildignoretable,omitempty" xml:"replicatewildignoretable,omitempty"`
	LastErrno uint32 `mysql:"Last_Errno,INT UNSIGNED NOT NULL" json:"lasterrno,omitempty" xml:"lasterrno,omitempty"`
	LastError string `mysql:"Last_Error,VARCHAR(255) NOT NULL" json:"lasterror,omitempty" xml:"lasterror,omitempty"`
	SkipCounter uint32 `mysql:"Skip_Counter,INT UNSIGNED NOT NULL" json:"skipcounter,omitempty" xml:"skipcounter,omitempty"`
	ExecMasterLogPos uint64 `mysql:"Exec_Master_Log_Pos,BIGINT UNSIGNED NOT NULL" json:"execmasterlogpos,omitempty" xml:"execmasterlogpos,omitempty"`
	RelayLogSpace uint64 `mysql:"Relay_Log_Space,BIGINT UNSIGNED NOT NULL" json:"relaylogspace,omitempty" xml:"relaylogspace,omitempty"`
	UntilCondition string `mysql:"Until_Condition,VARCHAR(255) NOT NULL" json:"untilcondition,omitempty" xml:"untilcondition,omitempty"`
	UntilLogFile string `mysql:"Until_Log_File,VARCHAR(255) NOT NULL" json:"untillogfile,omitempty" xml:"untillogfile,omitempty"`
	UntilLogPos uint64 `mysql:"Until_Log_Pos,BIGINT UNSIGNED NOT NULL" json:"untillogpos,omitempty" xml:"untillogpos,omitempty"`
	MasterSSLAllowed string `mysql:"Master_SSL_Allowed,VARCHAR(255) NOT NULL" json:"mastersslallowed,omitempty" xml:"mastersslallowed,omitempty"`
	MasterSSLCAFile string `mysql:"Master_SSL_CA_File,VARCHAR(255) NOT NULL" json:"mastersslcafile,omitempty" xml:"mastersslcafile,omitempty"`
	MasterSSLCAPath string `mysql:"Master_SSL_CA_Path,VARCHAR(255) NOT NULL" json:"mastersslcapath,omitempty" xml:"mastersslcapath,omitempty"`
	MasterSSLCert string `mysql:"Master_SSL_Cert,VARCHAR(255) NOT NULL" json:"mastersslcert,omitempty" xml:"mastersslcert,omitempty"`
	MasterSSLCipher string `mysql:"Master_SSL_Cipher,VARCHAR(255) NOT NULL" json:"mastersslcipher,omitempty" xml:"mastersslcipher,omitempty"`
	MasterSSLKey string `mysql:"Master_SSL_Key,VARCHAR(255) NOT NULL" json:"mastersslkey,omitempty" xml:"mastersslkey,omitempty"`
	SecondsBehindMaster uint64 `mysql:"Seconds_Behind_Master,BIGINT UNSIGNED NOT NULL" json:"secondsbehindmaster,omitempty" xml:"secondsbehindmaster,omitempty"`
	MasterSSLVerifyServerCert string `mysql:"Master_SSL_Verify_Server_Cert,VARCHAR(255) NOT NULL" json:"mastersslverifyservercert,omitempty" xml:"mastersslverifyservercert,omitempty"`
	LastIOErrno uint32 `mysql:"Last_IO_Errno,INT UNSIGNED NOT NULL" json:"lastioerrno,omitempty" xml:"lastioerrno,omitempty"`
	LastIOError string `mysql:"Last_IO_Error,VARCHAR(255) NOT NULL" json:"lastioerror,omitempty" xml:"lastioerror,omitempty"`
	LastSQLErrno uint32 `mysql:"Last_SQL_Errno,INT UNSIGNED NOT NULL" json:"lastsqlerrno,omitempty" xml:"lastsqlerrno,omitempty"`
	LastSQLError string `mysql:"Last_SQL_Error,VARCHAR(255) NOT NULL" json:"lastsqlerror,omitempty" xml:"lastsqlerror,omitempty"`
	ReplicateIgnoreServerIds string `mysql:"Replicate_Ignore_Server_Ids,VARCHAR(255) NOT NULL" json:"replicateignoreserverids,omitempty" xml:"replicateignoreserverids,omitempty"`
	MasterServerId uint32 `mysql:"Master_Server_Id,INT UNSIGNED NOT NULL" json:"masterserverid,omitempty" xml:"masterserverid,omitempty"`
	MasterSSLCrl string `mysql:"Master_SSL_Crl,VARCHAR(255) NOT NULL" json:"mastersslcrl,omitempty" xml:"mastersslcrl,omitempty"`
	MasterSSLCrlpath string `mysql:"Master_SSL_Crlpath,VARCHAR(255) NOT NULL" json:"mastersslcrlpath,omitempty" xml:"mastersslcrlpath,omitempty"`
	UsingGtid string `mysql:"Using_Gtid,VARCHAR(255) NOT NULL" json:"usinggtid,omitempty" xml:"usinggtid,omitempty"`
	GtidIOPos string `mysql:"Gtid_IO_Pos,VARCHAR(255) NOT NULL" json:"gtidiopos,omitempty" xml:"gtidiopos,omitempty"`
}

// Bind returns a slice of all columns for a SlaveStatus to use it with Scan.
// Bind only one SlaveStatus, reuse the []interface{} and store a copy after calling Scan.
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
		&v.MasterSSLCrl,
		&v.MasterSSLCrlpath,
		&v.UsingGtid,
		&v.GtidIOPos,
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
func (v *SlaveStatus) Query(location...string) (sql string, args []string) {
	if len(location) != 0 {
		panic("Query must be called without arguments")
	}
	return "SHOW SLAVE STATUS", nil
}


type User struct {
	User string `mysql:"user,CHAR BINARY NOT NULL" json:"user,omitempty" xml:"user,omitempty"`
	Host string `mysql:"host,CHAR BINARY NOT NULL" json:"host,omitempty" xml:"host,omitempty"`
}

// Bind returns a slice of all columns for a User to use it with Scan.
// Bind only one User, reuse the []interface{} and store a copy after calling Scan.
func (v *User) Bind() (cols []interface{}) {
	return []interface{}{
		&v.User,
		&v.Host,
	}
}

// copy returns a shallow copy of the User.
func (v *User) copy() *User {
	if v == nil {
		return nil
	}
	w := *v
	return &w
}

// Copy returns a shallow copy of the User.
func (v *User) Copy() Bindable {
	return v.copy()
}

// Query provides sql metadata used to query for User.
// It returns the SQL string and the arguments needed for prepared statements.
// Query panics if arguments are given.
func (v *User) Query(location...string) (sql string, args []string) {
	if len(location) != 0 {
		panic("Query must be called without arguments")
	}
	return "SELECT user, host FROM mysql.user", nil
}

