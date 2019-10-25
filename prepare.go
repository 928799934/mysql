package mysql

import (
	"database/sql"
)

// Prepare 预处理
func Prepare(query string) (*sql.Stmt, error) {
	return globle.Prepare(query)
}

// Prepare 预处理
func (db *Mysql) Prepare(query string) (*sql.Stmt, error) {
	return db.conn.Prepare(query)
}
