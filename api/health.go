package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Health check func
func health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}


