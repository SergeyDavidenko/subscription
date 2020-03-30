package models

// Claims ...
type Claims struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	Superuser bool   `json:"superuser"`
}
