package handlers

import (
	"discordMsgRead/internal/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader websocket.Upgrader = websocket.Upgrader{}

func SocketHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		upgrader.CheckOrigin = func(h *http.Request) bool { return true }
		conn, connErr := upgrader.Upgrade(c.Writer, c.Request, nil)
		if connErr != nil {
			log.Printf("Websocket handle connection error: %v", connErr)
		}
		conn.WriteMessage(websocket.TextMessage, utils.GetChannelMessage)
		for _, value := range utils.UserStatuses {
			writeErr := conn.WriteMessage(websocket.TextMessage, value)
			if writeErr != nil {
				conn.Close()
			}
		}
		utils.ConnSlice[len(utils.ConnSlice)] = conn
	}
}
