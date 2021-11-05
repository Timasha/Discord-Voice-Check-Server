package utils

import "github.com/gorilla/websocket"

var ConnSlice []*websocket.Conn = make([]*websocket.Conn, 0)
var GetChannelMessage []byte
var UserStatuses map[string][]byte = make(map[string][]byte, 0)

type Message struct {
	Cmd      string   `json:"cmd"`
	UserData struct{} `json:"data"`
	Evt      string   `json:"evt"`
	Nonce    string   `json:"nonce"`
}
type VoiceStateMessage struct {
	Cmd      string `json:"cmd"`
	UserData Data   `json:"data"`
	Evt      string `json:"evt"`
	Nonce    string `json:"nonce"`
}
type Data struct {
	Nick        string   `json:"nick"`
	Mute        bool     `json:"mute"`
	Volume      int      `json:"volume"`
	Pan         struct{} `json:"pan"`
	Voice_state struct{} `json:"voice_state"`
	User        struct{} `json:"user"`
}
