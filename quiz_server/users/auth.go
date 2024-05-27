package users

import (
	"net/http"
	"rtb/quizserver/db"

	"github.com/gin-gonic/gin"
)

func Authorize(f gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := c.Cookie("session")

		if err != nil {
			c.JSON(http.StatusUnauthorized, nil)
		}
		isActive, userId := db.SessionActive(id)

		if !isActive {
			c.JSON(http.StatusUnauthorized, nil)
		}

		actionId := c.GetHeader("action-id")
		actionIsGood := db.VerifyAction(actionId, userId)

		if !actionIsGood {
			c.JSON(http.StatusUnauthorized, nil)
		}
		f(c)
	}
}
