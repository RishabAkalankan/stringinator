package main

import (
	"github.com/RishabAkalankan/stringinator/api"
	"github.com/RishabAkalankan/stringinator/logger"
)

func main() {
	logger.Initialize()
	api.StartServer()
}
