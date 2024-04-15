package handlers

import (
	"bufio"
	"edns/models"
	"fmt"
	"os"
	"strings"
)

func readHostsFile() []models.HostEntry {
	// Open the /etc/hosts file
	file, err := os.Open("/etc/hosts")
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Slice to store host entries
	var hosts []models.HostEntry

	// Counter to track line numbers
	lineNumber := 1

	// Iterate through each line of the file
	for scanner.Scan() {
		line := scanner.Text()

		// Skip comments and empty lines
		if strings.HasPrefix(line, "#") || len(line) == 0 {
			lineNumber++
			continue
		}

		// Split the line by whitespace
		fields := strings.Fields(line)

		// Extract IP address and hostnames
		ip := fields[0]
		hostnames := fields[1:]

		// Create a HostEntry struct for each IP and hostname pair
		for _, hostname := range hostnames {
			entry := models.HostEntry{
				ID:       lineNumber,
				IP:       ip,
				Hostname: hostname,
			}
			hosts = append(hosts, entry)
		}

		lineNumber++
	}

	return hosts
}
