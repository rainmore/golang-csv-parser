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

	var artists = make(map[string]Artist)

	fmt.Printf("size: %v", len(records))

	for _, record := range records {
		var song Song
		song.title = record[0]
		song.artist = record[1]
		song.album = record[2]
		song.yearReleased = record[3]
		song.fileSize = record[4]
		song.songDuration = record[5]
		song.filename = record[6]
		song.genre = record[7]
		song.playCount = record[8]
		song.rating = record[9]
		song.addedToLibraryOnTimestamp = record[10]
		song.addedToLibraryOnEpoch = record[11]
		song.composer = record[12]
		song.comment = record[13]

		artist, ok := artists[song.artist]
		if ok {
			artist.albums[song.album] = append(artist.albums[song.album], song)
			artists[song.artist] = artist
		} else {
			var artistNew Artist
			artistNew.name = song.artist
			artistNew.albums = make(map[string][]Song)
			artistNew.albums[song.album] = []Song{song}
			artists[song.artist] = artistNew
		}
	}

	fmt.Printf("artist size: %v", len(artists))
}
