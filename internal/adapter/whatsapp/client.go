package whatsapp

import (
	"context"

	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/store/sqlstore"
	waLog "go.mau.fi/whatsmeow/util/log"
	_ "modernc.org/sqlite"
)

type Client struct {
	Client *whatsmeow.Client
}

func NewClient(ctx context.Context, dbUrl string) (*Client, error) {
	dbLog := waLog.Stdout("Database", "DEBUG", true)
	container, err := sqlstore.New(
		ctx,
		"sqlite",
		dbUrl,
		dbLog,
	)

	if err != nil {
		return nil, err
	}

	device, err := container.GetFirstDevice(ctx)
	if err != nil {
		return nil, err
	}

	clientLog := waLog.Stdout("WA", "DEBUG", true)
	client := whatsmeow.NewClient(device, clientLog)

	return &Client{
		Client: client,
	}, nil
}
