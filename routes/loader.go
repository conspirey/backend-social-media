package routes

import (
	"main/routes/auth"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func LoadRoutes(r *gin.Engine, db *mongo.Database) {
	auths := r.Group("/auth")

	auths.POST("/register", func(c *gin.Context) {
		auth.Register(c, db)
	})
	auths.POST("/login", func(c *gin.Context) {
		auth.Register(c, db)
	})
	auths.POST("/logout", func(c *gin.Context) {
		auth.Register(c, db)
	})
}