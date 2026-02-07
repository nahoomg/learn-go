package main

import (
	"fmt"
	"strings"
)

func FilterLogs(logs []string) []string {
	// Define it locally so it's fresh every time the function is called
	var filtered []string 

	defer fmt.Println("--- Log processing cycle complete ---")

	for _, v := range logs { // Parentheses around logs are optional in Go
		if strings.Contains(v, "ERROR") || strings.Contains(v, "CRITICAL") {
			filtered = append(filtered, v)
		}
	}
	return filtered
}

func main() {
	rawLogs := []string{
		"INFO: User login",
		"ERROR: Database connection failed",
		"DEBUG: Port 8080 open",
		"CRITICAL: Disk full",
	}

	// Capture the returned slice into a local variable
	result := FilterLogs(rawLogs)

	fmt.Printf("Length: %d, Capacity: %d\n", len(result), cap(result))
	fmt.Println("Filtered Logs:", result)
}