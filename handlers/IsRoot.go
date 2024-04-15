package handlers

import (
	"fmt"
	"os/user"
)

func IsRoot() bool {
	currentUser, err := user.Current()
	if err != nil {
		fmt.Println("Error:", err)
		return false
	}

	// Check if the program is running as root
	return currentUser.Username == "root"
}
