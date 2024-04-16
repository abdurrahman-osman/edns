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
		fmt.Printf("Line: %d\tID: %d\tIP: %s\tHostname: %s\n", entry.Line, entry.ID, entry.IP, entry.Hostname)
	}
}
