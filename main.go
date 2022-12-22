package main

import (
	"github.com/gin-gonic/gin"
	"main/routes"
	//f "github.com/ambelovsky/gosf"
)
var (
	r = gin.Default();
)
func main() {
	server := routes.LoadWebSocket(r)
	r.Static("/static", "./static")
	r.GET("/socket.io/", gin.WrapH(server))
	r.POST("/socket.io/", gin.WrapH(server))
	//r.GET("/socket.io/", gin.WrapH(f))
	r.GET("/", func(c *gin.Context) {
		c.String(200, "hello")
	})
	r.Run("localhost:8080")

	//defer f.Shutdown()
}