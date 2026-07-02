package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	allowedCommands := getCommands()
	config := &config{empty: true}
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			fmt.Printf("Error reading stdin: %v\n", scanner.Err())
			break
		}

		input := scanner.Text()
		cleanedInput := cleanInput(input)
		firstWord := cleanedInput[0]
		command, ok := allowedCommands[firstWord]
		if !ok {
			println("Unknown command")
			continue
		}

		err := command.callback(config)
		if err != nil {
			fmt.Printf("Error excecuting command: %v: %v\n", command.name, err)
		}

	}

}
