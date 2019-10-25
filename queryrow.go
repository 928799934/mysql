package mysql

import (
	"database/sql"
)

// QueryRow 查询
func QueryRow(query string, args ...interface{}) *sql.Row {
	return globle.QueryRow(query, args...)
}

// QueryRow 查询
func (db *Mysql) QueryRow(query string, args ...interface{}) *sql.Row {
	return db.conn.QueryRow(query, args...)
}
