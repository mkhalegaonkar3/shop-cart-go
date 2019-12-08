package list

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"time"

	listDb "github.com/mkhalegaonkar3/shop-cart-go/model/list"
	sessionDb "github.com/mkhalegaonkar3/shop-cart-go/model/session"

	"github.com/gin-gonic/gin"
)

// CreateList func
func CreateList(db *sql.DB, c *gin.Context) {
	userID, _ := strconv.Atoi(c.PostForm("userid"))
	listTitle := c.PostForm("listTitle")
	sid, _ := c.Cookie("session")
	var session sessionDb.Session
	fmt.Println("Print sid:", sid)
	session, _ = sessionDb.GetSessionId(sid, db)

	if sid == session.Sid && userID == session.UserID {
		var list listDb.List

		list.UserID = userID
		list.ListTitle = listTitle
		list.Create = time.Now().Format("2006-01-02 15:04:05")
		list.Update = time.Now().Format("2006-01-02 15:04:05")
		list.Delete = time.Now().Format("2006-01-02 15:04:05")

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
	} else {
		var msg string
		msg = "Your session is expired.Please login"
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": msg})
		return
	}

}

// AddItemsList func
func AddItemsList(db *sql.DB, c *gin.Context) {

	sid, _ := c.Cookie("session")
	var session sessionDb.Session
	fmt.Println("Print sid:", sid)
	session, _ = sessionDb.GetSessionId(sid, db)

	if sid == session.Sid {

		listID, _ := strconv.Atoi(c.PostForm("listID"))
		productID, _ := strconv.Atoi(c.PostForm("productid"))
		productName := c.PostForm("ProductName")
		listTitle := c.PostForm("listTitle")
		username := c.PostForm("username")
		var listDetails listDb.ListDetails

		listDetails.ListID = listID
		listDetails.ProductID = productID
		listDetails.ProductName = productName
		listDetails.ListTitle = listTitle
		listDetails.Create = time.Now().Format("2006-01-02 15:04:05")
		listDetails.Update = time.Now().Format("2006-01-02 15:04:05")
		listDetails.Delete = time.Now().Format("2006-01-02 15:04:05")
		listDetails.ModifiedBy = username

		fmt.Println("Received all the parameters for list of items in cart", listDetails)
		err := listDb.AddItemsList(listDetails, db)
		if err == nil {
			msg := "Product was added was successfully!!!"

			c.JSON(http.StatusOK, gin.H{
				"status":  http.StatusOK,
				"message": msg,
			})
			return

		} else {
			var msg string

			msg = "The list can not be found"
			c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": msg})
			return

		}
	} else {
		var msg string
		msg = "Your session is expired.Please login"
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": msg})
		return
	}

}

// DeleteItemList func
func DeleteItemList(db *sql.DB, c *gin.Context) {

	sid, _ := c.Cookie("session")
	var session sessionDb.Session
	fmt.Println("Print sid:", sid)
	session, _ = sessionDb.GetSessionId(sid, db)

	if sid == session.Sid {
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

			msg = "The item to be deleted can not found"
			c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": msg})
			return

		}
	} else {
		var msg string
		msg = "Your session is expired.Please login"
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": msg})
		return
	}

}

// DeleteList func
func DeleteList(db *sql.DB, c *gin.Context) {

	sid, _ := c.Cookie("session")
	var session sessionDb.Session
	fmt.Println("Print sid:", sid)
	session, _ = sessionDb.GetSessionId(sid, db)

	if sid == session.Sid {
		listid, _ := strconv.Atoi(c.PostForm("listID"))
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

			msg = "List not found"
			c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": msg})
			return

		}
	} else {
		var msg string
		msg = "Your session is expired.Please login"
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": msg})
		return
	}

}
