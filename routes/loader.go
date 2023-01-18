package routes

import (
	"main/functions/security"
	"main/routes/auth"
	"main/routes/cookie"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func LoadRoutes(r *gin.Engine, db *mongo.Database) {
	auths := r.Group("/auth/")
	api := r.Group("/api/")
	api.GET("/user", auth.GetUserData)
	r.POST("auth/register", func(c *gin.Context) {
		auth.Register(c, db)
	})
	
	//REMOVE THIS SOON
	//REMOVE THIS SOON
	//REMOVE THIS SOON
	r.POST("/cookie", func(c *gin.Context) {
		data, _ := security.Encrypt("{\"name\": \"Hello World!\"}", "1234567890123456")
		cookie.SetHeaderCookie(data, c)
		c.JSON(200, "make request to /cookie GET")
		// c.Redirect(307, "/cookie")
	})
	//REMOVE THIS SOON
	//REMOVE THIS SOON

//REMOVE THIS SOON
	r.GET("/cookie", func(c *gin.Context) {
		cookie.Cookie(c)
		// session := sessions.Default(c)
		// if session.Get("user") == nil {
		// 	s, _ := security.Encrypt("{}", session.EncrKey)
		// 	c.SetCookie(session.Name, s, session.Options.MaxAge, session.Options.Path, session.Options.Domain, session.Options.Secure, session.Options.HttpOnly)
		// 	c.JSON(200, "Created cookie")
		// 	return
		// } 
		// c.JSON(200, "Cookie Exists")
	})
	auths.POST("/login", func(c *gin.Context) {
		auth.Login(c, db)
	})
	auths.POST("/logout", func(c *gin.Context) {
		auth.Logout(c, db)
	})

}