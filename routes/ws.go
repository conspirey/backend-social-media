package routes

import (
	"fmt"
	"main/functions"
	mses "main/functions/sessions"

	"github.com/gin-gonic/gin"

	// "github.com/gorilla/securecookie"

	f "github.com/ambelovsky/gosf-socketio"
	"github.com/ambelovsky/gosf-socketio/transport"
)
type Echo struct {
	Text string `json:"text"`
}
func LoadWebSocket(r *gin.Engine, EncrKey string) *f.Server {
	server := f.NewServer(transport.GetDefaultWebsocketTransport())
	

	
	server.On(f.OnConnection, func(c *f.Channel) {
		fmt.Println("New client connected")

		c.Join("chat")

		
	})
	
	server.On("echo", func(c *f.Channel, msg Echo) any {
		//send event to all in room
		// fmt.Println(msg.Text)
		cookie, err := c.Request().Cookie("user")
		if err == nil {
			
			str, _ := mses.GetDec(cookie.Value, EncrKey)
			data := functions.StringToValue[map[string]any](str)
			
			// fmt.Println(data["user"].(map[string]any)["name"])
			server.BroadcastTo("chat", "echo", map[string]interface{}{
				"text": msg.Text,
				"user": data["user"],
			})
		}
		

		// c.Emit("echo", map[string]interface{}{
		// 		"text": msg.Text,
		// 	})
		return "OK"
	})

	// f.OnConnect(func(client *f.Client, request *f.Request) {
	// 	// f.Broadcast()
	// 	client.Join("test")
	// 	f.Broadcast("", "echo", f.NewSuccessMessage("hellkyrdyjrsjyjsryjsryjsrsyrzjo"))
	// 	fmt.Println("connected to client")
		
	// })
	// f.Listen("echo", func(client *f.Client, request *f.Request) *f.Message {
		
	// 	fmt.Print(client.Rooms)
		
	// 	return f.NewSuccessMessage("j")
	// })
	// fmt.Println("hello")
	// go f.Startup(map[string]interface{}{"port": 8080})
	return server
}