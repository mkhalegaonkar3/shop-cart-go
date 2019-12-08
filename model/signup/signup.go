package signup

import (
	"database/sql"
	"fmt"
)

// type Getter interface {
// 	GetAllRegisterdInDB(db *sql.DB) []UserData
// }

//for database storage
type UserData struct {
	ID        int    `json:"-"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Create    string `json:"create"`
	Update    string `json:"update"`
}

type List struct {
	Users []UserData
}

func RegisterInDB(newuser UserData, db *sql.DB) error {

	stmt, err := db.Prepare("insert into users (firstName,lastname,username,password,creation_time,modified_time) values(?,?,?,?,?,?);")
	if err != nil {
		fmt.Print(err.Error())
	}
	defer stmt.Close()
	_, err = stmt.Exec(newuser.Firstname, newuser.Lastname, newuser.Email, newuser.Password, newuser.Create, newuser.Update)

	if err != nil {
		fmt.Print(err.Error())
	}

	return err
}

// func GetAllRegisterdInDB(db *sql.DB) []UserData {

// 	var data UserData
// 	var allusers List
// 	rows, err := db.Query("select * from users;")
// 	defer rows.Close()
// 	if err != nil {
// 		fmt.Print(err.Error())
// 	}
// 	for rows.Next() {
// 		err = rows.Scan(&data.ID, &data.Firstname, &data.Lastname, &data.Email, &data.Password, &data.Create, &data.Update)
// 		allusers.Users = append(allusers.Users, data)
// 		if err != nil {
// 			fmt.Print(err.Error())
// 		}
// 	}
// 	//fmt.Println("data received : ", allusers.Users)
// 	return allusers.Users
// }
