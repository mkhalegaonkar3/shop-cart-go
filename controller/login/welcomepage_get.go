package login

import (
	"github.com/gin-gonic/gin"
)

func WelcomePage() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(200, "index.gohtml", gin.H{
			"title": "Home Page",
		})
	}
}
