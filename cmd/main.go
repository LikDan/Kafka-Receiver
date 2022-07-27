package main

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"receiver/internal/controller"
)

func main() {
	conf := kafka.ReaderConfig{
		Brokers:  []string{"localhost:9092"},
		Topic:    "user-messages",
		MaxBytes: 10,
	}

	fmt.Println(conf)

	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", "answer-messages", 0)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	ctrl := controller.NewController(conn)
	fmt.Println(ctrl)
}
