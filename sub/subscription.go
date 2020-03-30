package sub

import (
	"github.com/SergeyDavidenko/subscription/models"
	guuid "github.com/google/uuid"
)

// NewSubscriptions ...
func NewSubscriptions(name string, price float64) models.Subscription {
	return models.Subscription{
		ID:    genUUID(),
		Name:  name,
		Price: price,
	}
}

func genUUID() string {
	id := guuid.New()
	return id.String()
}
