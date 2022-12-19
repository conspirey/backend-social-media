package main

import (
	"github.com/gin-gonic/gin"
)
var (
	r = gin.Default();
)
func main() {
	r.Static("/static", "./static")
	r.Run(":8080")
}