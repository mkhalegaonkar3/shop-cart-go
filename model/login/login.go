package login

import (
	"database/sql"
	"fmt"

	signup "github.com/mkhalegaonkar3/shop-cart-go/model/signup"
)

//GetUserByUsername func
func GetUserByUsername(username string, db *sql.DB) (signup.UserData, error) {
	var data signup.UserData
	row := db.QueryRow("select userID,firstName,lastname,username,password from users where username= ?;", username)

	err := row.Scan(&data.ID, &data.Firstname, &data.Lastname, &data.Email, &data.Password)
	if err != nil {
		fmt.Print("No details found", err.Error())
	}

	return data, err

}

//VerifyUsername func
func VerifyUsername(email string, db *sql.DB) bool {
	var data signup.UserData
	row := db.QueryRow("select email from users where email= ?;", email)
	fmt.Println(row)
	err := row.Scan(&data.Email)
	if err != nil {
		fmt.Print("No details found", err.Error())
		return false
	}
	return true
}
