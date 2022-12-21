package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"

	f "github.com/ambelovsky/gosf"
)

func LoadWebSocket(r *gin.Engine) {
	

	// f.Listen("echo", func(client *f.Client, request *f.Request) *f.Message {
	// 	f.Broadcast("", request.Endpoint, f.NewSuccessMessage("hello"))
	// 	fmt.Print(request.Endpoint)
	// 	return f.NewSuccessMessage(request.Message.Text)
	//   })
	
	f.OnConnect(func(client *f.Client, request *f.Request) {
		// f.Broadcast()
		client.Join("test")
		f.Broadcast("", "echo", f.NewSuccessMessage("hellkyrdyjrsjyjsryjsryjsrsyrzjo"))
		fmt.Println("connected to client")
		
	})
	f.Listen("echo", func(client *f.Client, request *f.Request) *f.Message {
		
		fmt.Print(client.Rooms)
		
		return f.NewSuccessMessage("j")
	})
	fmt.Println("hello")
	f.Startup(map[string]interface{}{
		"port": 3000})
}