package user

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"main/functions/sessions"
)

func BanUser(c *gin.Context, db *mongo.Database) {
	session := sessions.Default(c)
	user := session.Get("user")
	idToBan := c.Query("id")
	if user == nil {
		c.JSON(401, gin.H{
			"error":         "user_not_logged_in",
			"error_message": "User is not logged in",
		})
		return
	}
	c.String(200, idToBan)

}
