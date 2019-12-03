package login

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegistrationPage() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "registration.gohtml", nil)
	}
}
