package login

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/mkhalegaonkar3/shop-cart-go/model/login"
	"github.com/mkhalegaonkar3/shop-cart-go/model/mail"
	"github.com/mkhalegaonkar3/shop-cart-go/model/session"
	"github.com/mkhalegaonkar3/shop-cart-go/model/signup"
	uuid "github.com/satori/go.uuid"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

//LoginPage func
func LoginPage() gin.HandlerFunc {
	return func(c *gin.Context) {
		sid, err := c.Cookie("session")
		if err != nil || sid == "" {
			c.HTML(http.StatusOK, "login.gohtml", nil)
		} else {
			c.Redirect(302, "/homepage")
		}
	}
}

// LoginPost func
func LoginPost(db *sql.DB, c *gin.Context) {

	sid, err := c.Cookie("session")
	fmt.Println("Print err:", err)
	fmt.Println("Print sid:", sid)
	if err != nil || sid == "" {
		//fmt.Println("err,sid::", err, sid)
		username := c.PostForm("username")
		password := c.PostForm("password")
		fmt.Println("The username is:", username)
		user, err := login.GetUserByUsername(username, db)

		if err == nil {
			err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
			if err == nil {

				cookie := uuid.Must(uuid.NewV4()).String()
				fmt.Println("The cookie obtained is:", cookie)
				c.SetCookie("session", cookie, 300, "/", "", false, true)
				details, _ := login.GetUserByUsername(username, db)

				session.Add(cookie, details, db)
				//TODO
				//c.Redirect(303, "/homepage?sid="+cookie)

			} else {
				fmt.Println("The err while login is:-", err)
				c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "password not matched"})

			}

		} else {
			c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "username not found"})

		}

	}

}

// Logout func
func Logout(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		sid, _ := c.Cookie("session")
		//fmt.Println("removing session in session table")
		session.RemoveSession(sid, db)
		//fmt.Println("session removed in session table")
		// fmt.Println("sid = ", sid)
		c.SetCookie("session", "", -1, "/", "", false, false)
		c.Redirect(302, "/login")
	}
}

// PasswordReset func
func PasswordReset(db *sql.DB, c *gin.Context) {

	c1, err := c.Cookie("selfserve")
	newpass := c.PostForm("newpass")
	if err == nil {
		stmt, err := db.Prepare("update users set password= ? where email= ?;")
		if err != nil {
			fmt.Print(err.Error())
		}
		password, err := bcrypt.GenerateFromPassword([]byte(newpass), 14)
		_, err = stmt.Exec(password, c1)
		if err != nil {
			fmt.Print(err.Error())
			c.HTML(http.StatusOK, "login.gohtml", "password reset failed")
		} else {
			c.SetCookie("selfserve", "", -1, "/", "", false, false)
			msg := "password reset successfully"
			c.HTML(303, "login.gohtml", msg)
		}
		comm := mail.Comms{}
		comm.Name = c1
		comm.Username = c1
		comm.Password = newpass
		m := mail.NewMail(c1, "Password reset successful")
		m.Send("resetmail.gohtml", comm)

	} else {
		c.String(400, "please verify user username")
	}

}

// ResetPasswordPage func
func ResetPasswordPage() gin.HandlerFunc {

	return func(c *gin.Context) {

		c.HTML(http.StatusOK, "resetpass.gohtml", nil)
	}

}

// ActiveSession func
func ActiveSession(db *sql.DB) gin.HandlerFunc {

	return func(c *gin.Context) {
		results, err := session.GetAllActiveSessions(db)
		if err != nil {
			fmt.Println(err)
		}
		//fmt.Println("results before sending", results)
		c.JSON(http.StatusOK, results)
	}

}

// RegistrationPost func
func RegistrationPost(db *sql.DB, c *gin.Context) {

	firstname := c.PostForm("firstname")
	lastname := c.PostForm("lastname")
	email := c.PostForm("email")
	password := c.PostForm("password")

	/* password hashing mechanism */
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 0)

	var newUser signup.Data

	newUser.Firstname = firstname
	newUser.Lastname = lastname
	newUser.Email = email
	newUser.Password = string(hashedPassword)
	newUser.Create = time.Now().Format("2006-01-02 15:04:05")
	newUser.Update = time.Now().Format("2006-01-02 15:04:05")

	comm := mail.Comms{}
	comm.Token = mail.GenerateToken()

	comm.Name = newUser.Firstname
	comm.Username = newUser.Email
	comm.Password = password
	fmt.Println("Received all the parameters for sign up", newUser)
	err := signup.RegisterInDB(newUser, db)
	if err == nil {
		msg := "Registration successful, please login !!!"
		m := mail.NewMail(newUser.Email, "Registration successful")
		m.Send("signupmail.gohtml", comm)

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

func VerifyUsername(db *sql.DB, c *gin.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.PostForm("username")

		exist := login.VerifyUsername(username, db)

		if exist {
			c.SetCookie("selfserve", username, 300, "/", "", false, true)
			c.HTML(http.StatusOK, "confirmpass.gohtml", username)
		} else {
			c.String(400, "username does not exist")
		}
	}
}
