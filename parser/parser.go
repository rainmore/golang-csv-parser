package parser

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const UNKNOWN = "unknown"

var GENRE = map[int]string{
	0:   "Blues",
	1:   "Classic Rock",
	2:   "Country",
	3:   "Dance",
	4:   "Disco",
	5:   "Funk",
	6:   "Grunge",
	7:   "Hip-Hop",
	8:   "Jazz",
	9:   "Metal",
	10:  "New Age",
	11:  "Oldies",
	12:  "Other",
	13:  "Pop",
	14:  "R&B",
	15:  "Rap",
	16:  "Reggae",
	17:  "Rock",
	18:  "Techno",
	19:  "Industrial",
	20:  "Alternative",
	21:  "Ska",
	22:  "Death Metal",
	23:  "Pranks",
	24:  "Soundtrack",
	25:  "Euro-Techno",
	26:  "Ambient",
	27:  "Trip-Hop",
	28:  "Vocal",
	29:  "Jazz+Funk",
	30:  "Fusion",
	31:  "Trance",
	32:  "Classical",
	33:  "Instrumental",
	34:  "Acid",
	35:  "House",
	36:  "Game",
	37:  "Sound Clip",
	38:  "Gospel",
	39:  "Noise",
	40:  "AlternRock",
	41:  "Bass",
	42:  "Soul",
	43:  "Punk",
	44:  "Space",
	45:  "Meditative",
	46:  "Instrumental Pop",
	47:  "Instrumental Rock",
	48:  "Ethnic",
	49:  "Gothic",
	50:  "Darkwave",
	51:  "Techno-Industrial",
	52:  "Electronic",
	53:  "Pop-Folk",
	54:  "Eurodance",
	55:  "Dream",
	56:  "Southern Rock",
	57:  "Comedy",
	58:  "Cult",
	59:  "Gangsta",
	60:  "Top 40",
	61:  "Christian Rap",
	62:  "Pop/Funk",
	63:  "Jungle",
	64:  "Native American",
	65:  "Cabaret",
	66:  "New Wave",
	67:  "Psychadelic",
	68:  "Rave",
	69:  "Showtunes",
	70:  "Trailer",
	71:  "Lo-Fi",
	72:  "Tribal",
	73:  "Acid Punk",
	74:  "Acid Jazz",
	75:  "Polka",
	76:  "Retro",
	77:  "Musical",
	78:  "Rock & Roll",
	79:  "Hard Rock",
	80:  "Folk",
	81:  "Folk-Rock",
	82:  "National Folk",
	83:  "Swing",
	84:  "Fast Fusion",
	85:  "Bebob",
	86:  "Latin",
	87:  "Revival",
	88:  "Celtic",
	89:  "Bluegrass",
	90:  "Avantgarde",
	91:  "Gothic Rock",
	92:  "Progressive Rock",
	93:  "Psychedelic Rock",
	94:  "Symphonic Rock",
	95:  "Slow Rock",
	96:  "Big Band",
	97:  "Chorus",
	98:  "Easy Listening",
	99:  "Acoustic",
	100: "Humour",
	101: "Speech",
	102: "Chanson",
	103: "Opera",
	104: "Chamber Music",
	105: "Sonata",
	106: "Symphony",
	107: "Booty Bass",
	108: "Primus",
	109: "Porn Groove",
	110: "Satire",
	111: "Slow Jam",
	112: "Club",
	113: "Tango",
	114: "Samba",
	115: "Folklore",
	116: "Ballad",
	117: "Power Ballad",
	118: "Rhythmic Soul",
	119: "Freestyle",
	120: "Duet",
	121: "Punk Rock",
	122: "Drum Solo",
	123: "A capella",
	124: "Euro-House",
	125: "Dance Hall",
	126: "Goa",
	127: "Drum & Bass",
	128: "Club-House",
	129: "Hardcore",
	130: "Terror",
	131: "Indie",
	132: "Britpop",
	133: "Negerpunk",
	134: "Polsk Punk",
	135: "Beat",
	136: "Christian Gangsta Rap",
	137: "Heavy Metal",
	138: "Black Metal",
	139: "Crossover",
	140: "Contemporary Christian",
	141: "Christian Rock",
	142: "Merengue",
	143: "Salsa",
	144: "Trash Metal",
	145: "Anime",
	146: "JPop",
	147: "Synthpop",
}

func Parser(file io.Reader, fromFolder string, targetFolder string) {
	artists := ReadCsv(file)

	targetNewFolder := filepath.Join(targetFolder, "new")

	err := os.MkdirAll(targetNewFolder, os.ModePerm)
	if err != nil {
		panic(err)
	}

	fmt.Printf("artist size: %v, target folder %v \n", len(artists), targetFolder)

	MovedSongs(artists, fromFolder, targetFolder)
}

