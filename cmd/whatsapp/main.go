package main

import (
	"context"

	"github.com/fardannozami/gin-wa/internal/adapter/whatsapp"
	"github.com/fardannozami/gin-wa/internal/usecase"
)

func main() {
	ctx := context.Background()

	// usecase
	messageUC := usecase.NewMessageUseCase()

	// handler
	eventHandler := whatsapp.NewEventHandler(messageUC)

	whatsappClient, err := whatsapp.NewClient(ctx, "file:examplestore.db?_pragma=foreign_keys(1)&_pragma=busy_timeout(5000)&_pragma=journal_mode(WAL)")

	if err != nil {
		panic(err)
	}

	whatsappClient.Client.AddEventHandler(eventHandler.HandleEvent)

	// login / connect
	if whatsappClient.Client.Store.ID == nil {

	}
}
