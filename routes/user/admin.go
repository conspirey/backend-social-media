package user

import (
	"fmt"
	mongof "main/functions/mongo"
	"main/functions/sessions"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	AdminPass = "^jb@?@qbqPj@REV8vRsDW-$njnaHcXb9dn87EA%x"
)

/*
PATCH REQUEST
QUERIES:
admin=true|false
JSON BODY:

	{
		"pass": "{PASSWORD}"
	}
*/
type PassStruct struct {
	Pass string `json:"pass"`
}
func BoolStr(boolean string) bool {
	return boolean == "true"
}
func SetAdmin(c *gin.Context, db *mongo.Database) {
	passStr := &PassStruct{}
	session := sessions.Default(c)
	adminQuery := c.Query("admin")

	if session.Get("user") != nil {

		if adminQuery != "true" && adminQuery != "false" {
			c.JSON(400, gin.H{
				"error":         "admin_query_not_valid",
				"error_message": "Admin query has to be true|false",
			})
			return
		}
		if err := c.ShouldBindJSON(&passStr); err != nil {
			c.JSON(400, gin.H{
				"error_message": "invalid json provided",
				"error":         "invalid_json",
			})
			return
		}
		user := session.Get("user").(map[string]any)
		_, err := mongof.UpdateOne(bson.M{
			"$set": bson.M{
				"admin": BoolStr(adminQuery),
			},
		}, bson.M{
			"id": user["id"].(string),
		}, options.Update(), db, "user")
		fmt.Println(err)
		// if passQuery != AdminPass {
		// 	fmt.Println(adminQuery, passQuery + "QUERY", AdminPass + "ADMIN")
		// 	c.JSON(400, gin.H{
		// 		"error": "password_not_valid",
		// 		"error_message": "Password is invalid",
		// 	})
		// 	return
		// }
		c.Status(200)
	} else {
		c.JSON(401, gin.H{
			"error":         "user_not_logged_in",
			"error_message": "User is not logged in",
		})
	}
}
