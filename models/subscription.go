package models

import (
	"context"
	"fmt"

	"github.com/SergeyDavidenko/subscription/db"
	log "github.com/sirupsen/logrus"
)

// Subscription ...
type Subscription struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

// CreateSubscriptionsOnDB ...
func (s *Subscription) CreateSubscriptionsOnDB() error {
	sql := `
	INSERT INTO subscriptions(id, name, price) VALUES ($1, $2, $3)
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
		s.Name,
		s.Price,
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

// DeleteSubscriptionsOnDB ...
func (s *Subscription) DeleteSubscriptionsOnDB() error {
	sql := `
	DELETE FROM subscriptions WHERE id = $1
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

// UpdateSubscriptionsOnDB ...
func (s *Subscription) UpdateSubscriptionsOnDB() error {
	sql := `
	UPDATE subscriptions SET price=$1, name=$2 WHERE id=$3;
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
		s.Price, s.Name, s.ID,
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

//GetSubscriptionsOnDBByID ...
func (s *Subscription) GetSubscriptionsOnDBByID() (Subscription, error) {
	sql := `
	SELECT * FROM subscriptions WHERE id = $1;
	`
	var sub Subscription
	errRow := db.DB.QueryRow(context.Background(), sql, s.ID).Scan(&sub.ID,
		&sub.Name, &sub.Price)
	if errRow != nil {
		return sub, errRow
	}
	log.Debug(sub)
	return sub, nil
}

//GetSubscriptionsOnDB ...
func (s *Subscription) GetSubscriptionsOnDB() ([]Subscription, error) {
	sql := `
	SELECT * FROM subscriptions LIMIT 1000;
	`
	var subs []Subscription
	rows, errRow := db.DB.Query(context.Background(), sql)
	if errRow != nil {
		log.Error(errRow)
		return subs, errRow
	}
	for rows.Next() {
		var sub Subscription
		err := rows.Scan(&sub.ID, &sub.Price, &sub.Name)
		if err != nil {
			log.Error(err)
			return subs, err
		}
		subs = append(subs, sub)
	}
	log.Debug(subs)
	return subs, nil
}

// ValidateSubscriptions ...
func (s *Subscription) ValidateSubscriptions() error {
	if s.ID == "" {
		return fmt.Errorf("ID not set")
	}
	if s.Name == "" {
		return fmt.Errorf("Name not set")
	}
	if len(s.Name) < 3 {
		return fmt.Errorf("Len name < 3")
	}
	if s.Price == 0 {
		return fmt.Errorf("Price not set")
	}
	return nil
}
