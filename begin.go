package mysql

import (
	"database/sql"
)

// Begin 开启事务
func Begin() (*sql.Tx, error) {
	return globle.Begin()
}

// Begin 开启事务
func (db *Mysql) Begin() (*sql.Tx, error) {
	return db.conn.Begin()
}
