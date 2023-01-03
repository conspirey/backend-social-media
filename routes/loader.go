package routes

import (
	"main/functions/security"
	"main/functions/sessions"
	"main/routes/auth"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func LoadRoutes(r *gin.Engine, db *mongo.Database) {
	auths := r.Group("/auth/")
	api := r.Group("/api/")
	api.GET("/user", auth.GetUserData)
	auths.POST("/register", func(c *gin.Context) {
		auth.Register(c, db)
	})
	r.GET("/cookie", func(c *gin.Context) {
		session := sessions.Default(c)
		if session.Get("user") == nil {
			s, _ := security.Encrypt("{}", session.EncrKey)
			c.SetCookie(session.Name, s, session.Options.MaxAge, session.Options.Path, session.Options.Domain, session.Options.Secure, session.Options.HttpOnly)
			c.JSON(200, "Created cookie")
			return
		} 
		c.JSON(200, "Cookie Exists")
	})
	auths.POST("/login", func(c *gin.Context) {
		auth.Login(c, db)
	})
	auths.POST("/logout", func(c *gin.Context) {
		auth.Logout(c, db)
	})
}