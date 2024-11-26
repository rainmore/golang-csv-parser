package main

import (
	"os"
	"rainmore/csv-parser/parser"
)

func main() {
	file, err := os.Open("resources/music.csv")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	parser.Parser(file)

}
