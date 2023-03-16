package blog

import (
	"main/structs"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateBlog(c *gin.Context, db *mongo.Database) {
	blog := &structs.Blog{}
	if err := c.ShouldBindJSON(&blog); err != nil {
		c.JSON(400, Error("invalid json body", "invalid_json"))
		return
	}
	c.JSON(200, blog)

}