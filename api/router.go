package api

import "github.com/gin-gonic/gin"

func InitRouter(messages chan<- string) *gin.Engine {
	router := gin.Default()

	router.POST("/", CollectHandlerInit(messages))

	return router
}
