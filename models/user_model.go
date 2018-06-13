package models

import (
	"Watermelon/db"
	"log"
)

type User struct {
	UserID int64
	Name   string
	Age    string
}

func Addone(user *User) (int64, error) {
	res, err := db.Orm.Exec("INSERT INTO userInfo(name, age) VALUES (?, ?)", user.Name, user.Age)
	if err != nil {
		log.Println("error on adding one into person :", err)
	}
	backId, err := res.LastInsertId()
	return backId, err
}
