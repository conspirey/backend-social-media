package auth

import (
	mongof "main/functions/mongo"
	mses "main/functions/sessions"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
)
func GetUserData(c *gin.Context, db *mongo.Database) {
	session := mses.Default(c)
	user := session.Get("user")
	if user != nil {
		// mongof.FindOne(bson.M{}, options.FindOne(), db, "user")
		c.JSON(200, user)
	} else {
		c.JSON(401, gin.H{
			"error": "user_not_logged_in", 
			"error_message": "User is not logged in",
		})
	}
}