package main

import (
	"context"

	"github.com/FabioSebs/whatsappAutomation/template"
	waclient "github.com/FabioSebs/whatsappAutomation/wa_client"
	"go.mau.fi/whatsmeow/proto/waE2E"
	"go.mau.fi/whatsmeow/types"
)

func main() {
	//TODO: get recipient information
	//TODO: send data in a loop
	return
}

func SendMessage(phone string, name string) (err error) {
	var (
		msg    = template.JEANLY_WA_MSG
		ctx    = context.Background()
		client = waclient.NewWhatsappClient()
	)
	_, err = client.SendMessage(ctx, types.JID{
		User:   phone,
		Server: "s.whatsapp.net",
	}, &waE2E.Message{
		Conversation: &msg,
	})
	return
}
