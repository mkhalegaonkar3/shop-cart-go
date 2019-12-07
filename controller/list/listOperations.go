package list

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	listDb "github.com/KaustubhLonkar/shop-cart-go/model/list"

	"github.com/gin-gonic/gin"
)

// CreateList func
func CreateList(db *sql.DB, c *gin.Context) {

	userID, _ := strconv.Atoi(c.PostForm("userid"))
	listTitle := c.PostForm("listTitle")

	var list listDb.List

	list.UserID = userID
	list.ListTitle = listTitle
	list.Create = time.Now().Format("Mon Jan _2 15:04:05 2006")
	list.Update = time.Now().Format("Mon Jan _2 15:04:05 2006")
	list.Delete = time.Now().Format("Mon Jan _2 15:04:05 2006")

	fmt.Println("Received all the parameters for creating a list:-", list.ListTitle)
	err := listDb.CreateList(list, db)
	if err == nil {
		msg := "Your cart has been created successful !!!"

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

// AddItemsList func
func AddItemsList(db *sql.DB, c *gin.Context) {

	var listDetails listDb.ListDetails

	productID, _ := strconv.Atoi(c.PostForm("productid"))
	productName := c.PostForm("ProductName")
	listTitle := c.PostForm("listTitle")
	username := c.PostForm("listTitle")

	listDetails.ProductID = productID
	listDetails.ProductName = productName
	listDetails.ListTitle = listTitle
	listDetails.Create = time.Now().Format("Mon Jan _2 15:04:05 2006")
	listDetails.Update = time.Now().Format("Mon Jan _2 15:04:05 2006")
	listDetails.Delete = time.Now().Format("Mon Jan _2 15:04:05 2006")
	listDetails.ModifiedBy = username

	fmt.Println("Received all the parameters for list of items in cart", listDetails)
	err := listDb.AddItemsList(listDetails, db)
	if err == nil {
		msg := "List creation was successful, please login !!!"

		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": msg,
		})
		return

	} else {
		var msg string
		if strings.Contains(err.Error(), "Error 1062") {
			msg = "List Details already exist,please try with another email !!!"
			c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": msg})
			return
		} else {
			msg = "Internal server error"
			c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": msg})
			return

		}
	}
}

// DeleteItemList func
func DeleteItemList(db *sql.DB, c *gin.Context) {

	productID, _ := strconv.Atoi(c.PostForm("productID"))
	productName := c.PostForm("productName")
	listTitle := c.PostForm("listTitle")

	fmt.Println("Received all the parameters for deleting item from list", productID, " ", productName, " ", listTitle)
	err := listDb.DeleteItemList(productID, productName, listTitle, db)
	if err == nil {
		msg := "Deletion of item from list successful !!!"

		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": msg,
		})
		return

	} else {
		var msg string
		if strings.Contains(err.Error(), "Error 1062") {
			msg = "email already exist,please try with another email !!!"
			c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": msg})
			return
		} else {
			msg = "Internal server error"
			c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": msg})
			return

		}
	}
}

// DeleteList func
func DeleteList(db *sql.DB, c *gin.Context) {

	listid, _ := strconv.Atoi(c.PostForm("listid"))
	listname := c.PostForm("listTitle")

	fmt.Println("Received all the parameters for deleting list", listid, listname)
	err := listDb.DeleteList(listid, listname, db)
	if err == nil {
		msg := "Deletion of list successful!!!"

		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": msg,
		})
		return

	} else {
		var msg string
		if strings.Contains(err.Error(), "Error 1062") {
			msg = "email already exist,please try with another email !!!"
			c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": msg})
			return
		} else {
			msg = "Internal server error"
			c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": msg})
			return

		}
	}
}
