package db

import (
	"database/sql"
	_ "github.com/glebarez/go-sqlite"
	"github.com/google/uuid"
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

func VerifyAction(actionId, userId string) bool {

	return true
}

func SessionActive(session string) (bool, string) {
	db, err := open()
	if err != nil {
		return false, ""
	}

	sessionId, err := uuid.Parse(session)

	if err != nil {
		return false, ""
	}

	sessionBytes, err := sessionId.MarshalBinary()

	if err != nil {
		return false, ""
	}

	rows, err := db.Query("SELECT session_id, user_id FROM session WHERE session_id = ? AND active = 1", sessionBytes)

	if err != nil {
		return false, ""
	}

	dbSession := ""
	dbUser := ""
	for rows.Next() {

		cols, err := rows.Columns()

		if err != nil {
			return false, ""
		}

		dbSessionString := cols[0]
		dbBytes := []byte(dbSessionString)
		dbUUID, err := uuid.ParseBytes(dbBytes)
		dbSession = dbUUID.String()

		if err != nil {
			return false, ""
		}

		dbUserString := cols[1]
		dbUserBytes := []byte(dbUserString)
		dbUserId, err := uuid.ParseBytes(dbUserBytes)

		if err != nil {
			return false, ""
		}

		dbUser = dbUserId.String()
	}

	return dbSession == session, dbUser
}
