package api

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func CollectHandlerInit(messages chan<- string) func(*gin.Context) {
	return func(c *gin.Context) {
		rawData, err := c.GetRawData()
		if err != nil {
			log.Error("raw data error", err)
			c.JSON(http.StatusBadRequest, gin.H{})
		}

		messages <- string(rawData)
	}
}
