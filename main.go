package main

import (
	"anime-locator/anime"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Keymap of available command argument
	argslegend := map[string]bool{
		"--info":         true,
		"--list":         true,
		"--count":        true,
		"--list-genre":   true,
		"--list-key":     true,
		"--reg-id-genre": true,
		"--reg-id-date":  true,
	}

	// Initialization of AnimeLocator and Listing of Animes
	locator := anime.AnimeLocator{}
	locator.Animes = locator.LoadList()

	// If no arguments found
	if (len(os.Args) == 2 && os.Args[1] == "--help") || len(os.Args) == 1 {
		ind := "\n\t\t\t\t\t\t"

		fmt.Println("No Argument(s) found.")
		fmt.Println("\nUsage: anime-locator [--list-genre <genre-list>...] [--list-key <key-list>...]")
		fmt.Println("\t\t     [--reg-id-genre <id> <genre-list>...]")
		fmt.Println("\t\t     [--reg-id-date <id> <s_date> <e_date>]")
		fmt.Println("\t\t     [--info <id>] [--list] [--count]")

		fmt.Println("\nOptions:")
		fmt.Println("\t--list\t\t\t\t\tLists all Animes both registered and unregistered.")
		fmt.Println("\t--count\t\t\t\t\tDisplays the total count of registered Animes.")
		fmt.Println("\t--info <id>\t\t\t\tDisplays the information of an Anime according" + ind + "to the ID.\n")
		fmt.Println("\t--list-genre <genre-list>\t\tLists and filters all Animes according to the" + ind + "genre-list.\n")
		fmt.Println("\t--list-key <key-list>\t\t\tLists and filters all Animes according to the" + ind + "key-list.\n")

		fmt.Println("\t--reg-id-genre <id> <genre-list>\tRegister an Anime with its genres using its ID" + ind + "from the list.\n")
		fmt.Println("\t--reg-id-date <id> <s_date> <e_date>\tRegister an Anime with its released date using" + ind + "its ID from the list.\n")

		return
	}

	// Map of arguments
	argsmap := make(map[string][]string)

	prev := "" // Placeholder of previous command argument
	// Mapping all command arguments into the map of arguments
	for _, arg := range os.Args[1:] {
		if argslegend[arg] {
			prev = arg
		} else if prev != "" {
			argsmap[prev] = append(argsmap[prev], arg)
		} else {
			fmt.Println("Invalid Argument(s)!")
			return
		}
	}

	// Execution of all listing command arguments lexicographically
	result := locator.ListAnimeByGenre(argsmap["--list-genre"], nil)
	result = locator.ListAnimeByKeyword(argsmap["--list-key"], result)

	// Modifiying the range
	if argsmap["--list"] != nil {
		length, err := strconv.Atoi(argsmap["--list"][0])
		if length < len(result) && err == nil {
			result = result[:length]
		}
	}

	// Displaying the contents after executing the command arguments
	if prev == "--list" || argsmap["--list-genre"] != nil || argsmap["--list-key"] != nil {
		fmt.Println("\n=~=~= Anime List =~=~=")
		for i, content := range result {
			fmt.Printf("%03d\t%s\n", i+1, content.AnimeName)
		}
	}

	// Displaying the total count of Animes according to [--list] argument
	if prev == "--count" {
		fmt.Println("\nNumber of registered Anime in Library:", len(result))
	}

	// Execution of all registry command arguments lexicographically
	locator.RegisterAnimeByDate(argsmap["--reg-id-date"], result)
	locator.RegisterAnimeByGenre(argsmap["--reg-id-genre"], result)

	// Execute displaying of information of a specific Anime
	locator.DisplayAnimeInformation(argsmap["--info"], result)

	// Save instance of the list (if something was changed)
	locator.SaveList()
}
