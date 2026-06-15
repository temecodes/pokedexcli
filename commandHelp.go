


package main

import(
	"fmt"
)
func commandHelp(cfg *config) error{
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\nhelp: Displays a help message\nexit: Exit the Pokedex\nmap: Tells the name of different areas\n")
	return nil
}
