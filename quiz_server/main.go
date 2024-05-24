package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"rtb/quizserver/db"
	"rtb/quizserver/users"
)

func postUsers(c *gin.Context) {
	user := users.CreateUserArgs{}
	if err := c.BindJSON(&user); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	if err := users.CreateUser(user.Name, user.Password, user.PasswordVerify); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	c.Header("Access-Control-Allow-Origin", "*")
	c.IndentedJSON(http.StatusCreated, gin.H{"status": "user created"})
}

func main() {
	homeDir := os.Getenv("HOME")
	dbPath := fmt.Sprintf("%s/Data/database", homeDir)
	db.Init(dbPath)

	router := gin.Default()
	router.POST("/users", postUsers)
	router.Use(cors.Default())
	err := router.Run(":3000")
	if err != nil {
		fmt.Println(err)
	}
}
