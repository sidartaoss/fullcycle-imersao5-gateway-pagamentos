package kafka

import (
	"log"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

type Consumer struct {
	*ckafka.ConfigMap
	Topics []string
}

func NewConsumer(configMap *ckafka.ConfigMap, topics []string) *Consumer {
	return &Consumer{
		ConfigMap: configMap,
		Topics:    topics,
	}
}

func (c *Consumer) Consume(msgChan chan *ckafka.Message) error {
	consumer, err := ckafka.NewConsumer(c.ConfigMap)
	if err != nil {
		log.Println(err)
		return err
	}
	err = consumer.SubscribeTopics(c.Topics, nil)
	if err != nil {
		log.Println(err)
		return err
	}
	for {
		msg, err := consumer.ReadMessage(-1)
		if err != nil {
			log.Println(err)
			return err
		}
		log.Println("reading message")
		msgChan <- msg
	}
}
