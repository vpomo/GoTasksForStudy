package main

import (
	"fmt"
	"log"
	"time"

	"github.com/IBM/sarama"
)

func main() {
	// Конфигурация продюсера
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Timeout = 5 * time.Second

	// Создание синхронного продюсера
	producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		log.Fatalf("Error creating producer: %v", err)
	}
	defer producer.Close()

	// Параметры сообщения
	topic := "test-topic"
	message := "Hello Kafka!"

	// Отправка 5 сообщений
	for i := 0; i < 5; i++ {
		msg := &sarama.ProducerMessage{
			Topic: topic,
			Value: sarama.StringEncoder(fmt.Sprintf("%s %d", message, i)),
		}

		// Отправка сообщения
		partition, offset, err := producer.SendMessage(msg)
		if err != nil {
			log.Printf("Failed to send message: %v", err)
			continue
		}

		log.Printf("Sent message %d to partition %d, offset %d", i, partition, offset)
		time.Sleep(1 * time.Second)
	}
}
