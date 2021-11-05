package discordconnect

import (
	"discordMsgRead/internal/utils"

	"github.com/gorilla/websocket"
)

func close(conn *websocket.Conn, message []byte) {
	conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(1000, "woops"))
	conn.Close()
}
func Broadcast(messege []byte, wsSlice []*websocket.Conn) {
	var internal []*websocket.Conn
	internal = append(internal, wsSlice...)
	for i := 0; i < len(internal); i++ {
		err := internal[i].WriteMessage(websocket.TextMessage, messege)
		if err != nil {
			utils.ConnSlice[i].Close()
			utils.ConnSlice = append(utils.ConnSlice[:i], utils.ConnSlice[:i+1]...)
		}
	}
}
