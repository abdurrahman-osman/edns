package handlers

import "fmt"

func ShowInfo() {
	// Print some information
	fmt.Println("No action specified. Showing program information.")
	fmt.Println("This program reads the /etc/hosts file and extracts IP addresses and hostnames.")
	fmt.Println("To list the IP addresses and hostnames, run the program with 'list' as the first argument.")
	fmt.Println("To search for IP addresses or hostnames, run the program with 'search <query>' as the arguments.")
}
