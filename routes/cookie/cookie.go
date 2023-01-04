package cookie

import (
	// "fmt"
	"main/functions/security"
	"main/functions/sessions"
	"main/structs"

	"github.com/gin-gonic/gin"
)

//GET REQUEST
func Cookie(c *gin.Context) {

	GetHeaderAndSetCookie(c)
	c.JSON(200, gin.H{
		"success": "Added cookie",
	})
}
func SetHeaderCookie(value string, c *gin.Context) {
	c.Header("Set-Cookie", value)
}
func GetHeaderAndSetCookie(c *gin.Context) {
	var key = structs.Key
	val := c.GetHeader("Set-Cookie")
	if val != "" {
		data, _ := security.Decrypt(val, key)
		session := sessions.Default(c)
		c.SetCookie(session.Name, data, session.Options.MaxAge, session.Options.Path, session.Options.Domain, session.Options.Secure, session.Options.HttpOnly) 
	}
}