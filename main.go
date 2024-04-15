package main

import (
	"edns/handlers"
	"fmt"
	"os"
)

func main() {
	// Check if the program is running as root
	if !handlers.IsRoot() {
		fmt.Println("Error: This program must be run as root.")
		return
	}

	// Read command-line arguments
	args := os.Args[1:]

	// Check if the first argument is "list"
	if len(args) > 0 && args[0] == "list" {
		handlers.ListHosts()
	} else if len(args) > 0 && args[0] == "search" {
		if len(args) < 2 {
			fmt.Println("Error: Insufficient arguments for search.")
			return
		}
		handlers.SearchHosts(args[1])
	} else {
		handlers.ShowInfo()
	}
}
