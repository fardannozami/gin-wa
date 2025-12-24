package usecase

import (
	"log"

	domain "github.com/fardannozami/gin-wa/internal/domain/message"
)

type MessageUseCase struct{}

func NewMessageUseCase() *MessageUseCase {
	return &MessageUseCase{}
}

func (uc *MessageUseCase) HandleIncomingMessage(message domain.Message) {
	// bisnis logi disni
	log.Println("Received message from", message.From, ":", message.Content)
	//simpan db dll
}
