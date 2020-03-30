package api

import (
	"net/http"
	"github.com/SergeyDavidenko/subscription/sub"
	"github.com/gin-gonic/gin"
)

func allSubscriptions(c *gin.Context) {
	sub := sub.NewSubscriptions("", 0)
	subs, err := sub.GetSubscriptionsOnDB()
	if err != nil {
		respondWithError(http.StatusInternalServerError, err.Error(), c)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": subs})
}