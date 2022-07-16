package api

import (
	log "github.com/sirupsen/logrus"
)

func Run(done chan<- bool, messages chan<- string) {
	router := InitRouter(messages)

	// TODO: config
	if err := router.Run("0.0.0.0:8080"); err != nil {
		log.Error("api run error", err)
		done <- true
	}
}
