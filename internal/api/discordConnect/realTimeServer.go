package discordconnect

import (
	"discordMsgRead/internal/utils"

	"github.com/gorilla/websocket"
)

func Broadcast(messege []byte) {
	for i := 0; i < len(utils.ConnSlice); i++ {
		err := utils.ConnSlice[i].WriteMessage(websocket.TextMessage, messege)
		if err != nil {
			go utils.ConnSlice[i].Close()
			utils.ConnSlice = append(utils.ConnSlice[:i], utils.ConnSlice[i+1:]...)
		}
	}
}
