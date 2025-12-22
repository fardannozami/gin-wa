package usecase

import (
	"log"

	"github.com/fardannozami/gin-wa/internal/domain"
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
