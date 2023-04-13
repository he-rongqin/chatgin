package main

import (
	"org.chatgin/src/config"
	"org.chatgin/src/router"
)

func main() {
	r := router.Router()
	r.Run(config.ServerConfig.ServerPort)
}
