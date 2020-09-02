package csvformat

import (
	"encoding/csv"
	"log"
	"os"
)

type Book struct {
	Author string
	Title  string
}

func main() {
	books := []Book{
		{
			"F Scott Fitzgerald",
			"The Great Gatsby",
		},
		{
			"J D Salinger",
			"The Catcher in the Rye",
		},
	}
	f, _ := os.Create("xx.csv")
	defer f.Close()
	writer := csv.NewWriter(f)
	writer.Write([]string{"Author", "Title"})

	for _, b := range books {
		err := writer.Write([]string{b.Author, b.Title})
		if err != nil {
			log.Fatal(err)
		}
	}
	writer.Flush()
}
