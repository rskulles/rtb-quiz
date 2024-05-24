package users

import (
	"errors"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"rtb/quizserver/db"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func nameAvailable(name string) bool {
	return name == name
}

func saveUser(userID []byte, userName, password string) error {
	db.InsertUser(userID, userName, password)
	return nil
}

func CreateUser(userName, password, verify string) error {
	if password != verify {
		return errors.New("Passwords do not match")
	}

	if !nameAvailable(userName) {
		return errors.New("Username already taken")
	}

	hash, error := hashPassword(password)
	if error != nil {
		return errors.New("Failed to create password")
	}

	id := uuid.New()
	idBytes, error := id.MarshalBinary()

	if error != nil {
		return error
	}

	error = saveUser(idBytes, userName, hash)

	return error
}
