package auth

import (
	mses "main/functions/sessions"

	"github.com/gin-gonic/gin"
)
func GetUserData(c *gin.Context) {
	session := mses.Default(c)
	user := session.Get("user")
	if user != nil {
		c.JSON(200, user)
	} else {
		c.JSON(401, gin.H{
			"error": "user_not_logged_in", 
			"error_message": "User is not logged in",
		})
	}
}