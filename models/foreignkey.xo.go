// Package models contains the types for schema 'public'.
package models

// GENERATED BY XO. DO NOT EDIT.

// ForeignKey represents a foreign key.
type ForeignKey struct {
	ForeignKeyName string // foreign_key_name
	TableName      string // table_name
	ColumnName     string // column_name
	RefIndexName   string // ref_index_name
	RefTableName   string // ref_table_name
	RefColumnName  string // ref_column_name
	Comment        string // comment
}

// PgForeignKeysBySchema runs a custom query, returning results as ForeignKey.
func PgForeignKeysBySchema(db XODB, schema string) ([]*ForeignKey, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`r.conname, ` + // ::varchar AS foreign_key_name
		`a.relname, ` + // ::varchar AS table_name
		`b.attname, ` + // ::varchar AS column_name
		`i.relname, ` + // ::varchar AS ref_index_name
		`c.relname, ` + // ::varchar AS ref_table_name
		`d.attname, ` + // ::varchar AS ref_column_name
		`'' ` + // ::varchar AS comment
		`FROM pg_constraint r ` +
		`JOIN ONLY pg_class a ON a.oid = r.conrelid ` +
		`JOIN ONLY pg_attribute b ON b.attisdropped = false AND b.attnum = ANY(r.conkey) AND b.attrelid = r.conrelid ` +
		`JOIN ONLY pg_class i on i.oid = r.conindid ` +
		`JOIN ONLY pg_class c on c.oid = r.confrelid ` +
		`JOIN ONLY pg_attribute d ON d.attisdropped = false AND d.attnum = ANY(r.confkey) AND d.attrelid = r.confrelid ` +
		`JOIN ONLY pg_namespace n ON n.oid = r.connamespace ` +
		`WHERE r.contype = 'f' AND n.nspname = $1 ` +
		`ORDER BY r.conname, a.relname, b.attname`

	// run query
	XOLog(sqlstr, schema)
	q, err := db.Query(sqlstr, schema)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*ForeignKey{}
	for q.Next() {
		fk := ForeignKey{}

		// scan
		err = q.Scan(&fk.ForeignKeyName, &fk.TableName, &fk.ColumnName, &fk.RefIndexName, &fk.RefTableName, &fk.RefColumnName, &fk.Comment)
		if err != nil {
			return nil, err
		}

		res = append(res, &fk)
	}

	return res, nil
}

// MyForeignKeysBySchema runs a custom query, returning results as ForeignKey.
func MyForeignKeysBySchema(db XODB, schema string) ([]*ForeignKey, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`constraint_name AS foreign_key_name, ` +
		`table_name AS table_name, ` +
		`column_name AS column_name, ` +
		`'' AS ref_index_name, ` +
		`referenced_table_name AS ref_table_name, ` +
		`referenced_column_name AS ref_column_name ` +
		`FROM information_schema.key_column_usage ` +
		`WHERE referenced_table_name IS NOT NULL AND table_schema = ? ` +
		`ORDER BY table_name, constraint_name`

	// run query
	XOLog(sqlstr, schema)
	q, err := db.Query(sqlstr, schema)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*ForeignKey{}
	for q.Next() {
		fk := ForeignKey{}

		// scan
		err = q.Scan(&fk.ForeignKeyName, &fk.TableName, &fk.ColumnName, &fk.RefIndexName, &fk.RefTableName, &fk.RefColumnName)
		if err != nil {
			return nil, err
		}

		res = append(res, &fk)
	}

	return res, nil
}
