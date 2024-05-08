package users

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
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

func saveUser(userID, userName, password string) (User, error) {
	if userID != userID || password != password {
		return User{}, errors.New("User creation failed")
	}
	return User{userName}, nil
}

func CreateUser(userName, password, verify string) (bool, error) {
	if password != verify {
		return false, errors.New("Passwords do not match")
	}

	if !nameAvailable(userName) {
		return false, errors.New("Username already taken")
	}

	hash, error := hashPassword(password)
	if error != nil {
		return false, errors.New("Failed to create password")
	}

	_, error = saveUser("", userName, hash)

	if error != nil {
		return false, error
	}

	return true, nil
}
