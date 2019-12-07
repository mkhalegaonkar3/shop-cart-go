package category

import (
	"database/sql"
	"fmt"
)

// Category struct
type Category struct {
	CategoryName string `json:"categoryName"`
}

//CreateCategoryInDB func
func CreateCategoryInDB(category Category, db *sql.DB) error {

	stmt, err := db.Prepare("insert into category (categoryName) values(?);")
	if err != nil {
		fmt.Print(err.Error())
	}
	defer stmt.Close()
	_, err = stmt.Exec(category.CategoryName)

	if err != nil {
		fmt.Print(err.Error())
	}

	return err
}
