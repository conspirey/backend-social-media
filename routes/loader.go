package routes

import (
	gosocketio "github.com/ambelovsky/gosf-socketio"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"main/routes/auth"
	"main/routes/blog"
	"main/routes/message"
	"main/routes/updatedata"
	"main/routes/user"
)

func LoadRoutes(r *gin.Engine, db *mongo.Database, server *gosocketio.Server) {

	auths := r.Group("/auth/")
	api := r.Group("/api/")
	api.GET("/user", func(c *gin.Context) {
		auth.GetUserData(c, db)
	})
	api.GET("/encryption", func(ctx *gin.Context) {
		updatedata.UpdateUserData(ctx, db)
	})
	api.GET("/blog", func(c *gin.Context) {
		blog.GetBlogs(c, db)
	})
	api.PATCH("/user/ban", func(c *gin.Context) {
		user.BanUser(c, db)
	})
	api.POST("/blog", func(c *gin.Context) {
		blog.CreateBlog(c, db)
	})
	api.PATCH("/user/admin", func(c *gin.Context) {
		user.SetAdmin(c, db)
	})
	r.POST("auth/register", func(c *gin.Context) {
		auth.Register(c, db)
	})
	api.POST("/message", func(c *gin.Context) {
		message.CreateMessage(c, server, db)
	})

	auths.POST("/login", func(c *gin.Context) {
		auth.Login(c, db)
	})
	auths.POST("/logout", func(c *gin.Context) {
		auth.Logout(c, db)
	})

}
