package service

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type GinInterface interface {
	Start()
}

func NewRestAPI(port uint16) {
	router := gin.Default()
	router.GET("/version", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"version": "1.0.0",
		})
	})

	host := fmt.Sprintf("localhost:%d", port)

	router.Run(host)
}
