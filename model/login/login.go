package login

import (
	"database/sql"
	"fmt"

	signup "github.com/mkhalegaonkar3/shop-cart-go/model/signup"
)

//GetUserByUsername
func GetUserByUsername(username string, db *sql.DB) (signup.Data, error) {
	var data signup.Data
	row := db.QueryRow("select firstName,lastname,username,creation_time from users where username= ?;", username)

	err := row.Scan(&data.Firstname, &data.Lastname, &data.Email, &data.Create)
	if err != nil {
		fmt.Print("No details found", err.Error())
	}

	return data, err

}

func VerifyUsername(email string, db *sql.DB) bool {
	var data signup.Data
	row := db.QueryRow("select email from users where email= ?;", email)
	fmt.Println(row)
	err := row.Scan(&data.Email)
	if err != nil {
		fmt.Print("No details found", err.Error())
		return false
	}
	return true
}
