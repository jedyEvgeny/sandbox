package main

import (
	"context"
	"fmt"
	"log"
	"time"

	kafka "github.com/segmentio/kafka-go"
)

const (
	topic              = "my-learning-topic"
	kafkaBrokerAddress = "localhost:9093"
)

func main() {
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{kafkaBrokerAddress},
		Topic:   topic,
	})
	defer w.Close()

	msgs := []string{
		"Привет, го-разработчики!",
		"Это первые сообщения через Кафку",
		"По сути - это общение между двумя микросервисами",
	}

	sendMsgToKafka(w, msgs)

	fmt.Println("Producer finished.")
}

func sendMsgToKafka(w *kafka.Writer, msgs []string) {
	for i, msgBody := range msgs {
		msg := kafka.Message{
			Key:   []byte(fmt.Sprintf("Сообщение №%d", i+1)),
			Value: []byte(msgBody),
			Time:  time.Now(),
		}

		err := w.WriteMessages(context.Background(), msg)
		if err != nil {
			log.Printf("ошибка отправки сообщения в Кафку '%s': %v\n", msgBody, err)
		} else {
			fmt.Printf("Отправленное в Кафку сообщение: %s\n", msgBody)
		}
		time.Sleep(200 * time.Millisecond)
	}
}
