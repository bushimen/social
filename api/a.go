package main

import (
	"net/http"

	"github.com/bushimen/social/api/handlers"
	"github.com/bushimen/social/api/models"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	models.Connect()
	defer models.Disconnect()

	// Disable Console Color
	gin.DisableConsoleColor()

	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	router.GET("/instagram/:shortcode", handlers.GetInstagram)
	router.GET("/instagram", handlers.GetInstagrams)
	router.PUT("/instagram", handlers.PutInstagram)

	router.GET("/bilibili/:aid", handlers.GetBilibili)
	router.GET("/bilibili", handlers.GetBilibilis)
	router.PUT("/bilibili", handlers.PutBilibili)

	router.Run()
}
