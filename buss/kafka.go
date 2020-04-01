package buss

import (
	"encoding/json"

	"github.com/SergeyDavidenko/subscription/config"
	"github.com/SergeyDavidenko/subscription/models"

	log "github.com/sirupsen/logrus"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

// ConsumerKafka ... func
var ConsumerKafka *kafka.Consumer

// KafkaConnect ...
func KafkaConnect() *kafka.Consumer {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": config.Conf.Kafka.Address,
		"group.id":          "subscription",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		log.Fatal(err)
	}
	ConsumerKafka = c
	return c
}

// RunKafkaConsumer ...
func RunKafkaConsumer(consumer *kafka.Consumer) {
	consumer.SubscribeTopics([]string{config.Conf.Kafka.Topic}, nil)
	log.Info("Start kafka consumer on topic ", config.Conf.Kafka.Topic)
	defer consumer.Close()
	var sub models.SubscriptionUser
	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			//log.Info(fmt.Sprintf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value)))
			if errJSON := json.Unmarshal(msg.Value, &sub); errJSON != nil {
				log.Error(errJSON)
			}
			errCreateSub := sub.CreateSubscriptionUserOnDB()
			if errCreateSub != nil {
				log.Error(errCreateSub)
			}
			log.Info(sub)
		} else {
			log.Error(err)
		}
	}
}
