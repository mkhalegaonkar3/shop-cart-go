package login

import (
	"database/sql"
	"net/http"

	"github.com/KaustubhLonkar/shop-cart-go/model/login"
	"github.com/KaustubhLonkar/shop-cart-go/model/session"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

func LoginPost(db *sql.DB, c *gin.Context) {

	sid, err := c.Cookie("session")
	//	fmt.Println("err,sid:", err, sid)
	if err != nil || sid == "" {
		//fmt.Println("err,sid::", err, sid)
		username := c.PostForm("username")
		password := c.PostForm("password")

		user, err := login.GetUserByUsername(username, db)

		if err == nil {
			err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
			if err == nil {

				cookie := uuid.Must(uuid.NewV4()).String()
				c.SetCookie("session", cookie, 300, "/", "", false, true)
				details, _ := login.GetUserByUsername(username, db)

				session.Add(cookie, details, db)
				//TODO
				c.Redirect(303, "/homepage?sid="+cookie)

			} else {
				c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "password not matched"})
				//c.String(400, "password not matched")
			}

		} else {
			c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "username not found"})

			//c.JSON(404, "username not found")
		}

	}

}
