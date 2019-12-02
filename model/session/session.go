package session

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/KaustubhLonkar/shop-cart-go/model/signup"
)

type Session struct {
	Sid         string `json:"sid"`
	Username    string `json:"username"`
	Create_time string `json:"create_time"`
	Update_time string `json:"update_time"`
}

type Repo struct {
	Sessions []Session `json:"sessions"`
}

func Add(sid string, details signup.Data, db *sql.DB) {
	//stmt, err := db.Prepare("insert into users (firstname,lastname,email,password ,creation_time,updation_time) values(?,?,?,?,?,?);")
	stmt, err := db.Prepare("insert into sessions  (sid,username,create_time,update_time) values(?,?,?,?)")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	email := details.Email
	create_time := time.Now().Format("Mon Jan _2 15:04:05 2006")
	update_time := time.Now().Format("Mon Jan _2 15:04:05 2006")
	defer stmt.Close()
	_, err = stmt.Exec(sid, email, create_time, update_time)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

}

func GetAllActiveSessions(db *sql.DB) ([]Session, error) {
	rows, err := db.Query("select * from sessions;")
	defer rows.Close()
	var result Session
	var results Repo

	if err != nil {
		fmt.Println("error in fetching details", err.Error())
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&result.Sid, &result.Username, &result.Create_time, &result.Update_time)
		if err != nil {
			fmt.Println("error in scanning all session", err.Error())
		}
		//fmt.Println("printing fetched results", result.Sid)
		results.Sessions = append(results.Sessions, result)

	}
	//fmt.Println("in session get call: ", results.Sessions)
	return results.Sessions, nil

}

func RemoveSession(sid string, db *sql.DB) {

	result := db.QueryRow("delete from sessions where sid=?", sid)

	fmt.Println("sessions after removing:", result)

}

func Search(sid string, db *sql.DB) bool {

	result := db.QueryRow("select * from table where sid=?", sid)
	fmt.Println(result)
	return true
}
