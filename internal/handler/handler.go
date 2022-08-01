package handler

import (
	"context"
	"fmt"
	"receiver/internal/controller"
	"receiver/internal/entities"
)

type Handler struct {
	controller controller.IController
}

func NewHandler(controller controller.IController) *Handler {
	handler := Handler{controller: controller}
	go handler.proceed()

	return &handler
}

func (h *Handler) receive(c chan entities.Message) {
	for {
		message, err := h.controller.Read(context.Background())
		if err != nil {
			fmt.Println("Read error: ", err)
			continue
		}

		c <- message
	}
}

func (h *Handler) proceed() {
	messageChan := make(chan entities.Message)

	go func() {
		for {
			message := <-messageChan

			fmt.Println("received message: ", message)

			if err := h.controller.Answer(context.Background(), message, "Answer"); err != nil {
				fmt.Println("Answer error: ", err)
				continue
			}
		}
	}()

	h.receive(messageChan)
}
