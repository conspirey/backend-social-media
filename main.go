package main

import (
	"context"
	"log"
	"main/routes"
	"main/functions"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/go-cmp/cmp/internal/function"

	//f "github.com/ambelovsky/gosf"
	"github.com/gin-contrib/sessions/mongo/mongodriver"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
var (
	r = gin.Default();
)
func main() {
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
    ApplyURI("mongodb+srv://redobot:dbuserpassword123@cluster0.rhc8q.mongodb.net/Cluster0?retryWrites=true&w=majority").
    SetServerAPIOptions(serverAPIOptions)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
    	log.Fatal(err)
	}
	db := client.Database("conspir")
	
	routes.LoadRoutes(r, db)
	cs := db.Collection("sessions")
	store := mongodriver.NewStore(cs, 3600*48, true, []byte(functions.RandStringRunes(15))) // change 3600 time how to: delete everything in mongodb collection
	server := routes.LoadWebSocket(r)
	r.Use(sessions.Sessions("mysession", store))
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