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

	var targetFolder = "resources/target"
	os.RemoveAll(targetFolder)

	parser.Parser(file, "resources/", targetFolder)

	// parser.ConvertRestMp3("resources/iPod_Control/Music/")
}
