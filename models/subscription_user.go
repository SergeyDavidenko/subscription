package models

import (
	"context"

	"github.com/SergeyDavidenko/subscription/db"
	log "github.com/sirupsen/logrus"
)

// Login ...
type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// KafkaSubscriptionUser ...
type KafkaSubscriptionUser struct {
	Action           string           `json:"action"`
	UserSubscription SubscriptionUser `json:"subscription"`
}

// SubscriptionUser ...
type SubscriptionUser struct {
	ID               string `json:"id"`
	UserID           string `json:"user_id"`
	StartSubcruption int64  `json:"start_subscription"`
	ExpiresAt        int64  `json:"expipe_subscription"`
	Activate         bool   `json:"activate"`
}

// CreateSubscriptionUserOnDB ...
func (s *SubscriptionUser) CreateSubscriptionUserOnDB() error {
	sql := `
	INSERT INTO subscriptions_user(id, 
		user_id, 
		start_subscription, 
		expipe_subscription, 
		activate) VALUES ($1, $2, $3, $4, $5)
	`
	tx, err := db.DB.Begin(context.Background())
	if err != nil {
		log.Error(err)
		return err
	}
	defer tx.Rollback(context.Background())
	_, err = tx.Exec(
		context.Background(), sql,
		s.ID,
		s.UserID,
		s.StartSubcruption,
		s.ExpiresAt,
		s.Activate,
	)
	if err != nil {
		log.Error(err)
		return err
	}
	err = tx.Commit(context.Background())
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

// UpdateSubscriptionUserOnDB ...
func (s *SubscriptionUser) UpdateSubscriptionUserOnDB() error {
	sql := `
	UPDATE subscriptions_user SET activate=$1 WHERE user_id=$2;
	`
	tx, err := db.DB.Begin(context.Background())
	if err != nil {
		log.Error(err)
		return err
	}
	defer tx.Rollback(context.Background())
	log.Debug(s)
	_, err = tx.Exec(
		context.Background(), sql,
		s.Activate,
		s.UserID,
	)
	if err != nil {
		log.Error(err)
		return err
	}
	err = tx.Commit(context.Background())
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}
