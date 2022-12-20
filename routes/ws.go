package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func LoadWebSocket(r *gin.Engine) {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin:     func(r *http.Request) bool { return true },
	}

	// Create a handler function for the WebSocket endpoint
	r.GET("/ws", func(c *gin.Context) {
		// Upgrade the HTTP connection
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer conn.Close()

		// Start listening for messages
		for {
			// Read a message from the client
			_, msg, err := conn.ReadMessage()
			if err != nil {
				fmt.Println(err)
				break;
			}
			fmt.Println(string(msg))

			// Send a message back to the client
			if err = conn.WriteMessage(websocket.TextMessage, []byte("Hello from the server!")); err != nil {
				fmt.Println(err)
				break
			}
		}
	})

}