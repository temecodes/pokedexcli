package main

import (
	"fmt"
	"bufio"
	//"strings"
	"os"
)

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
	}
}

func commandExit() error{
	fmt.Printf("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}

func commandHelp() error{
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\nhelp: Displays a help message\nexit: Exit the Pokedex\n")
	return nil
}

func main() {
	
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
