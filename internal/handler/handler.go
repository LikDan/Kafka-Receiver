package handler

import (
	"context"
	"log"
	"receiver/internal/controller"
	"receiver/internal/entities"
)

type Handler struct {
	controller controller.IController
}

func NewHandler(controller controller.IController) *Handler {
	return &Handler{controller: controller}
}

func (h *Handler) receive(c chan entities.Message) {
	for {
		message, err := h.controller.Read(context.Background())
		if err != nil {
			log.Println("Read error: ", err)
			continue
		}

		c <- message
	}
}

func (h *Handler) Proceed() {
	messageChan := make(chan entities.Message)

	go func() {
		for {
			message := <-messageChan
			log.Println("received message: ", message)

			ctx := context.Background()
			if err := h.controller.Answer(ctx, message, "Answer"); err != nil {
				log.Println("Answer error: ", err)
				continue
			}
		}
	}()

	h.receive(messageChan)
}
