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

// CreatList func
func CreateList(db *sql.DB, c *gin.Context) {

	userID, _ := strconv.Atoi(c.PostForm("userid"))
	listTitle := c.PostForm("listTitle")

	var list listDb.List

	list.UserID = userID
	list.ListTitle = listTitle
	list.Create = time.Now().Format("Mon Jan _2 15:04:05 2006")
	list.Update = time.Now().Format("Mon Jan _2 15:04:05 2006")
	list.Delete = time.Now().Format("Mon Jan _2 15:04:05 2006")

	fmt.Println("Received all the parameters for sign up", list)
	err := listDb.CreateList(list, db)
	if err == nil {
		msg := "Registration successful, please login !!!"

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

// CreatList func
func AddItemsList(db *sql.DB, c *gin.Context) {

	userID, _ := strconv.Atoi(c.PostForm("userid"))
	listTitle := c.PostForm("listTitle")

	var list listDb.List

	list.UserID = userID
	list.ListTitle = listTitle
	list.Create = time.Now().Format("Mon Jan _2 15:04:05 2006")
	list.Update = time.Now().Format("Mon Jan _2 15:04:05 2006")
	list.Delete = time.Now().Format("Mon Jan _2 15:04:05 2006")

	fmt.Println("Received all the parameters for sign up", list)
	err := listDb.CreateList(list, db)
	if err == nil {
		msg := "Registration successful, please login !!!"

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

// CreatList func
func DeleteItemList(db *sql.DB, c *gin.Context) {

	userID, _ := strconv.Atoi(c.PostForm("userid"))
	listTitle := c.PostForm("listTitle")

	var list listDb.List

	list.UserID = userID
	list.ListTitle = listTitle
	list.Create = time.Now().Format("Mon Jan _2 15:04:05 2006")
	list.Update = time.Now().Format("Mon Jan _2 15:04:05 2006")
	list.Delete = time.Now().Format("Mon Jan _2 15:04:05 2006")

	fmt.Println("Received all the parameters for sign up", list)
	err := listDb.CreateList(list, db)
	if err == nil {
		msg := "Registration successful, please login !!!"

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

// CreatList func
func DeleteList(db *sql.DB, c *gin.Context) {

	userID, _ := strconv.Atoi(c.PostForm("userid"))
	listTitle := c.PostForm("listTitle")

	var list listDb.List

	list.UserID = userID
	list.ListTitle = listTitle
	list.Create = time.Now().Format("Mon Jan _2 15:04:05 2006")
	list.Update = time.Now().Format("Mon Jan _2 15:04:05 2006")
	list.Delete = time.Now().Format("Mon Jan _2 15:04:05 2006")

	fmt.Println("Received all the parameters for sign up", list)
	err := listDb.CreateList(list, db)
	if err == nil {
		msg := "Registration successful, please login !!!"

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



