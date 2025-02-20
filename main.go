package main

import (
	"context"
	"encoding/gob"
	"github.com/joho/godotenv"
	"os"
	"strings"

	"log"
	"main/functions"
	"main/routes"

	"net/http"
	"time"

	mongof "main/functions/mongo"
	mses "main/functions/sessions"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	keypair = functions.RandStringRunes(32)
)

func main() {
	// gin.SetMode(gin.DebugMode)

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
	r := gin.Default()
	gob.Register(map[string]interface{}{})
	gob.Register(map[interface{}]interface{}{})
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI(os.Getenv("MONGO")).
		// ApplyURI("mongodb+srv://redobot:dbuserpassword123@cluster0.rhc8q.mongodb.net/Cluster0?retryWrites=true&w=majority").
		SetServerAPIOptions(serverAPIOptions)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	var db = client.Database("conspir_test")
	if len(os.Args) > 1 {
		if os.Args[2] == "true" {
			db = client.Database("conspir")
		}
	}
	if !mongof.CollectionExists("user", db) {
		db.CreateCollection(context.TODO(), "user")
	}

	r.Use(cors.New(cors.Config{
		//  AllowAllOrigins: true,
		AllowWildcard:          true,
		AllowWebSockets:        true,
		AllowBrowserExtensions: true,
		AllowOrigins:           []string{"http://localhost:5173", "http://127.0.0.1:3000", "http://127.0.0.1:5173", "https://127.0.0.1:443", "http://127.0.0.1:80", "http://192.168.8.114:5173", "https://conspirey.xyz", "http://vm5.spacerv.ovh:3623/", "http://localhost:3623", "http://127.0.0.1:3623"},
		AllowMethods:           []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:           []string{"Origin", "Content-Length", "Content-Type", "Accept", "Cookie", "Set-Cookie"},
		AllowCredentials:       true,
	}))
	r.Use(mses.MiddleWare("user", keypair, 3600*48, "", "/", false, false, http.SameSiteNoneMode))
	// cs := db.Collection("sessions")
	// store := mongodriver.NewStore(cs, 3600*48, true, []byte(keypair)) // change 3600 time how to: delete everything in mongodb collection

	// r.Use(sessions.Sessions("mysession", store))

	server := routes.LoadWebSocket(r, keypair)
	routes.LoadRoutes(r, db, server)

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

	// if os.Getenv("GIN_MODE") == "release" {
	// 	listener, err := net.Listen("tcp4", "0.0.0.0:3200")
	// 	if err != nil {
	// 		// handle error
	// 	}
	// 	r.RunListener(listener)
	// } else {
	// 	r.Run(":3200")
	// }
	// if err := CMD(r).Execute(); err != nil {
	// 	fmt.Fprintln(os.Stderr, err)
	// 	os.Exit(1)
	//   }

	if len(os.Args) > 1 {
		if strings.ToLower(os.Args[1]) == "release" {
			gin.SetMode(gin.ReleaseMode)
			r.RunTLS(":"+os.Getenv("PORT"), "./ssl/cert.pem", "./ssl/keys.pem")
		} else {
			gin.SetMode(gin.DebugMode)
			r.Run("localhost:" + os.Getenv("PORT"))
		}

	} else {
		gin.SetMode(gin.DebugMode)
		r.Run(":" + os.Getenv("PORT"))
	}

	// r.Run(":3100")
	//defer f.Shutdown()
}
func RunHTTPHandler(h http.Handler, c *gin.Context) {
	h.ServeHTTP(c.Writer, c.Request)
}

// func CMD(r *gin.Engine) *cobra.Command {
// 	var rootCMD = &cobra.Command{
// 		Use: "conspirey",
// 		Short: "Help is to help you",

// 		Run: func(c *cobra.Command, args []string) {

// 			if args[0] == "release" {
// 				listener, err := net.Listen("tcp4", "0.0.0.0:3200")
// 				if err != nil {

// 				}
// 				r.RunListener(listener)
// 			} else {
// 				r.Run(":3200")
// 			}

// 		},
// 	}
// 	return rootCMD
// }
