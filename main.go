package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			break
		}

		arrText := cleanInput(scanner.Text())

		if len(arrText) > 0 {
			firstword := arrText[0]
			// check where arrText[0] (first value) in cliCommand.name

			command, exists := getCommands()[firstword]

			if exists {
				err := command.callback()
				if err != nil {
					fmt.Println("Error: %v", err)
				}

			} else {
				fmt.Println("Unknown Command")
			}
		}

	}
}
