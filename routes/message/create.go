package message

import (
	"fmt"
	mongof "main/functions/mongo"
	"main/functions/sessions"
	"main/structs"
	"strconv"
	"strings"
	"time"

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
- Query type=basic|server|server_timer
*/
func CreateMessage(c *gin.Context, f *sock.Server, db *mongo.Database) {
	session := sessions.Default(c)
	var msg structs.Message = structs.Message{ServerMessage: &structs.ServerMessage{}}
	user := session.Get("user")
	if user == nil {
		c.JSON(401, Error("Not Authorized", "not_authorized_3"))
		return
	}
	if err := c.ShouldBindJSON(&msg); err != nil {
		c.JSON(400, gin.H{
			"error_message": "invalid json provided",
			"error":         "invalid_json",
		})
		return
	}
	var typeC string = strings.ToLower(c.Query("type"))
	var msgT = structs.MessageType{}
	if err := msgT.Apply(typeC); err != nil {
		c.JSON(400, Error(err.Error()+"| all types: basic, server, server_timer", "invalid_type_3"))
		return
	}
	// msg.SetUser(session)
	// id, name := user.(map[string]any)["id"].(string), user.(map[string]any)["name"].(string)
	msg = *structs.NewMessage(msg.Text, typeC, session, &structs.ServerMessage{
		Timer:   msg.ServerMessage.Timer,
		Delay:   msg.ServerMessage.Delay,
		Message: msg.ServerMessage.Message,
	})

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
		data, err := mongof.FindOne(bson.M{"id": msg.User.ID}, options.FindOne(), db, "user")
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
				"error":         "not_an_admin",
				"error_message": "You have to be admin to continue",
			})
			return
		}
		f.BroadcastTo("chat", "echo", msg.ToMap())

		c.JSON(200, gin.H{
			"success": "message_is_sent",
		})

	} else if msgT.ServerTimer == typeC {
		if msg.ServerMessage.Timer == 0 || msg.ServerMessage.Timer > 100 {
			c.JSON(400, Error("timer has to be more than 0 and cant be greater than 100", "invalid_timer"))
			return
		}
		if msg.ServerMessage.Delay < -1 || msg.ServerMessage.Delay > 10000 {
			c.JSON(400, Error("Delay must be from 0ms to 10000ms", "invalid_delay"))
			return
		}
		data, err := mongof.FindOne(bson.M{"id": msg.User.ID}, options.FindOne(), db, "user")
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
				"error":         "not_an_admin",
				"error_message": "You have to be admin to continue",
			})
			return
		}
		matches := structs.ServerMessageTimerRegex.FindAllString(msg.Text, -1)
		ogString := msg.Text
		if len(matches) > 0 {
			for i := msg.ServerMessage.Timer; i >= 1; i-- {
				//matchesS := structs.ServerMessageTimerRegex.FindAllString(msg.Text, -1)

				outputString := msg.Text
				for _, match := range matches {
					outputString = strings.Replace(ogString, match, strconv.Itoa(int(i)), 1)

				}

				time.Sleep(time.Duration(msg.ServerMessage.Delay) * time.Millisecond)
				msg.Text = outputString
				msg.Type = "server"
				fmt.Println()
				f.BroadcastTo("chat", "echo", msg.ToMap())
			}
		} else {
			c.JSON(400, Error("You need to use timers to use this message type", "no_timer_found"))
			return
		}
		//for i:= 0;

		c.JSON(200, gin.H{
			"success": "message_is_sent",
		})
	}
}

func Error(text string, err string) gin.H {
	return gin.H{
		"error":         err,
		"error_message": text,
	}
}
func Success(text string) gin.H {
	return gin.H{
		"success": text,
	}
}
