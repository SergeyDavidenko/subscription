package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Info check func
func info(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "running",
	})
}
