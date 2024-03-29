package blog

import (
	"main/structs"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GetBlogs
/*
Get blog posts

- Queries:

- limit - how many blogs are returned - 0 or undefined will return all blogs



*/
type BlogJSONOptions struct {
	Description string `json:"description"`
}

func GetBlogs(c *gin.Context, db *mongo.Database) {
	limit, _ := strconv.Atoi(c.Query("limit"))
	// author := c.Query("author")
	// id := c.Query()
	var filter = map[string]any{}


	// if len(body) != 0 {
	// 	if errJS := c.BindJSON(&filter); errJS != nil {
	// 		fmt.Println(errJS, filter)
	// 		c.JSON(400, Error("failed binding json", "json_error"))
	// 		return
	// 	}
	// }
	option := options.Find()
	if limit != 0 {
		option.SetLimit(int64(limit))
	}

	data, err := structs.GetBlogs(filter, option, db)
	if err != nil {
		c.JSON(400, Error(err.Error(), "error"))
		return
	}
	if len(data) == 0 {
		c.JSON(200, []bson.M{})
		return
	}
	c.JSON(200, data)

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
