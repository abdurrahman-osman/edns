package handlers

import "fmt"

func ListHosts() {
	hosts := readHostsFile()
	if len(hosts) == 0 {
		fmt.Println("No hosts found.")
		return
	}

	// Print the host entries
	for _, entry := range hosts {
		fmt.Printf("ID: %d\tIP: %s\tHostname: %s\n", entry.ID, entry.IP, entry.Hostname)
	}
}
