
package main

import(
	"strings"
	"fmt"
	"os"
	"bufio"
)

func cleanInput(text string) []string{
	lower := strings.ToLower(text)
	words := strings.Fields(lower)
	return words
}

type cliCommand struct{
	name string
	description string
	callback func(*config) error
}

type config struct{
	nextLoc *string
	prevLoc *string
}

func getCommands() map[string]cliCommand{
	return map[string]cliCommand{
		"exit":{
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},
		"help":{
			name: "help",
			description: "Displays a help message",
			callback: commandHelp,
		},
		"map":{
			name:"map",
			description: "Tells the name of different areas",
			callback: commandMap,
		},
		"next":{
			name:"next",
			description: "Gives the next 20 locations",
			callback: commandMap,
		},
	}
}


func startRepl() {
	startURL := "https://pokeapi.co/api/v2/location-area"

	cfg := &config{
		nextLoc: &startURL,
		prevLoc: nil,
	}

	scanner := bufio.NewScanner(os.Stdin)
	gettingCommands:= getCommands()

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

		commandSaid := cleanText[0]
		
		command, ok := gettingCommands[commandSaid]
		if !ok{
			fmt.Println("Unknown command")
			continue
		}	

		err := command.callback(cfg)
		if err != nil{
			fmt.Println(err)
		}
		
	}
	return

}
