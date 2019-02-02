package handlers

import (
	"net/http"

	"github.com/bushimen/social/api/models"
	"github.com/gin-gonic/gin"
)

// GetInstagram get a single instagram post
func GetInstagram(c *gin.Context) {
	shortcode := c.Param("shortcode")
	ins, err := models.GetInstagram(shortcode)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err})
	}

	exists, err := ins.Exists()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err})
		return
	}

	if !exists {
		c.JSON(http.StatusOK, gin.H{"data": nil})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": ins})
}

// GetInstagrams get multiple instagram posts
func GetInstagrams(c *gin.Context) {
	var query models.FetchQuery
	c.BindQuery(&query)

	ins, err := models.GetInstagrams(query)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": ins})
}

// PutInstagram upserts an instagram post
func PutInstagram(c *gin.Context) {
	var ins models.Instagram
	if err := c.BindJSON(&ins); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	err := models.UpsertInstagram(&ins)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": ins})
}
