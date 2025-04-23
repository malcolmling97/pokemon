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
			fmt.Printf("Your command was: %v \n", arrText[0])
		}

	}
}
