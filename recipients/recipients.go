package recipients

import (
	"fmt"
	"log"

	"github.com/xuri/excelize/v2"
)

func GetRecipients() People {
	// opening
	f, err := excelize.OpenFile("data/official.xlsx")
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
	rows, err := f.GetRows("Data")
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
			Phone:       row[8], //wa.me/628111357100
			Name:        row[3],
			Institution: row[5],
			ContactPerson: ContactPerson{
				Title: row[9],
				Name:  row[10],
				Phone: row[13],
			},
			Familiar: row[14],
		}
		people = append(people, person)
	}

	return people
}

type Person struct {
	Phone         string
	Name          string
	Institution   string
	ContactPerson ContactPerson
	Familiar      string
}

type ContactPerson struct {
	Title string
	Name  string
	Phone string
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
		},
	}
)
