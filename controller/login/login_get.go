package login

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// func LoginPage() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		sid, err := c.Cookie("session")
// 		if err != nil || sid == "" {
// 			c.HTML(http.StatusOK, "login.gohtml", nil)
// 		} else {
// 			c.Redirect(302, "/homepage")
// 		}
// 	}
// }
