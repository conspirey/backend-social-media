package message

import (
	"main/functions/sessions"
	"main/structs"

	sock "github.com/ambelovsky/gosf-socketio"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)
type Text struct {
	Text string `json:"text"`
}
func CreateMessage(c *gin.Context, f *sock.Server, db *mongo.Database) {
	session := sessions.Default(c)
	var Text Text;
	user := session.Get("user")
	if user != nil {
		c.JSON(401, Error("Not Authorized", "not_authorized_3"))
		return
	}
	if err := c.ShouldBindJSON(&Text); err != nil {
		c.JSON(400, gin.H{
			"error_message": "invalid json provided",
			"error": "invalid_json",
		})
	}
	var typeC string = c.Query("type")
	var msgT = structs.NewMessage(Text.Text, session)
	if err := msgT.Apply(typeC); err != nil {
		c.JSON(400, Error(err.Error(), "invalid_type_3"))
		return
	}
	
	if msgT.Basic == typeC {
		f.BroadcastTo("chat", "echo", msgT.ToMap())
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
