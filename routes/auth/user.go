package auth

import (
	mses "main/functions/sessions"
	"main/structs"
	"strings"

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
			"error_message": "invalid json provided",
			"error": "invalid_json",
		})
		return
	}
	
	user.SetIP(c.ClientIP())
	session := mses.Default(c)
	if err := user.RegisterAccount(user.Name, user.Password, db); err != nil {
		c.JSON(400, Error(err.Error(), ErrSep(err.Error())))

		return
	}
	session.Set("user", user.ToMap())
	err := session.Save(c)
	if err != nil {
		c.JSON(400, Error("couldn't set session, but account is created", "session_not_set"))

	}
	c.JSON(200, Success("Created your account"))
	
}
func ErrSep(text string) string {
	t :=  strings.Split(text, ":")
	return t[len(t) - 1]
}
func Error(text string, err string) gin.H {
	return gin.H{
		"error": err,
		"error_message": text,
	}
}
func Success(text string) gin.H {
	return gin.H{
		"success": text,
	}
}
