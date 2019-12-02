package model

import (
	"database/sql"
	"fmt"

	signup "github.com/KaustubhLonkar/shop-cart-go/model/signup"
)

//GetUserByUsername
func GetUserByUsername(email string, db *sql.DB) (signup.Data, error) {
	var data signup.Data
	row := db.QueryRow("select * from users where email= ?;", email)

	err := row.Scan(&data.ID, &data.Firstname, &data.Lastname, &data.Email, &data.Password, &data.Create, &data.Update)
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
