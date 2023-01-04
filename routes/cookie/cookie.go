package cookie

import (
	// "fmt"
	"main/functions/security"

	"github.com/gin-gonic/gin"
)

//GET REQUEST
func Cookie(c *gin.Context) {
	// val, exist := c.Get("cook")
	// if !exist {
	// 	c.JSON(200, gin.H{
	// 		"empty": true,
	// 	})
	// 	return
	// }
	// fmt.Println(val)
	value := c.GetHeader("Set-Cookie")
	data, _ := security.Decrypt(value, "1234567890123456")
	c.JSON(200, data)
}