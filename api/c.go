package handlers

import (
	"net/http"
	"strconv"

	"github.com/bushimen/social/api/models"
	"github.com/gin-gonic/gin"
)

// GetBilibili get a single bilibili post
func GetBilibili(c *gin.Context) {
	aid, _ := strconv.Atoi(c.Param("aid"))
	b, err := models.GetBilibili(aid)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	exists, err := b.Exists()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if !exists {
		c.JSON(http.StatusOK, gin.H{"data": nil})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": b})
}

// GetBilibilis get multiple bilibili posts
func GetBilibilis(c *gin.Context) {
	var query models.FetchQuery
	c.BindQuery(&query)

	bs, err := models.GetBilibilis(query)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": bs})
}

// PutBilibili upserts an bilibili post
func PutBilibili(c *gin.Context) {
	var b models.Bilibili
	if err := c.BindJSON(&b); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err := models.UpsertBilibili(&b)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": b})
}
