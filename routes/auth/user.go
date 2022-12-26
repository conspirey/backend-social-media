package auth

import (
	"fmt"
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
		return
	}
	if !user.IsValidName() {
		fmt.Println()
		c.JSON(400, Error("Invalid username"))
		return
	}
	if !user.IsValidPass() {
		c.JSON(400, Error("Invalid password, it should be 8-32 in length"))
		return
	}
	
	if err := user.RegisterAccount(user.Name, user.Password, db); err != nil {
		c.JSON(400, Error(err.Error()))
	}
	c.JSON(200, Error("Created your account"))
	
}
func Error(text string) gin.H {
	return gin.H{
		"error": text,
	}
}