package main

import (
	"github.com/gin-gonic/gin"
	"main/routes"
)
var (
	r = gin.Default();
)
func main() {
	r.Static("/static", "./static")
	routes.LoadWebSocket(r);
	r.Run("localhost:8080")
}