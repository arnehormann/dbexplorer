dbexplorer
==========

Scan your database and create code.

See `gen.go` for an example with administrative queries.

Exact workings are undocumented so far, but generated code may be used somewhat like this:

```go
// appendMystruct fetches Mystruct values from the database and appends them to dest.
func appendMystruct(dest []*Mystruct, db *sql.DB) []*Mystruct {
	bound := &Mystruct{}
	dest := bound.Bind()
	// second arg is only used for variable identifiers like schemas, tables, users and hosts
	query, _ := bound.Query()

	rows, err := db.Query(query)
	if err != nil {
		panic(err) // TODO fix
	}
	for rows.Next() {
		err = rows.Scan(dest...)
		if err != nil {
			panic(err) // TODO fix
		}
		dest = append(dest, bound.copy())
	}
	return dest
}
```
