package model

import (
	"database/sql"
	"fmt"
)

type Getter interface {
	GetAllRegisterdInDB(db *sql.DB) []Data
}

//for database storage
type Data struct {
	ID        int    `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Create    string `json:"create"`
	Update    string `json:"update"`
}

type List struct {
	Users []Data
}

func RegisterInDB(newuser Data, db *sql.DB) error {

	stmt, err := db.Prepare("insert into users (firstname,lastname,email,password ,creation_time,updation_time) values(?,?,?,?,?,?);")
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

func GetAllRegisterdInDB(db *sql.DB) []Data {

	var data Data
	var allusers List
	rows, err := db.Query("select * from users;")
	defer rows.Close()
	if err != nil {
		fmt.Print(err.Error())
	}
	for rows.Next() {
		err = rows.Scan(&data.ID, &data.Firstname, &data.Lastname, &data.Email, &data.Password, &data.Create, &data.Update)
		allusers.Users = append(allusers.Users, data)
		if err != nil {
			fmt.Print(err.Error())
		}
	}
	//fmt.Println("data received : ", allusers.Users)
	return allusers.Users
}
