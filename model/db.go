package model

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Start() {
	var err error

	db, err = sql.Open("mysql", "root:root@/OrderManagement?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect to database")
	}

}
