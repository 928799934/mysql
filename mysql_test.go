package mysql

import (
	"log"
	"testing"
	"time"
)

func TestRun(t *testing.T) {
	cfg := PoolConfig{"test:test@tcp(10.60.82.115:3306)/test?parseTime=true&loc=Local&charset=utf8", 10, 10}
	if err := Init(cfg); err != nil {
		log.Printf("Init(%v) error(%v)", cfg, err)
		return
	}
	defer Close()
	rows, err := Query("SELECT id FROM test")
	if err != nil {
		log.Printf("Query(\"SELECT id FROM test\") error(%v)", err)
		return
	}
	defer rows.Close()
	var list []time.Time
	for rows.Next() {
		t := time.Time{}
		if err := rows.Scan(&t); err != nil {
			continue
		}
		list = append(list, t)
	}
	log.Printf("%v", list)
}
