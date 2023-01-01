package routes

import (
	"main/routes/auth"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func LoadRoutes(r *gin.Engine, db *mongo.Database) {
	auths := r.Group("/auth/")
	api := r.Group("/api/")
	api.GET("/user/", auth.GetUserData)
	auths.POST("/register/", func(c *gin.Context) {
		auth.Register(c, db)
	})
	auths.POST("/login/", func(c *gin.Context) {
		auth.Login(c, db)
	})
	auths.POST("/logout/", func(c *gin.Context) {
		auth.Logout(c, db)
	})
}