package main

import (
	"bufio"
	"fmt"
	"os"
	"os/user"
	"sort"
	"strconv"
	"strings"
)

type HostEntry struct {
	ID       int
	IP       string
	Hostname string
}

type ByIP []HostEntry
type ByHostname []HostEntry

func (a ByIP) Len() int           { return len(a) }
func (a ByIP) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByIP) Less(i, j int) bool { return lessIP(a[i].IP, a[j].IP) }

func (a ByHostname) Len() int           { return len(a) }
func (a ByHostname) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByHostname) Less(i, j int) bool { return a[i].Hostname < a[j].Hostname }

func main() {
	// Check if the program is running as root
	if !isRoot() {
		fmt.Println("Error: This program must be run as root.")
		return
	}

	// Read command-line arguments
	args := os.Args[1:]

	// Check if the first argument is "list"
	if len(args) > 0 && args[0] == "list" {
		listHosts()
	} else if len(args) > 0 && args[0] == "search" {
		if len(args) < 2 {
			fmt.Println("Error: Insufficient arguments for search.")
			return
		}
		searchHosts(args[1])
	} else {
		showInfo()
	}
}

func isRoot() bool {
	currentUser, err := user.Current()
	if err != nil {
		fmt.Println("Error:", err)
		return false
	}

	// Check if the program is running as root
	return currentUser.Username == "root"
}

func listHosts() {
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

func searchHosts(query string) {
	hosts := readHostsFile()
	if len(hosts) == 0 {
		fmt.Println("No hosts found.")
		return
	}

	var searchResults []HostEntry

	// Check if the query is an IP address or hostname substring
	isIP := isPotentialIP(query)

	// Sort the hosts slice based on IP or hostname
	if isIP {
		sort.Sort(ByIP(hosts))
	} else {
		sort.Sort(ByHostname(hosts))
	}

	// Perform search based on IP or hostname
	for _, entry := range hosts {
		match := false
		if isIP {
			if strings.Contains(entry.IP, query) {
				match = true
			}
		} else {
			if strings.Contains(entry.Hostname, query) {
				match = true
			}
		}

		if match {
			searchResults = append(searchResults, entry)
		}
	}

	// Check if there are at least three matches before printing
	if len(searchResults) < 1 {
		fmt.Println("No matching hosts found.")
		return
	}

	// Print the search results
	for _, entry := range searchResults {
		fmt.Printf("ID: %d\tIP: %s\tHostname: %s\n", entry.ID, entry.IP, entry.Hostname)
	}
}

func readHostsFile() []HostEntry {
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
	var hosts []HostEntry

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
			entry := HostEntry{
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

func lessIP(ip1, ip2 string) bool {
	parts1 := strings.Split(ip1, ".")
	parts2 := strings.Split(ip2, ".")
	for i := 0; i < 4; i++ {
		num1 := atoi(parts1[i])
		num2 := atoi(parts2[i])
		if num1 < num2 {
			return true
		} else if num1 > num2 {
			return false
		}
	}
	return false
}

func atoi(s string) int {
	v, _ := strconv.Atoi(s)
	return v
}

func showInfo() {
	// Print some information
	fmt.Println("No action specified. Showing program information.")
	fmt.Println("This program reads the /etc/hosts file and extracts IP addresses and hostnames.")
	fmt.Println("To list the IP addresses and hostnames, run the program with 'list' as the first argument.")
	fmt.Println("To search for IP addresses or hostnames, run the program with 'search <query>' as the arguments.")
}

func isPotentialIP(s string) bool {
	// Check if the string contains only numbers and dots
	for _, c := range s {
		if c != '.' && (c < '0' || c > '9') {
			return false
		}
	}
	return true
}
