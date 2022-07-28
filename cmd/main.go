package main

import (
	"fmt"
	"github.com/segmentio/kafka-go"
	"receiver/internal/controller"
	"receiver/internal/repository"
)

func main() {
	w := &kafka.Writer{
		Addr:     kafka.TCP("localhost:9092"),
		Topic:    "answer-messages",
		Balancer: &kafka.LeastBytes{},
	}

	repo := repository.NewRepository(w)
	ctrl := controller.NewController(repo)
	fmt.Println(ctrl)
}
