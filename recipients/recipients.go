package recipients

import (
	"fmt"
	"log"

	"github.com/xuri/excelize/v2"
)

func GetRecipients() People {
	// opening
	f, err := excelize.OpenFile("data/recipients.xlsx")
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}

	// closing
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	// choosing sheet
	rows, err := f.GetRows("Contacts")
	if err != nil {
		log.Fatalf("Failed to get rows: %v", err)
	}

	// traversing
	var people People
	for i, row := range rows {
		if i == 0 {
			continue
		}
		person := Person{
			Phone:       row[0],
			Name:        row[1],
			Title:       row[2],
			Institution: row[3],
		}
		people = append(people, person)
	}

	return people
}

type Person struct {
	Phone       string
	Name        string
	Title       string
	Institution string
}

type People []Person

const (
	MOCK_MSG = `HELLO %s %s FROM WHATSAPP ICCT SCRIPT`
)

var (
	MOCK_PPL People = People{
		Person{
			Phone: "‪6281807408933‬",
			Name:  "Jeanly",
			Title: "Kak",
		},
	}
)
