package main

import (
	"context"
	"encoding/gob"
	"fmt"
	"log"
	"main/functions"
	"main/routes"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	//f "github.com/ambelovsky/gosf"
	"github.com/gin-contrib/sessions/mongo/mongodriver"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
var (
	r = gin.Default();
)
func main() {
	gob.Register(map[string]interface{}{})
	gob.Register(map[interface{}]interface{}{})
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
		c.JSON(200, sessions.Default(c).Get("user"))
	})
	r.GET("/set", func(c *gin.Context) {
		session := sessions.Default(c)
		session.Set("user", map[string]any{
			"password": "123",
			"name": "mrredo",
		})
		fmt.Println(session.Save())
		c.JSON(200, sessions.Default(c).Get("user"))
	})
	r.Run("localhost:8080")

	//defer f.Shutdown()
}