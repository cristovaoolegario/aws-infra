package kafka

import (
	"fmt"
	"log"
	"os"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaConsumer struct {
	MessageChan chan *ckafka.Message
}

func NewKafkaConsumer(msgChan chan *ckafka.Message) *KafkaConsumer {
	return &KafkaConsumer{
		MessageChan: msgChan,
	}
}

func (kc *KafkaConsumer) Consume() {
	configMap := &ckafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KafkaBootstrapServers"),
		"group.id":          os.Getenv("KafkaConsumerGroupId"),
	}
	c, err := ckafka.NewConsumer(configMap)
	if err != nil {
		log.Fatalf("error consuming kafka messages:" + err.Error())
	}

	topics := []string{os.Getenv("KafkaReadTopic")}

	c.SubscribeTopics(topics, nil)

	fmt.Println("Kafka consumer has been started")

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			kc.MessageChan <- msg
		}
	}
}
