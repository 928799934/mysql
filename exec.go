package mysql

import (
	"database/sql"
)

// Exec 运行语句
func Exec(query string, args ...interface{}) (sql.Result, error) {
	return globle.Exec(query, args...)
}

// Exec 运行语句
func (db *Mysql) Exec(query string, args ...interface{}) (sql.Result, error) {
	return db.conn.Exec(query, args...)
}
