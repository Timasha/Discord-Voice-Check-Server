package api

import (
	discordconnect "discordMsgRead/internal/api/discordConnect"
	"discordMsgRead/internal/api/handlers"
	"discordMsgRead/internal/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func Middleware(flags utils.Config) {
	header := http.Header{}
	header.Add("Origin", "https://streamkit.discord.com")
	conn, _, connErr := websocket.DefaultDialer.Dial(("ws://127.0.0.1:6463/?v=1&client_id=" + flags.UserID), header)
	if connErr != nil {
		log.Fatalf("Websocket connection error: %v", connErr)
	}
	go discordconnect.RecieveMsg(&conn, flags)
	r := gin.Default()
	r.GET("/", handlers.SocketHandler())
	r.Run("localhost:8080")
}
