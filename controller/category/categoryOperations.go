package category

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	categoryDb "github.com/mkhalegaonkar3/shop-cart-go/model/category"
)

// AddCategory func
func AddCategory(db *sql.DB, c *gin.Context) {

	categoryName := c.PostForm("categoryName")

	var category categoryDb.Category
	category.CategoryName = categoryName

	err := categoryDb.CreateCategoryInDB(category, db)

	if err == nil {
		msg := "Category has been created successfully !!!"

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
