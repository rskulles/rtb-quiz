package db

import (
	"database/sql"
	_ "github.com/glebarez/go-sqlite"
)

var connection = ""

func Init(connectionString string) {
	connection = connectionString
}

func open() (*sql.DB, error) {
	db, err := sql.Open("sqlite", connection)
	err = db.Ping()
	return db, err
}

func InsertUser(userid []byte, userName, password string) error {
	db, err := open()
	if err != nil {
		return err
	}

	_, err = db.Exec("INSERT INTO user (id, name, password) VALUES (?,?,?)", userid, userName, password)

	return err
}

func FindName(userName string) (bool, error) {

	db, err := open()

	if err != nil {
		return false, err
	}

	rows, err := db.Query("SELECT name FROM user WHERE name = ?", userName)

	count := 0

	for rows.Next() {
		count++
	}

	return (count > 0), err
}
