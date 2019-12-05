package list

import (
	"database/sql"
	"fmt"
)

type List struct {
	UserID    int `json:"userID"`
	ListTitle string `json:"listTitle"`
	Create    string `json:"create"`
	Update    string `json:"update"`
	Delete    string `json:"delete"`
	Status    string `json:"status"`
}

func CreateList(list List, db *sql.DB) error {
	fmt.Println("New User:-", list.UserID)
	fmt.Println("New User:-", list.ListTitle)
	fmt.Println("New User:-", list.Create)
	fmt.Println("New User:-", list.Update)
	fmt.Println("New User:-", list.Delete)
	fmt.Println("New User:-", list.Status)

	stmt, err := db.Prepare("insert into list (userID,listTitlename,creation_time,modified_time,deletion_time,status) values(?,?,?,?,?,?);")
	if err != nil {
		fmt.Print(err.Error())
	}
	defer stmt.Close()
	_, err = stmt.Exec(list.UserID, list.ListTitle, list.Create, list.Update, list.Delete, list.Status)

	if err != nil {
		fmt.Print(err.Error())
	}

	return err
}
