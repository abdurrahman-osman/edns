package handlers

func isPotentialIP(s string) bool {
	// Check if the string contains only numbers and dots
	for _, c := range s {
		if c != '.' && (c < '0' || c > '9') {
			return false
		}
	}
	return true
}
