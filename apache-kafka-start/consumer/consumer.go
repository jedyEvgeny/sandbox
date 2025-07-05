package main

import (
	"context"
	"fmt"
	"log"

	kafka "github.com/segmentio/kafka-go"
)

const (
	topic         = "my-learning-topic"
	brokerAddress = "localhost:9093"
	groupID       = "my-learning-go-group"
)

func main() {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
		GroupID: groupID,
	})
	defer r.Close()

	fmt.Printf("Консьюмер подписан на топик '%s' в группе '%s'\n\n", topic, groupID)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	readMsgs(ctx, r)
}

func readMsgs(ctx context.Context, r *kafka.Reader) {
	for {
		m, err := r.ReadMessage(ctx)
		if err != nil {
			log.Fatalf("ошибка чтения из Kafka: %v\n", err)
		}
		fmt.Printf("Сообщение в топике %v, партиция %v, offset %v: \n\t%s %s\n\n",
			m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
	}
}
