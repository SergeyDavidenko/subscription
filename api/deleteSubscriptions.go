package api

import (
	"net/http"

	"github.com/SergeyDavidenko/subscription/models"
	"github.com/gin-gonic/gin"
)

func deleteSubscriptions(c *gin.Context) {
	var subscriptions models.Subscription
	err := c.BindJSON(&subscriptions)
	if err != nil {
		respondWithError(http.StatusBadRequest, "bad reqeust", c)
		return
	}

	errCreate := subscriptions.DeleteSubscriptionsOnDB()
	if errCreate != nil {
		respondWithError(http.StatusInternalServerError, errCreate.Error(), c)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "delete success", "id": subscriptions.ID})
}
