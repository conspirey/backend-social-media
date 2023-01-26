package message

import (
	// "fmt"
	"main/functions/sessions"
	"main/structs"
	"strings"

	sock "github.com/ambelovsky/gosf-socketio"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

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
	msg = *structs.NewMessage(msg.Text, session)
	// var msg = structs.NewMessage()
	if msgT.Basic == typeC {
		// fmt.Println(user.(map[string]any)["id"])
		// id := user.(map[string]any)["id"]
		// msg.User.Name = id.(string) //user.(map[string]any)["name"].(string)
		// msg.User.ID = user.(map[string]any)["id"].(string)

		f.BroadcastTo("chat", "echo", msg.ToMap())
		// fmt.Println(msg.ToMap(), msg, user)
		c.Status(200)
	} else if msgT.Server == typeC {

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
