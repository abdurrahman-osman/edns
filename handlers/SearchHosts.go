package handlers

import (
	"edns/models"
	"fmt"
	"sort"
	"strings"
)

func SearchHosts(query string) {
	hosts := readHostsFile()
	if len(hosts) == 0 {
		fmt.Println("No hosts found.")
		return
	}

	var searchResults []models.HostEntry

	// Check if the query is an IP address or hostname substring
	isIP := isPotentialIP(query)

	// Sort the hosts slice based on IP or hostname
	if isIP {
		sort.Sort(models.ByIP(hosts))
	} else {
		sort.Sort(models.ByHostname(hosts))
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
