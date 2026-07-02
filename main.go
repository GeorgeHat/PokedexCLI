package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	allowedCommands := getCommands()

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		input := scanner.Text()
		cleanedInput := cleanInput(input)
		firstWord := cleanedInput[0]
		command, ok := allowedCommands[firstWord]
		if !ok {
			println("Unknown command")
			continue
		}
		err := command.callback()
		if err != nil {
			fmt.Printf("Error excecuting command: %v: %v", command.name, err)
		}

	}

}
