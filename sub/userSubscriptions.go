package sub

import (
	"time"

	"github.com/SergeyDavidenko/subscription/models"
)

// NewUserSubscriptions ...
func NewUserSubscriptions(subID, userID string) models.SubscriptionUser {
	now := time.Now()
	return models.SubscriptionUser{
		ID:               subID,
		UserID:           userID,
		StartSubcruption: now.UTC().Unix(),
		ExpiresAt:        now.AddDate(0, 0, 31).UTC().Unix(),
		Activate:         true,
	}
}
