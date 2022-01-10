package anime

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

type AnimeLocator struct {
	Animes []*Anime
}

var AnimeDirectory = [...]string{"D:\\`Anime Storage\\Anime", "D:\\`Anime Storage"}

func (locator *AnimeLocator) ListAnimeByGenre(genres []string, optlist []*Anime) []*Anime {
	list := locator.Animes
	if genres == nil { // If the genres argument is empty, just return the list
		return list
	}

	// Filter according to the genre arguments
	for i, lsize, n := 0, 0, len(genres); i < n; i++ {
		lsize = len(list) // Original length each filter iteration
		for _, anime := range list {
			for _, genre := range anime.AnimeGenre {
				ngenre := strings.Join(strings.Split(genres[i], "-"), " ")
				if strings.ToLower(genre) == ngenre ||
					strings.ToUpper(genre) == ngenre {
					list = append(list, anime)
				}
			}
		}
		// Easy work around for filter by removing elems with regards to the original length
		list = list[lsize:]
	}
	return list
}

func (locator *AnimeLocator) ListAnimeByKeyword(keywords []string, optlist []*Anime) []*Anime {
	list := optlist
	if keywords == nil { // If the keywords argument is empty, just return the list
		return list
	}
	// Filter according to the keyword arguments
	for i, lsize, n := 0, 0, len(keywords); i < n; i++ {
		lsize = len(list) // Original length each filter iteration
		for _, anime := range list {
			nkey := strings.Join(strings.Split(keywords[i], "-"), " ")
			if strings.Contains(strings.ToLower(anime.AnimeName), nkey) ||
				strings.Contains(strings.ToUpper(anime.AnimeName), nkey) {
				list = append(list, anime)
			}
		}
		// Easy work around for filter by removing elems with regards to the original length
		list = list[lsize:]
	}
	return list
}

func (AnimeLocator) RegisterAnimeByGenre(genres []string, optlist []*Anime) {
	list := optlist
	if len(genres) < 2 { // If the keywords argument is empty, just return the list
		return
	}

	// Getting the ID from the list
	id, _ := strconv.Atoi(genres[0])

	// Formatting the genre into a proper Title format
	for i, genre := range genres[1:] {
		genres[i+1] = strings.Title(strings.ToLower(strings.Join(strings.Split(genre, "-"), " ")))
	}
	// Overwriting the current list of genre and outputs it
	sort.Strings(genres[1:])
	list[id-1].AnimeGenre = genres[1:]
	fmt.Println("\n=~=~= Anime Modification Information =~=~=")
	fmt.Println("Anime Name:\t", list[id-1].AnimeName)
	fmt.Println("Genres:\t\t", strings.Join(list[id-1].AnimeGenre, ", "))
	fmt.Println("")
}

func (AnimeLocator) RegisterAnimeByDate(dates []string, optlist []*Anime) {
	list := optlist
	if len(dates) != 3 { // If the keywords argument is empty, just return the list
		return
	}

	// Month Format
	dtf := [...]string{
		"Jan", "Feb", "Mar", "Apr", "May", "Jun",
		"Jul", "Aug", "Sept", "Oct", "Nov", "Dec",
	}

	// Parsing the 'dates' slice into its components
	id, _ := strconv.Atoi(dates[0])
	s_date, e_date := strings.Split(dates[1], "-"), strings.Split(dates[2], "-")

	// Parsing the integer date of the month
	dtf_s, _ := strconv.Atoi(s_date[0])
	dtf_e, _ := strconv.Atoi(e_date[0])

	// Formatting the starting and ending date of the released date
	output := (dtf[dtf_s-1] + " " + s_date[1] + ", " + s_date[2])
	output += " to "
	output += (dtf[dtf_e-1] + " " + e_date[1] + ", " + e_date[2])

	// Overwriting the current released date and outputs it
	list[id-1].ReleaseDate = output
	fmt.Println("\n=~=~= Anime Modification Information =~=~=")
	fmt.Println("Anime Name:\t", list[id-1].AnimeName)
	fmt.Println("Released Date:\t", list[id-1].ReleaseDate)
	fmt.Println("")
}

func (AnimeLocator) DisplayAnimeInformation(ids []string, optlist []*Anime) {
	list := optlist
	if ids == nil { // If the keywords argument is empty, just return the list
		return
	}

	// Parsing the ID and locates the Anime according to the ID
	nid, _ := strconv.Atoi(ids[0])
	if nid > len(list) {
		return
	}
	anime := list[nid-1]

	// Displays the information of the specified Anime
	fmt.Println("\n=~=~= Anime Information =~=~=")
	fmt.Println("Anime Name:\t", anime.AnimeName)
	fmt.Println("Anime Genre:\t", strings.Join(anime.AnimeGenre, ", "))
	fmt.Println("Released Date:\t", anime.ReleaseDate)
	fmt.Println("")
}

func (AnimeLocator) LoadList() []*Anime {
	// Accessing the necessary files/directories
	anilist, _ := os.ReadDir(AnimeDirectory[0])
	genrelist, _ := os.ReadFile(AnimeDirectory[1] + "\\AnimeMasterList.txt")

	// Slice of Animes
	animes := make([]*Anime, len(anilist))
	// Slice of Animes with registered information (genre and released date)
	reglist := strings.Split(string(genrelist), "=+=")
	sort.Strings(reglist) // Sort in ascending order (if something was out of arrangement)

	// Storing all Anime available in the directory
	for i, elem := range anilist {
		animes[i] = &Anime{}
		animes[i].AnimeName = elem.Name()[1:]
	}

	curid := 0 // Placeholder for the ID that is registered with information
	for _, elem := range reglist {
		p := strings.Split(elem, "|")           // Slice of information
		for ; curid < len(animes); curid += 1 { // Scanning the list with regards to already registered Anime
			if animes[curid].AnimeName == p[0] {
				animes[curid].AnimeGenre = strings.Split(p[1], ",")
				animes[curid].ReleaseDate = strings.TrimSpace(p[2])
				break
			}
		}
	}

	return animes
}

func (locator *AnimeLocator) SaveList() {
	// Delimeters-> Newline:"=+=", Separator:"|", Indices:","
	list := ""

	for _, anime := range locator.Animes {
		list += ("=+=" + anime.AnimeName + "|" + strings.Join(anime.AnimeGenre, ",") + "|" + anime.ReleaseDate + "\n")
	}

	list = list[3:]

	// Creation/Overwrite of the MasterList file
	file, _ := os.Create(AnimeDirectory[1] + "\\AnimeMasterList.txt")

	// Writing the new information to the Master file
	io.WriteString(file, list)
}
