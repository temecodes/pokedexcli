package main

import (
	"fmt"
	"bufio"
	//"strings"
	"os"
)


func main() {

	scanner := bufio.NewScanner(os.Stdin)

	for{
		fmt.Print("Pokedex > ")
		if !scanner.Scan(){
			break
		}
		text := scanner.Text()
		cleanText := cleanInput(text)
		if len(cleanText) == 0{
			continue
		}
		fmt.Printf("Your command was: %s\n", cleanText[0])

	}
	return

}
