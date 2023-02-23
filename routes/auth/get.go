package auth

import (
	mongof "main/functions/mongo"
	mses "main/functions/sessions"
	"main/structs"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
func GetUserData(c *gin.Context, db *mongo.Database) {
	session := mses.Default(c)
	user := session.Get("user")
	if user != nil {
		id := c.Query("id")
		if id == "" {
			id = user.(map[string]any)["id"].(string)
		}
		userM, err := mongof.FindOne(bson.M{
			"id": id,
		}, options.FindOne(), db, "user")
		if err != nil {
			c.JSON(502, gin.H{
				"error": "database_error",
			})
			return
		}
		if len(userM) == 0 {
			c.JSON(404, gin.H{
				"error": "user_not_found",
			})
			return
		}
		structs.StripMapOfImportantInfo(userM)
		c.JSON(200, userM)
	} else {
		c.JSON(401, gin.H{
			"error": "user_not_logged_in", 
			"error_message": "User is not logged in",
		})
	}
}