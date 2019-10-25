package mysql

import (
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
)

var (
	// ErrConfig 配置类型错误
	ErrConfig = errors.New("config type unknow")
	globle    *Mysql
)

// PoolConfig 配置
type PoolConfig struct {
	Addr    string `gcfg:"addr"`    // 连接地址
	MaxOpen int    `gcfg:"maxopen"` // 最大打开数
	MaxIdle int    `gcfg:"maxidle"` // 最大活跃数
}

// Conn 链接
type Conn interface {
	Query(string, ...interface{}) (*sql.Rows, error)
	QueryRow(string, ...interface{}) *sql.Row
	Exec(string, ...interface{}) (sql.Result, error)
	Prepare(string) (*sql.Stmt, error)
	Begin() (*sql.Tx, error)
	Close()
}

// Mysql 对象
type Mysql struct {
	conn Conn
}

// NewMysql 创建对象
func NewMysql(conf interface{}) (*Mysql, error) {
	switch inst := conf.(type) {
	case PoolConfig:
		return newMysql(NewPool(inst.Addr, inst.MaxIdle, inst.MaxOpen))
	default:
		return nil, ErrConfig
	}
}

func newMysql(conn Conn, err error) (*Mysql, error) {
	if err != nil {
		return nil, err
	}
	return &Mysql{conn}, nil
}

// Close 关闭
func (db *Mysql) Close() {
	db.conn.Close()
}

// Init 初始化
func Init(conf interface{}) error {
	var err error
	globle, err = NewMysql(conf)
	return err
}

// Close 关闭
func Close() {
	if globle != nil {
		globle.Close()
	}
}
