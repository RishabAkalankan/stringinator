package main

import (
	"github.com/RishabAkalankan/stringinator/logger"
	"github.com/RishabAkalankan/stringinator/server"
)

func main() {
	logger.Initialize()
	server.Start()
}
