package whatsapp

import (
	"log"

	"github.com/fardannozami/gin-wa/internal/usecase"
	"go.mau.fi/whatsmeow/types/events"
)

type EventHandler struct {
	MessageUC *usecase.MessageUseCase
}

func NewEventHandler(messageUC *usecase.MessageUseCase) *EventHandler {
	return &EventHandler{
		MessageUC: messageUC,
	}
}

func (h *EventHandler) HandleEvent(evt interface{}) {
	switch evt.(type) {

	case *events.PairSuccess:
		log.Println("🎉 Pairing SUCCESS")

	case *events.Connected:
		log.Println("🔗 Connected")

	case *events.Disconnected:
		log.Println("❌ Disconnected")
	}
}
