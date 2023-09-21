package main

import (
	config "apiGo/Config"
	"apiGo/router"
)

var (
	logger *config.Logger
)

func main() {

	logger = config.GetLogger("main")
	// Initialize Configs
	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}
	// Initialize router
	router.Initialize()
}
