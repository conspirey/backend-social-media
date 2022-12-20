package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"

)

func LoadWebSocket(r *gin.Engine) {
	server := socketio.NewServer(nil)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	server.On("connection", func(so socketio.Socket) {
		fmt.Println("on connection")
		so.Join("chat")
		so.On("chat message", func(msg string) {
			fmt.Println("emit:", so.Emit("chat message", msg))
			so.BroadcastTo("chat", "chat message", msg)
		})
		so.On("disconnection", func() {
			fmt.Println("on disconnect")
		})
	})
	server.OnConnect("error", func(so socketio.Socket, err error) {
		fmt.Println("error:", err)
	})
	r.GET("/socket.io/*any", gin.WrapH(server))
	r.POST("/socket.io/*any", gin.WrapH(server))

//https://github.com/googollee/go-socket.io/blob/master/_examples/gin-gonic/main.go
}