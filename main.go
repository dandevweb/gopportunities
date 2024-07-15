package main

import (
	"github.com/dandevweb/gopportunities/config"
	"github.com/dandevweb/gopportunities/router"
)

var (
	logger *config.Logger
)

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func main() {
	logger = config.GetLogger("main")
	err := config.Init()
	if err != nil {
		logger.Errorf("Error initializing config: %v", err)
		return
	}

	router.Initialize()
}
