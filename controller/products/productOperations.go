package products

import (
	"database/sql"
	"net/http"
	"strconv"

	productsDb "github.com/KaustubhLonkar/shop-cart-go/model/products"
	"github.com/gin-gonic/gin"
)

// AddProduct func
func AddProduct(db *sql.DB, c *gin.Context) {

	productName := c.PostForm("productName")
	productQuantity, _ := strconv.Atoi(c.PostForm("productQuantity"))
	categoryID, _ := strconv.Atoi(c.PostForm("categoryID"))

	var product productsDb.Product
	product.ProductName = productName
	product.ProductQuantity = productQuantity
	product.CategoryID = categoryID

	err := productsDb.AddProductsInDB(product, db)

	if err == nil {
		msg := "Your product has been added successfully !!!"

		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": msg,
		})
		return

	} else {
		var msg string

		msg = "Internal server error"
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": msg})
		return

	}

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
