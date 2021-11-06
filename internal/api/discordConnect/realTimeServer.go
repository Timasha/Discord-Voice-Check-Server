package discordconnect

import (
	"discordMsgRead/internal/utils"

	"github.com/gorilla/websocket"
)

func close(conn *websocket.Conn) {
	conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(1000, "woops"))
	conn.Close()
}
func Broadcast(messege []byte) {
	for i := 0; i < len(utils.ConnSlice); i++ {
		err := utils.ConnSlice[i].WriteMessage(websocket.TextMessage, messege)
		if err != nil {
			utils.BroadcastGroup.Wait()
			utils.DeleteGroup.Add(1)
			close(utils.ConnSlice[i])
			utils.ConnSlice = append(utils.ConnSlice[:i], utils.ConnSlice[:i+1]...)
			utils.DeleteGroup.Done()
			i--
		}
	}
}
