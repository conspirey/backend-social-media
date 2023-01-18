package message

import (
	"main/functions/sessions"
	"main/structs"

	sock "github.com/ambelovsky/gosf-socketio"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)
func CreateMessage(c *gin.Context, f *sock.Server, db *mongo.Database) {
	session := sessions.Default(c)
	user := session.Get("user")
	if user != nil {
		c.JSON(401, Error("Not Authorized", "not_authorized_3"))
		return
	}
	var typeC string =c.Query("type")
	var msg = structs.MessageType{}
	if err := msg.Apply(typeC); err != nil {
		c.JSON(400, Error(err.Error(), "invalid_type_3"))
		return
	}
	if msg.Basic == typeC {
		f.BroadcastTo("chat")
	} else if msg.Server == typeC {

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
