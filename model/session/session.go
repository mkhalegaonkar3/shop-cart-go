package session

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/mkhalegaonkar3/shop-cart-go/model/signup"
)

type Session struct {
	Sid        string `json:"sid"`
	UserID     int    `json:"userID"`
	createTime string `json:"createTime"`
	updateTime string `json:"updateTime"`
}

type Repo struct {
	Sessions []Session `json:"sessions"`
}

func Add(sid string, details signup.UserData, db *sql.DB) {

	stmt, err := db.Prepare("insert into sessions  (sid,userID,creation_time,modified_time) values(?,?,?,?)")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	userID := details.ID
	createTime := time.Now().Format("2006-01-02 15:04:05")
	updateTime := time.Now().Format("2006-01-02 15:04:05")
	defer stmt.Close()
	_, err = stmt.Exec(sid, userID, createTime, updateTime)
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
		err = rows.Scan(&result.Sid, &result.UserID, &result.createTime, &result.updateTime)
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

func GetSessionId(sid string, db *sql.DB) (Session, error) {

	var session Session
	row := db.QueryRow("select sid,userID from sessions where sid=?", sid)
	// fmt.Println("Getting the SessionId", result)
	err := row.Scan(&session.Sid, &session.UserID)
	if err != nil {
		fmt.Print("No details found", err.Error())
	}

	return session, err

}
