package login

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/KaustubhLonkar/shop-cart-go/model/session"

	"github.com/gin-gonic/gin"
)

func Homepage(db *sql.DB) gin.HandlerFunc {

	return func(c *gin.Context) {

		type Response struct {
			Page       int `json:"page"`
			PerPage    int `json:"per_page"`
			Total      int `json:"total"`
			TotalPages int `json:"total_pages"`
			Data       []struct {
				ID        int    `json:"id"`
				Email     string `json:"email"`
				FirstName string `json:"first_name"`
				LastName  string `json:"last_name"`
				Avatar    string `json:"avatar"`
			} `json:"data"`
		}

		sid, err := c.Cookie("session")
		if err != nil || sid == "" {
			c.Redirect(302, "/login")
		} else if session.Search(sid, db) {
			//sessionname, _ := c.Cookie("session")
			//sid := c.Query("sid")
			result, _ := http.Get("https://reqres.in/api/users?page=2")
			data, _ := ioutil.ReadAll(result.Body)
			var response Response
			err := json.Unmarshal(data, &response)
			if err != nil {
				fmt.Println(err.Error())
			}
			c.HTML(http.StatusOK, "homepage.gohtml", response.Data)
		}
	}
}
