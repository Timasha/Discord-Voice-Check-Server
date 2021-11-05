package handlers

import (
	"discordMsgRead/internal/utils"
	"encoding/json"
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
			data, marshErr := json.Marshal(value)
			if marshErr != nil {
				log.Printf("Marshal user status error: %v", marshErr)
			}
			writeErr := conn.WriteMessage(websocket.TextMessage, data)
			if writeErr != nil {
				conn.Close()
			}
		}
		utils.ConnSlice = append(utils.ConnSlice, conn)
	}
}
