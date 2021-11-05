package discordconnect

import (
	"discordMsgRead/internal/utils"

	"github.com/gorilla/websocket"
)

func close(i int) {
	utils.ConnSlice[i].WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(1000, "woops"))
	utils.ConnSlice[i].Close()
}
func Broadcast(messege []byte) {
	for i := 0; i < len(utils.ConnSlice); i++ {
		err := utils.ConnSlice[i].WriteMessage(websocket.TextMessage, messege)
		if err != nil {
			go close(i)
			utils.ConnSlice = append(utils.ConnSlice[:i], utils.ConnSlice[i+1:]...)
		}
	}
}
