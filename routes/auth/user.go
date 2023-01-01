package auth

import (
	"main/structs"
	mses "main/functions/sessions"
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
		c.JSON(422, gin.H{
			"error": "invalid json provided",
		})
		return
	}
	
	if !user.IsValidName() {
		c.JSON(422, Error("Invalid username"))
		return
	}
	if !user.IsValidPass() {
		c.JSON(422, Error("Invalid password, it should be 8-32 in length"))
		return
	}
	user.SetIP(c.ClientIP())
	session := mses.Default(c)
	if err := user.RegisterAccount(user.Name, user.Password, db); err != nil {
		c.JSON(422, Error(err.Error()))
		return
	}
	session.Set("user", user.ToMap())
	err := session.Save(c)
	if err != nil {
		c.JSON(422, Error("couldn't set session, but account is created"))
	}
	c.JSON(200, Success("Created your account"))
	
}
func Error(text string) gin.H {
	return gin.H{
		"error": text,
	}
}
func Success(text string) gin.H {
	return gin.H{
		"success": text,
	}
}