package main

import (
	"github.com/segmentio/kafka-go"
	"receiver/internal/controller"
	"receiver/internal/handler"
	"receiver/internal/provider"
)

func main() {
	w := &kafka.Writer{
		Addr:     kafka.TCP("localhost:9092"),
		Topic:    "answer-messages",
		Balancer: &kafka.LeastBytes{},
	}

	conf := kafka.ReaderConfig{
		Brokers:  []string{"localhost:9092"},
		Topic:    "user-messages",
		MaxBytes: 10,
	}

	r := kafka.NewReader(conf)

	repo := provider.NewProvider(w, r)
	ctrl := controller.NewController(repo)
	h := handler.NewHandler(ctrl)
	h.Proceed()
}
