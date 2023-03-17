package blog

import (
	"fmt"
	"main/functions/sessions"
	"main/structs"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

/*
	{
	  "id": "blog123",
	  "title": "My First Blog Post",
	  "body": "This is my first blog post. I'm very excited to share my thoughts with the world!",
	  "created_at": "2022-03-15T16:22:00Z",
	  "updated_at": "2022-03-15T16:45:00Z",
	  "author": {
	    "name": "John Doe",
	    "id": "author123"
	  },
	  "tags": [
	    "blog",
	    "first post",
	    "excitement"
	  ]
	}
*/
func CreateBlog(c *gin.Context, db *mongo.Database) {
	session := sessions.Default(c)
	user := session.Get("user")
	if user == nil {
		c.JSON(401, Error("user is not authorized", "not_authorized"))
		return
	}
	userSTR := &structs.User{}
	userSTR.ID = user.(map[string]any)["id"].(string)
	if !userSTR.AccountExists(db) {
		c.JSON(404, Error("user account does not exist", "user_not_exist"))
		return
	}
	errF := userSTR.FetchData(db)
	if errF != nil {
		fmt.Println(errF)
		c.JSON(400, Error("failed fetching data", "failed_fetching_data"))
		return
	}

	blog := &structs.Blog{}
	if err := c.ShouldBindJSON(&blog); err != nil {
		c.JSON(400, Error("invalid json body", "invalid_json"))
		return
	}
	if blog.Title == "" || blog.Body == "" || blog.Description == "" {
		c.JSON(400, Error("Title, Body, Description fields can not be empty", "empty_text"))
		return
	}
	timeNow := time.Now().Unix()
	blog.CreatedAt = timeNow
	blog.UpdatedAt = timeNow
	blog.GenerateID()
	if !blog.ValidID() {
		c.JSON(400, Error("try again, failed generating id", "failed_generating_id"))
		return
	}
	blog.Author.ID = userSTR.ID
	blog.Author.Name = userSTR.Name
	if err := blog.Create(db); err != nil {
		c.JSON(400, Error("failed inserting blog to database", "database_error"))
		return
	}
	c.JSON(200, blog)

}
