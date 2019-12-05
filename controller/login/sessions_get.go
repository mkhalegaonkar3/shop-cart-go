package login

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/KaustubhLonkar/shop-cart-go/model/session"

	"github.com/gin-gonic/gin"
)

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
