package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/fardannozami/gin-wa/internal/adapter/whatsapp"
	"github.com/fardannozami/gin-wa/internal/usecase"
	"go.mau.fi/whatsmeow"
)

func main() {
	ctx := context.Background()

	// usecase
	messageUC := usecase.NewMessageUseCase()

	// handler
	eventHandler := whatsapp.NewEventHandler(messageUC)

	whatsappClient, err := whatsapp.NewClient(ctx, "file:examplestore.db?_pragma=foreign_keys(1)&_pragma=busy_timeout(5000)&_pragma=journal_mode(WAL)")

	if err != nil {
		log.Fatal(err)
	}

	client := whatsappClient.Client
	client.AddEventHandler(eventHandler.HandleEvent)
	// 1️⃣ CONNECT DULU (WAJIB)
	err = client.Connect()
	if err != nil {
		panic(err)
	}

	// 2️⃣ CEK LOGIN
	if client.Store.ID == nil {
		log.Println("📱 Belum login, generate pairing code...")

		pairingCode, err := client.PairPhone(ctx, "628985066454", true, whatsmeow.PairClientChrome, "Chrome (Linux)")
		if err != nil {
			log.Println("⚠️ PairPhone warning:", err)
		}

		fmt.Println("🔢 Pairing Code:", pairingCode)
		fmt.Println("➡️ WhatsApp > Linked Devices > Link with phone number")
	}

	// Wait for shutdown signal to exit cleanly instead of Ctrl+C status code.
	shutdownCtx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	<-shutdownCtx.Done()
	log.Println("🛑 Shutdown signal received, disconnecting...")
	client.Disconnect()

}
