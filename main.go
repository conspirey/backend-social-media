package main

import (
	"context"
	"encoding/gob"
	"fmt"
	"log"
	"main/functions"
	"main/routes"
	"net/http"
	"time"

	// "github.com/gin-contrib/sessions"
	mongof "main/functions/mongo"
	mses "main/functions/sessions"

	"github.com/gin-gonic/gin"

	//f "github.com/ambelovsky/gosf"
	// "github.com/gin-contrib/sessions/mongo/mongodriver"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
var (
	r = gin.Default();
	keypair = functions.RandStringRunes(32)
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
	if !mongof.CollectionExists("user", db) {
		db.CreateCollection(context.TODO(), "user")
	}
	routes.LoadRoutes(r, db)
	// cs := db.Collection("sessions")
	// store := mongodriver.NewStore(cs, 3600*48, true, []byte(keypair)) // change 3600 time how to: delete everything in mongodb collection
	server := routes.LoadWebSocket(r, keypair)
	// r.Use(sessions.Sessions("mysession", store))




	r.Use(mses.MiddleWare("user", keypair, 3600*48))



	r.GET("/test/set", func(c *gin.Context) {
		session := mses.Default(c)
		session.Set("user", map[string]any{
			"name": "hello world!",
		})
		fmt.Println(session.Save(c))
		c.JSON(200, session.Get("user"))
	})
	r.GET("/test/", func(c *gin.Context) {
		session := mses.Default(c)
		// session.Set("user", map[string]any{
		// 	"name": "hello world!",
		// })
		// fmt.Println(session.Save(c))
		c.JSON(200, session.Get("user"))
	})

	r.Static("/static", "./static")
	r.GET("/socket.io/", func(c *gin.Context) {
		RunHTTPHandler(server, c)
	})
	r.POST("/socket.io/", gin.WrapH(server))
	//r.GET("/socket.io/", gin.WrapH(f))




	r.Run("localhost:8080")

	//defer f.Shutdown()
}
func RunHTTPHandler(h http.Handler, c *gin.Context) {
	h.ServeHTTP(c.Writer, c.Request)
}