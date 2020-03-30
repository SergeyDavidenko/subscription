package api

import (
	"net/http"

	"github.com/SergeyDavidenko/subscription/models"
	"github.com/gin-gonic/gin"
)

func updateSubscriptions(c *gin.Context) {
	var subscriptions models.Subscription
	err := c.BindJSON(&subscriptions)
	if err != nil {
		respondWithError(http.StatusBadRequest, "bad reqeust", c)
		return
	}

	errCreate := subscriptions.UpdateSubscriptionsOnDB()
	if errCreate != nil {
		respondWithError(http.StatusInternalServerError, errCreate.Error(), c)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "update success", "subscription": subscriptions})
}