package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
	// "github.com/gorilla/securecookie"

	f "github.com/ambelovsky/gosf-socketio"
	"github.com/ambelovsky/gosf-socketio/transport"
)
type Echo struct {
	Text string `json:"text"`
}
func LoadWebSocket(r *gin.Engine, keypair []byte) *f.Server {
	server := f.NewServer(transport.GetDefaultWebsocketTransport())
	
	// f.Listen("echo", func(client *f.Client, request *f.Request) *f.Message {
	// 	f.Broadcast("", request.Endpoint, f.NewSuccessMessage("hello"))
	// 	fmt.Print(request.Endpoint)
	// 	return f.NewSuccessMessage(request.Message.Text)
	//   })
	
	server.On(f.OnConnection, func(c *f.Channel) {
		fmt.Println("New client connected")
		// if len(c.Request().Cookies()) > 0 {
		// 	// cook, _ := c.Request().Cookie("Cookie")
		// 	// var keypairs = [][]byte{keypair}
			
		// 	// codecs := securecookie.CodecsFromPairs(keypairs...)
		// 	// var value =  map[string]any{}
		// 	// fmt.Println(value == nil, codecs == nil)
			
		// 	// securecookie.DecodeMulti("mysession", cook.Value, &value, codecs...)
			
		// }

		
		
		
 		
		// session.Set("user", map[string]any{
		// 	"password": "123",
		// 	"name": "mrredo",
		// })
		// fmt.Println(session.Save())
		//join them to room
		c.Join("chat")

		
	})
	server.On("echo", func(c *f.Channel, msg Echo) any {
		//send event to all in room
		// server.BroadcastTo("chat", "echo", map[string]interface{}{
		// 	"text": "hello",
		// })
		server.BroadcastTo("chat", "echo", map[string]interface{}{
					"text": msg.Text,
				})
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