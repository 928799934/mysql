package mysql

import (
	"database/sql"
)

type connPool struct {
	db *sql.DB
}

// NewPool 创建对象
func NewPool(addr string, initial, maximum int) (Conn, error) {
	db, err := sql.Open("mysql", addr)
	if err != nil {
		return nil, err
	}
	db.SetMaxIdleConns(initial)
	db.SetMaxOpenConns(maximum)

	return &connPool{db}, nil
}

// Close 关闭
func (con *connPool) Close() {
	con.db.Close()
}

// Query 查询
func (con *connPool) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return con.db.Query(query, args...)
}

// QueryRow 查询
func (con *connPool) QueryRow(query string, args ...interface{}) *sql.Row {
	return con.db.QueryRow(query, args...)
}

// Exec 执行语句
func (con *connPool) Exec(query string, args ...interface{}) (sql.Result, error) {
	return con.db.Exec(query, args...)
}

// Prepare 预加载
func (con *connPool) Prepare(query string) (*sql.Stmt, error) {
	return con.db.Prepare(query)
}

// Begin 开启事务
func (con *connPool) Begin() (*sql.Tx, error) {
	return con.db.Begin()
}
