package mysql

import (
	"database/sql"
)

// Query 查询
func Query(query string, args ...interface{}) (*sql.Rows, error) {
	return globle.Query(query, args...)
}

// Query 查询
func (db *Mysql) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return db.conn.Query(query, args...)
}
