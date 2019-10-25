package mysql

// Ping 检测链接
func Ping() error {
	return globle.Ping()
}

// Ping 检测链接
func (db *Mysql) Ping() error {
	_, err := db.conn.Exec("SELECT 1")
	return err
}
