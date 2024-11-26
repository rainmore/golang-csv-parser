package parser

import (
	"encoding/csv"
	"fmt"
	"io"
)

func Parser(file io.Reader) {
	reader := csv.NewReader(file)

	records, err := reader.ReadAll()

	if err != nil {
		panic(err)
	}

	for _, record := range records {
		fmt.Println(record)
	}
}
