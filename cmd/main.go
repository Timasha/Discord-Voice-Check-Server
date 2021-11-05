package main

import (
	"discordMsgRead/internal/api"
	"discordMsgRead/internal/utils"
)

func main() {
	config := utils.ConfigParce()
	api.Middleware(config)
}
