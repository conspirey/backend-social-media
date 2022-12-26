package auth

import (
	"main/structs"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func Login(c *gin.Context, db *mongo.Database) {

}
func Logout(c *gin.Context, db *mongo.Database) {
	
}
func Register(c *gin.Context, db *mongo.Database) {
	var user structs.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{
			"error": "invalid json provided",
		})
	}
	
}