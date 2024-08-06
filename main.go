package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/FabioSebs/whatsappAutomation/recipients"
	waclient "github.com/FabioSebs/whatsappAutomation/wa_client"
	"go.mau.fi/whatsmeow/proto/waE2E"
	"go.mau.fi/whatsmeow/types"
)

func main() {
	//TODO: get recipient information
	people := recipients.MOCK_PPL
	if err := SendMessage(people); err != nil {
		panic(err)
	}
}

func SendMessage(people recipients.People) (err error) {
	var (
		ctx = context.Background()
	)

	client := waclient.NewWhatsappClient()

	for _, person := range people {
		msg := fmt.Sprintf(recipients.MOCK_MSG, person.Title, person.Name)
		if _, err = client.SendMessage(ctx, types.JID{
			User:   strings.Trim(person.Phone, " "),
			Server: "s.whatsapp.net",
		}, &waE2E.Message{
			Conversation: &msg,
		}); err != nil {
			return err
		}
	}

	return
}
