package discordconnect

import (
	"discordMsgRead/internal/utils"
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
)

var (
	msgStruct           utils.Message
	voiceStateMsgStruct utils.VoiceStateMessage
)

func RecieveMsg(conn **websocket.Conn, flags utils.Config) {
	for {
		msgType, msg, readErr := (*conn).ReadMessage()
		if msgType == websocket.CloseAbnormalClosure || msgType == websocket.CloseMessage || msgType == websocket.CloseNormalClosure || msgType == -1 {
			(*conn).Close()
			log.Println("Streamkit disconnected")
			return
		}
		if readErr != nil {
			(*conn).Close()
			log.Printf("Read websocket streamkit message error: %v.", readErr)
			return
		}
		json.Unmarshal(msg, &msgStruct)
		if msgStruct.Evt == "READY" {
			(*conn).WriteMessage(websocket.TextMessage, []byte(`{"cmd":"AUTHENTICATE","args":{"access_token":"`+flags.AccessToken+`"},"nonce":"d5b2d0a1-2898-4723-9aad-63545a4f731c"}`))
		} else if msgStruct.Cmd == "AUTHENTICATE" {
			(*conn).WriteMessage(websocket.TextMessage, []byte(`{"cmd":"GET_CHANNEL","args":{"channel_id":"`+flags.ChannelID+`"},"nonce":"cf52e192-3f93-4165-ab08-a5f60c8e3847"}`))
			(*conn).WriteMessage(websocket.TextMessage, []byte(`{"cmd":"SUBSCRIBE","args":{"channel_id":"`+flags.ChannelID+`"},"evt":"VOICE_STATE_CREATE","nonce":"ae965ff7-845c-40f7-9ab0-90194e04f7af"}`))
			(*conn).WriteMessage(websocket.TextMessage, []byte(`{"cmd":"SUBSCRIBE","args":{"channel_id":"`+flags.ChannelID+`"},"evt":"VOICE_STATE_DELETE","nonce":"7457b728-e7d3-4e0c-964d-3fa2bbe84315"}`))
			(*conn).WriteMessage(websocket.TextMessage, []byte(`{"cmd":"SUBSCRIBE","args":{"channel_id":"`+flags.ChannelID+`"},"evt":"VOICE_STATE_UPDATE","nonce":"3a0ddacf-51d3-4382-8db0-5ca6073c6345"}`))
			(*conn).WriteMessage(websocket.TextMessage, []byte(`{"cmd":"SUBSCRIBE","args":{"channel_id":"`+flags.ChannelID+`"},"evt":"SPEAKING_START","nonce":"efe7f0ec-f640-4367-b672-87b290c362e8"}`))
			(*conn).WriteMessage(websocket.TextMessage, []byte(`{"cmd":"SUBSCRIBE","args":{"channel_id":"`+flags.ChannelID+`"},"evt":"SPEAKING_STOP","nonce":"d2b24e55-78c8-438a-a180-4bba592ee02d"}`))
		} else if msgStruct.Cmd == "GET_CHANNEL" {
			utils.GetChannelMessage = msg
		} else {
			if msgStruct.Evt == "VOICE_STATE_CREATE" || msgStruct.Evt == "VOICE_STATE_UPDATE" {
				json.Unmarshal(msg, &voiceStateMsgStruct)
				utils.UserStatuses[voiceStateMsgStruct.UserData.Nick] = msg
			} else if msgStruct.Evt == "VOICE_STATE_DELETE" {
				json.Unmarshal(msg, &voiceStateMsgStruct)
				delete(utils.UserStatuses, voiceStateMsgStruct.UserData.Nick)
			}
			go Broadcast(msg, utils.ConnSlice)
		}
	}
}
