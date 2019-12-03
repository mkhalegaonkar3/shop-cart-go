package login

import (
	"database/sql"

	"github.com/KaustubhLonkar/shop-cart-go/model/session"

	"github.com/gin-gonic/gin"
)

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
