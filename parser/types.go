package parser

type Song struct {
	title                     string
	artist                    string
	album                     string
	yearReleased              string
	fileSize                  string
	songDuration              string
	filename                  string
	genre                     string
	playCount                 string
	rating                    string
	addedToLibraryOnTimestamp string
	addedToLibraryOnEpoch     string
	composer                  string
	comment                   string
}

type Artist struct {
	name   string
	albums map[string][]Song
}
