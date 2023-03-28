package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/nicolerobin/log"
)

func main() {
	db, err := sql.Open("mysql", "robin:robin@tcp(localhost)/test")
	if err != nil {
		log.Error("sql.Open() failed, err:%s", err)
	}

	rows, err := db.Query("select * from test_gap_lock")
	if err != nil {
		log.Error("db.Exec() failed, err:%s", err)
	}
	log.Debug("rows:%+v", rows)
	for rows.Next() {
		var (
			id   int64
			name string
		)
		if err := rows.Scan(&id, &name); err != nil {
			log.Error("rows.Scan() failed, err:%s", err)
			continue
		}
		log.Debug("id:%d, name:%s", id, name)
	}
}
