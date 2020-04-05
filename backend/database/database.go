package databaseutil

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Postinfo struct {
	firstname string
	lastname  string
	date      string
	mail      string
}

func Createperson(firstname string, lastname string, mail string, password string) int {
	database, _ := sql.Open("sqlite3", "./db.db")
	var preparement string = "SELECT * FROM persons WHERE mail='" + mail + "'"

	rows, _ := database.Query(preparement)

	if rows.Next() {
		//not unique e-mail
		return 1
	} else {
		statement, _ := database.Prepare("INSERT INTO persons (firstname, lastname, mail, password) VALUES (?, ?, ?, ?)")
		statement.Exec(firstname, lastname, mail, password)
		//0 = succesful operation
	}

	return 0
}

func Getposts() []Postinfo {
	database, _ := sql.Open("sqlite3", "./db.db")
	var preparement string = "SELECT * FROM posts"

	var size int = 0
	var slice = make([]Postinfo, size)
	var info Postinfo
	rows, _ := database.Query(preparement)

	for rows.Next() {
		err := rows.Scan(&info.firstname, &info.lastname, &info.date, &info.mail)
		if err != nil {
			log.Fatal(err)
		}
		slice = append(slice, info)
	}
	return slice
}
