package sub

import (
	"github.com/SergeyDavidenko/subscription/models"
)

// NewSubscriptions ...
func NewSubscriptions(name string, price float64) models.Subscription {
	return models.Subscription{
		ID:    genUUID(),
		Name:  name,
		Price: price,
	}
}
