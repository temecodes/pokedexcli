
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
	callback func() error
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
			callback: commandMap
		}
	}
}


func startRepl() {
	
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

		err := command.callback()
		if err != nil{
			fmt.Println(err)
		}
		
	}
	return

}
