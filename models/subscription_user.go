package models

// Login ...
type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// SubscriptionUser ...
type SubscriptionUser struct {
	ID               string `json:"id"`
	UserID           string `json:"user_id"`
	StartSubcruption int64  `json:"start_subscription"`
	ExpiresAt        int64  `json:"expipe_subscription"`
	Activate         bool   `json:"activate"`
}
