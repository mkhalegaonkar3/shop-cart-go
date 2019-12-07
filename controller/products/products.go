package products

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	//"github.com/KaustubhLonkar/order-management-go/model"
	"github.com/gin-gonic/gin"
)

type Product struct {
	productName     string `json:"productName"`
	productQuantity int    `json:"productQuantity"`
	categoryID      int    `json:"categoryID"`
}

// AddProduct func
func AddProduct(db *sql.DB, c *gin.Context) {

	productName := c.PostForm("productName")
	productQuantity, _ := strconv.Atoi(c.PostForm("productQuantity"))
	categoryID, _ := strconv.Atoi(c.PostForm("categoryID"))

	var product Product
	product.productName = productName
	product.productQuantity = productQuantity
	product.categoryID = categoryID

	stmt, err := db.Prepare("insert into products (productName,quantity,categoryID) values(?,?,?);")
	if err != nil {
		fmt.Print(err.Error())
	}
	defer stmt.Close()
	_, err = stmt.Exec(product.productName, product.productQuantity, product.categoryID)

	if err != nil {
		fmt.Print(err.Error())
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Product added successfully",
	})

}

// //IsProductAvailable to check availability
// func IsProductAvailable(pname string, qty int) (bool, product, int) {
// 	available := false
// 	amt := 0
// 	var prod product
// 	model.Db.Where("product_name = ?", pname).First(&prod)

// 	if pname == prod.ProductName && qty <= prod.ProductQuantity {
// 		available = true
// 		amt = prod.ProductPrice * qty
// 		remainingQuantity := prod.ProductQuantity - qty
// 		model.Db.Model(&prod).Update("product_quantity", remainingQuantity)
// 		return available, prod, amt
// 	}
// 	return available, prod, amt
// }
