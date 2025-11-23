package service

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (m *MioEngine) NewRestAPI(port uint16) {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:    []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:   []string{"Content-Length"},
	}))

	router.GET("/version", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"version": "1.0.0",
		})
	})

	router.GET("/config", func(c *gin.Context) {
		c.JSON(200, m.options)
	})

	router.GET("/traffic", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"upstream":   m.GetUpStream(),
			"downstream": m.GetDownStream(),
		})
	})

	router.POST("/config", func(c *gin.Context) {
		var options MioOptions
		if err := c.ShouldBindJSON(&options); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"status": "config updated"})
	})

	host := fmt.Sprintf("localhost:%d", port)

	router.Run(host)
}