func MovedSongs(artists map[string]Artist, fromFolder string, targetFolder string) {
	for artistName, artist := range artists {
		artistName = strings.TrimSpace(artistName)
		if artistName == "" {
			artistName = UNKNOWN
		}

		artistFolder := filepath.Join(targetFolder, artistName)

		err := os.MkdirAll(artistFolder, os.ModePerm)
		if err != nil {
			panic(err)
		}

		for albumName, songs := range artist.albums {
			albumName = strings.TrimSpace(albumName)
			if albumName == "" {
				albumName = UNKNOWN
			}

			albumFolder := filepath.Join(artistFolder, albumName)

			err := os.MkdirAll(albumFolder, os.ModePerm)
			if err != nil {
				panic(err)
			}

			for _, song := range songs {
				songPath := filepath.Join(fromFolder, song.filename)

				songFileState, err1 := os.Stat(songPath)
				if err1 != nil {
					panic(err)
				} else {
					targetSongPath := filepath.Join(albumFolder, songFileState.Name())

					ProcessFile(song, songPath, targetSongPath)
				}

			}
		}
	}
}

func ProcessFile(song Song, from string, to string) {
	if strings.HasSuffix(from, ".mp3") {
		os.Rename(from, to)
		ApplyID3v2Tags(song, to)
	} else if strings.HasSuffix(from, ".m4a") {
		mp3FilePath := strings.ReplaceAll(to, ".m4a", ".mp3")

		ConvertToMP3(from, mp3FilePath)
		os.Remove(from)
		ApplyID3v2Tags(song, mp3FilePath)
	} else {
		fmt.Printf("unsupported format from: %s to: %s \n", from, to)
	}
}

func ConvertToMP3(from string, to string) {
	commandArray := []string{
		"-i",
		from,
		"-c:v",
		"copy",
		"-c:a",
		"libmp3lame",
		"-q:a",
		"4",
		to,
	}

	fmt.Println("/usr/local/bin/ffmpeg " + strings.Join(commandArray[:], " "))

	exec.Command("/usr/local/bin/ffmpeg", commandArray[:]...).Output()
}

func ApplyID3v2Tags(song Song, path string) {
	// /usr/local/bin/id3v2 -t "The Memory Of Trees" -A "The Memory Of Trees" -a "Enya" -y 0 -g 13 resources/target/Enya/The Memory Of Trees/NLXI.m4a
	commandArray := []string{"-2", "--to-v2.4"}

	if song.title != "" {
		commandArray = append(commandArray, "-t", QuoteStr(song.title))
	}

	if song.album != "" {
		commandArray = append(commandArray, "-A", QuoteStr(song.album))
	}

	if song.artist != "" {
		commandArray = append(commandArray, "-a", QuoteStr(song.artist))
	}

	if song.yearReleased != "" && song.yearReleased != "0" {
		commandArray = append(commandArray, "-Y", song.yearReleased)
	}

	if song.genre != "" {
		genre, ok := GenreKey(song.genre)

		if ok {
			commandArray = append(commandArray, "-G", fmt.Sprintf("%d", genre))
		}
	}

	if len(commandArray) > 0 {
		commandArray = append(commandArray, path)

		// fmt.Println(absPath)

		fmt.Println("/usr/local/bin/eyeD3 " + strings.Join(commandArray[:], " "))
		exec.Command("/usr/local/bin/eyeD3", commandArray[:]...).Output()

		// if err != nil {
		// 	panic(err)
		// }
		// fmt.Println(out)
	}
}

func GenreKey(str string) (int, bool) {
	var genre int

	for k, v := range GENRE {
		if v == str {
			return k, true
		}
	}
	return genre, false
}

func QuoteStr(str string) string {
	// return strings.ReplaceAll(str, " ", "\\ ")
	return strings.TrimSpace(str)
	// return fmt.Sprintf("\"%s\"", str)
}

func ReadCsv(file io.Reader) map[string]Artist {
	reader := csv.NewReader(file)

	records, err := reader.ReadAll()

	if err != nil {
		panic(err)
	}

	var artists = make(map[string]Artist)

	fmt.Printf("size: %v \n", len(records))

	for _, record := range records[1:] {
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

		// if song.artist != "ABBA" {
		// 	continue
		// }

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

	return artists
}

func ConvertRestMp3(source string) {
	folder := filepath.Join(source, "**/*.*")

	files, err := filepath.Glob(folder)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(files)

	for _, song := range files {
		if strings.HasSuffix(song, ".mp3") {
			fmt.Printf("mp3 %s \n", song)
		} else if strings.HasSuffix(song, ".m4a") {
			mp3FilePath := strings.ReplaceAll(song, ".m4a", ".mp3")
			ConvertToMP3(song, mp3FilePath)
			os.Remove(song)
		} else {
			fmt.Printf("unsupported format from: %s \n", song)
		}
	}
}

// eyeD3 -2 --to-v2.4 -t "Thank You For The Music" -A "More ABBA Gold" -a "ABBA" -Y 2010 -g 13 "resources/target/ABBA/More ABBA Gold/CAPT.mp3"
