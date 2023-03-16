package blog

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"main/structs"
	"strings"
)

// GetBlogs
/*
Get blog posts

- Queries:

- limit - how many blogs are returned - 0 or undefined will return all blogs



*/
func GetBlogs(c *gin.Context, db *mongo.Database) {
	blog := structs.Blog{}
	user := structs.User{}
	idChan := make(chan string)
	id2 := make(chan string)
	go func() {
		id := blog.GenerateID()
		idChan <- id
	}()
	go func() {
		id := user.CreateID()
		id2 <- id
	}()
	ids1 := <-idChan
	ids2 := <-id2
	c.String(200, fmt.Sprintf("blog: %s \n user: %s", ids1, ids2))
	//session := sessions.Default(c)
	//user := session.Get("user")
	//if user != nil {
	//	userSTR := &structs.User{}
	//	userSTR.MapToUser(user.(map[string]any))
	//	if !userSTR.AccountExists(db) {
	//		c.JSON(404, Error("user not found", "user_not_found"))
	//		return
	//	}
	//	if err := userSTR.FetchData(db); err != nil {
	//		c.JSON(400, Error(err.Error(), "error"))
	//		return
	//	}
	//	userM := userSTR.ToMap()
	//	structs.StripMapOfImportantInfo(userM)
	//	c.JSON(200, userM)
	//
	//} else {
	//	c.JSON(200, Error("User is not logged in", "user_not_logged_in"))
	//}
}
func ErrSep(text string) string {
	t := strings.Split(text, ":")
	return t[len(t)-1]
}
func Error(text string, err string) gin.H {
	return gin.H{
		"error":         err,
		"error_message": text,
	}
}
func Success(text string) gin.H {
	return gin.H{
		"success": text,
	}
}
