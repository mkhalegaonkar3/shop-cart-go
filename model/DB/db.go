package DB

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Start() (*sql.DB, error) {
	var err error

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/shopcart")
	//fmt.Println("db", db, "err", err)
	return db, err

}
