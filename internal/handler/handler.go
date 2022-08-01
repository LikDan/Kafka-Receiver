package handler

import (
	"context"
	"fmt"
	"receiver/internal/controller"
)

type Handler struct {
	controller controller.IController
}

func NewHandler(controller controller.IController) *Handler {
	handler := Handler{controller: controller}
	go handler.receive()

	return &handler
}

func (h *Handler) receive() {
	for {
		message, err := h.controller.Read(context.Background())
		if err != nil {
			fmt.Println("Read error: ", err)
			continue
		}

		fmt.Println("received message: ", message)

		if err = h.controller.Answer(context.Background(), message, "Answer"); err != nil {
			fmt.Println("Answer error: ", err)
			continue
		}
	}
}
