package api

import (
	"net/http"

	"github.com/SergeyDavidenko/subscription/models"
	"github.com/SergeyDavidenko/subscription/sub"
	"github.com/SergeyDavidenko/subscription/utils"
	"github.com/gin-gonic/gin"
)

func createSubscriptions(c *gin.Context) {
	var subscriptions models.Subscription
	err := c.BindJSON(&subscriptions)
	if err != nil {
		respondWithError(http.StatusBadRequest, "bad reqeust", c)
		return
	}
	sub := sub.NewSubscriptions(subscriptions.Name, subscriptions.Price)
	errValid := sub.ValidateSubscriptions()
	if errValid != nil {
		respondWithError(http.StatusBadRequest, errValid.Error(), c)
		return
	}
	errCreate := sub.CreateSubscriptionsOnDB()
	if errCreate != nil {
		if utils.ErrorCodePG(errCreate) == "23505" {
			respondWithError(http.StatusBadRequest, "subscription already exist", c)
			return
		}
		respondWithError(http.StatusInternalServerError, errCreate.Error(), c)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "create success", "subscriptions": sub})
}
