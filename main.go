package main

import (
	"context"
	"fmt"
	"os"
	"strings"
	"unicode"

	"github.com/FabioSebs/whatsappAutomation/recipients"
	"github.com/FabioSebs/whatsappAutomation/template"
	waclient "github.com/FabioSebs/whatsappAutomation/wa_client"
	"github.com/samber/lo"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/proto/waE2E"
	"go.mau.fi/whatsmeow/types"
	"google.golang.org/protobuf/proto"
)

func main() {
	people := recipients.GetRecipients()

	peopleCleaned := lo.Map(people, func(person recipients.Person, idx int) recipients.Person {
		person.Phone = removeCommas(removeControlCharacters(person.Phone))
		person.Name = removeControlCharacters(person.Name)
		person.Title = removeControlCharacters(person.Title)
		return person
	})

	if err := SendMessage(peopleCleaned); err != nil {
		panic(err)
	}
}

func removeControlCharacters(str string) string {
	return strings.Map(func(r rune) rune {
		// Remove control characters like U+202A, U+202B, U+202C, U+202D, U+202E
		if unicode.Is(unicode.Cc, r) || (r >= 0x202A && r <= 0x202E) {
			return -1
		}
		return r
	}, str)
}

func removeCommas(numberStr string) string {
	return strings.ReplaceAll(numberStr, ",", "")
}

func SendMessage(people recipients.People) (err error) {
	var (
		ctx = context.Background()
	)

	client := waclient.NewWhatsappClient()

	fileData, err := os.ReadFile("./icct.png")
	if err != nil {
		return
	}

	resp, err := client.Upload(ctx, fileData, whatsmeow.MediaImage)
	if err != nil {
		return
	}

	for _, person := range people {
		toJID := types.JID{
			User:   person.Phone,
			Server: types.DefaultUserServer,
		}

		message := &waE2E.ImageMessage{
			Caption:       proto.String(fmt.Sprintf(template.JEANLY_WA_MSG, person.Name, person.Institution)),
			URL:           &resp.URL,
			DirectPath:    &resp.DirectPath,
			Mimetype:      proto.String("image/png"),
			FileLength:    &resp.FileLength,
			MediaKey:      resp.MediaKey,
			FileEncSHA256: resp.FileEncSHA256,
			FileSHA256:    resp.FileSHA256,
		}

		// media message
		if _, err = client.SendMessage(ctx, toJID, &waE2E.Message{
			ImageMessage: message,
		}); err != nil {
			return err
		}

		// // chat message
		// chatMsg := fmt.Sprintf(recipients.MOCK_MSG, person.Title, person.Name)
		// if _, err = client.SendMessage(ctx, types.JID{
		// 	User:   strings.Trim(person.Phone, " "),
		// 	Server: "s.whatsapp.net",
		// }, &waE2E.Message{
		// 	Conversation: &chatMsg,
		// }); err != nil {
		// 	return err
		// }
	}

	return
}
