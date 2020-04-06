package api

import (
	"net/http"

	"github.com/SergeyDavidenko/subscription/models"
	"github.com/gin-gonic/gin"
)

func createUserSubscription(c *gin.Context) {
	var subscriptions models.SubscriptionUser
	err := c.BindJSON(&subscriptions)
	if err != nil {
		respondWithError(http.StatusBadRequest, "bad reqeust", c)
		return
	}
	errCreate := subscriptions.CreateSubscriptionUserOnDB()
	if errCreate != nil {
		respondWithError(http.StatusBadRequest, "bad reqeust", c)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "create success", "subscriptions": subscriptions})
}
