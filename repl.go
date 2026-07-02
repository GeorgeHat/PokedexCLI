package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
	empty bool
}

func cleanInput(text string) []string {
	if text == "" {
		return []string{""}
	}
	return strings.Fields(strings.ToLower(text))

}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the names of 20 location areas in the Pokemon world. Each subsequent call to map displays thhe next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the names of 20 previous areas",
			callback:    commandMapb,
		},
	}
}

func commandExit(config *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(config *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	for _, command := range getCommands() {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}

func commandMap(locationArea *config) error {
	url := ""
	if !locationArea.empty {
		if locationArea.Next == nil {
			fmt.Println("you're on the last page")
			return nil
		}
		url = *locationArea.Next
	} else {
		url = "https://pokeapi.co/api/v2/location-area/"
		locationArea.empty = false
	}

	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, locationArea)
	if err != nil {
		return err
	}

	for _, result := range locationArea.Results {
		fmt.Println(result.Name)
	}

	return nil
}

func commandMapb(locationArea *config) error {
	url := ""
	if !locationArea.empty {
		if locationArea.Previous == nil {
			fmt.Println("you're in the first page")
			return nil
		}
		url = *locationArea.Previous
	} else {
		return fmt.Errorf("Incorrect usage: you need to use map first at least once")
	}

	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, locationArea)
	if err != nil {
		return err
	}

	for _, result := range locationArea.Results {
		fmt.Println(result.Name)
	}

	return nil
}
