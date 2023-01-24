package main

import (
	"context"
	"encoding/gob"
	"os"
	"net"

	// "fmt"
	"log"
	"main/functions"
	"main/routes"

	// "main/routes/auth"
	"net/http"
	"time"

	// "github.com/gin-contrib/sessions"
	mongof "main/functions/mongo"
	mses "main/functions/sessions"

	"github.com/gin-contrib/cors"
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
	r.Use(cors.New(cors.Config{
		 AllowAllOrigins: true,
		 AllowWildcard: true,
		 AllowWebSockets: true,
		 AllowBrowserExtensions: true,
		//AllowOrigins: []string{"http://localhost:5173", "http://127.0.0.1:3000", "http://127.0.0.1:5173", "http://192.168.8.114:5173"},
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders: []string{"Origin","Content-Length", "Content-Type", "Accept", "Cookie", "Set-Cookie"},
		AllowCredentials: true,
		
	}))
	r.Use(mses.MiddleWare("user", keypair, 3600*48, "", "/", false, false, http.SameSiteNoneMode))
	// cs := db.Collection("sessions")
	// store := mongodriver.NewStore(cs, 3600*48, true, []byte(keypair)) // change 3600 time how to: delete everything in mongodb collection
	
	// r.Use(sessions.Sessions("mysession", store))


	server := routes.LoadWebSocket(r, keypair)
	routes.LoadRoutes(r, db, server)
	r.GET("/e", func(c *gin.Context) {
		server.BroadcastTo("chat", "echo", map[string]any{
			"text": "BOZOZZOO",
			"user": map[string]any{
				"name": "mrredo",
			},
		})
	})
	
	r.Static("/assets/", "./frontend/dist/assets")
	r.Static("/static/", "./frontend/dist")
	r.GET("/socket.io/", func(c *gin.Context) {
		RunHTTPHandler(server, c)
	})
	
	r.POST("/socket.io/", gin.WrapH(server))
	//r.GET("/socket.io/", gin.WrapH(f))


	r.NoRoute(func(c *gin.Context) {
		c.File("./frontend/dist/index.html")
	})


	if os.Getenv("GIN_MODE") == "release" {
		listener, err := net.Listen("tcp4", "0.0.0.0:3200")
		if err != nil {
			// handle error
		}
		r.RunListener(listener)
	} else {
		r.Run(":3200")
	}


	//defer f.Shutdown()
}
func RunHTTPHandler(h http.Handler, c *gin.Context) {
	h.ServeHTTP(c.Writer, c.Request)
}
