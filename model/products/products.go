package products

import (
	"database/sql"
	"fmt"
)

// Product struct for storing product info
type Product struct {
	ProductName     string `json:"productName"`
	ProductQuantity int    `json:"productQuantity"`
	CategoryID      int    `json:"categoryID"`
}

// AddProductsInDB func
func AddProductsInDB(product Product, db *sql.DB) error {

	stmt, err := db.Prepare("insert into products (productName,quantity,categoryID) values(?,?,?);")
	if err != nil {
		fmt.Print(err.Error())
	}
	defer stmt.Close()
	_, err = stmt.Exec(product.ProductName, product.ProductQuantity, product.CategoryID)

	if err != nil {
		fmt.Print(err.Error())
	}

	return err
}
