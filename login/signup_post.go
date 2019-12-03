package login

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/KaustubhLonkar/shop-cart-go/model/mail"
	"github.com/KaustubhLonkar/shop-cart-go/model/signup"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func RegistrationPost(db *sql.DB, c *gin.Context) {

	firstname := c.PostForm("firstname")
	lastname := c.PostForm("lastname")
	email := c.PostForm("email")
	password := c.PostForm("password")
	//fmt.Println("The first name is:-",firstname)
	// firstname := c.Param("first_name")
	// lastname := c.Param("last_name")
	// email := c.Param("email")
	// password := c.Param("password")
	//fmt.Println("Received all the parameters for sign up", firstname)
	/* password hashing mechanism */
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 14)

	var newuser signup.Data

	newuser.Firstname = firstname
	newuser.Lastname = lastname
	newuser.Email = email
	newuser.Password = string(hashedPassword)
	newuser.Create = time.Now().Format("Mon Jan _2 15:04:05 2006")
	newuser.Update = time.Now().Format("Mon Jan _2 15:04:05 2006")

	comm := mail.Comms{}
	comm.Token = mail.GenerateToken()
	//comm.OTP = mail.GenerateOTP(6)
	comm.Name = newuser.Firstname
	comm.Username = newuser.Email
	comm.Password = password
	fmt.Println("Received all the parameters for sign up", newuser)
	err := signup.RegisterInDB(newuser, db)
	if err == nil {
		msg := "Registration successful, please login !!!"
		m := mail.NewMail(newuser.Email, "Registration successful")
		m.Send("signupmail.gohtml", comm)

		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": msg,
		})
		return
		//c.HTML(http.StatusPermanentRedirect, "login.gohtml", msg)
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
			//	c.HTML(500, "registration.gohtml", msg)
		}
	}

}

// c.JSON(http.StatusOK, gin.H{
// 	"status":  http.StatusOK,
// 	"message": "placed order is successfull...!",
// 	"data":    ord,
// })
// return
// } else {
// c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No Product found !!"})
// return
// }
