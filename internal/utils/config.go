package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	UserID      string `json:"userID"`
	AccessToken string `json:"accessToken"`
	ChannelID   string `json:"channelID"`
}

func ConfigParce() (config Config) {
	file, openErr := os.Open("./config.json")
	if openErr != nil {
		log.Fatalf("Open config error: %v", openErr)
	}
	data, readErr := ioutil.ReadAll(file)
	if readErr != nil {
		log.Fatalf("Read file error: %v", readErr)
	}
	unmarshErr := json.Unmarshal(data, &config)
	if unmarshErr != nil {
		log.Fatalf("JSON unmarshal error: %v", unmarshErr)
	}
	return
}
