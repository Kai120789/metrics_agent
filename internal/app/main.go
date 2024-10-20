package app

import (
	"agent/internal/config"
	"fmt"
)

func StartApp() {
	// init config
	cfg, err := config.GetConfig()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(cfg)

	// init logger

	// init service

	// init handler

	// init router

	//start server
}
