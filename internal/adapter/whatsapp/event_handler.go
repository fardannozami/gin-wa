package whatsapp

import (
	"github.com/fardannozami/gin-wa/internal/domain"
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
	switch v := evt.(type) {
	case *events.Message:
		h.MessageUC.HandleIncomingMessage(domain.Message{
			From:    v.Info.Sender.User,
			Content: v.Message.GetConversation(),
		})
	}
}
