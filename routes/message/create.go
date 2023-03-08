package message

import (

	mongof "main/functions/mongo"
	"main/functions/sessions"
	"main/structs"
	"strings"

	sock "github.com/ambelovsky/gosf-socketio"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*
Requirements to use
- Valid Cookie header - example - Cookie: user={cookie_string}
- json data { "text": "hello world!"}
- Query type=basic|server
*/
func CreateMessage(c *gin.Context, f *sock.Server, db *mongo.Database) {
	session := sessions.Default(c)
	var msg structs.Message;
	user := session.Get("user")
	if user == nil {
		c.JSON(401, Error("Not Authorized", "not_authorized_3"))
		return
	}
	if err := c.ShouldBindJSON(&msg); err != nil {
		c.JSON(400, gin.H{
			"error_message": "invalid json provided",
			"error": "invalid_json",
		})
		return
	}
	var typeC string = strings.ToLower(c.Query("type"))
	var msgT = structs.MessageType{}
	if err := msgT.Apply(typeC); err != nil {
		c.JSON(400, Error(err.Error() + "| all types: basic, server", "invalid_type_3"))
		return
	}
	// msg.SetUser(session)
	// id, name := user.(map[string]any)["id"].(string), user.(map[string]any)["name"].(string)
	msg = *structs.NewMessage(msg.Text,typeC, session)
	// var msg = structs.NewMessage()
	if msgT.Basic == typeC {
		// fmt.Println(user.(map[string]any)["id"])
		// id := user.(map[string]any)["id"]
		// msg.User.Name = id.(string) //user.(map[string]any)["name"].(string)
		// msg.User.ID = user.(map[string]any)["id"].(string)

		f.BroadcastTo("chat", "echo", msg.ToMap())
		// fmt.Println(msg.ToMap(), msg, user)
		c.JSON(200, gin.H{
			"success": "message_is_sent",
		})
		c.Status(200)
	} else if msgT.Server == typeC {	
		data, err := mongof.FindOne(bson.M{"id":msg.User.ID}, options.FindOne(), db, "user")
		if err != nil {
			c.JSON(500, gin.H{
				"error": "server_database_error",
			})
			return
		}
		user := &structs.User{}
		user.MapToUser(data)
		if !user.Admin {
			c.JSON(400, gin.H{
				"error": "not_an_admin",
				"error_message": "You have to be admin to continue",
			})
		}
		f.BroadcastTo("chat", "echo", msg.ToMap())

		c.JSON(200, gin.H{
			"success": "message_is_sent",
		})
		
	}
}

func Error(text string, err string) gin.H {
	return gin.H{
		"error": err,
		"error_message": text,
	}
}
func Success(text string) gin.H {
	return gin.H{
		"success": text,
	}
}
