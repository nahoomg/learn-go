package main

import (
	"fmt"
	"strings"
)

func FilterLogs(logs []string) []string {
	var filtered []string 

	defer fmt.Println("--- Log processing cycle complete ---")

	for _, v := range logs { 
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

	result := FilterLogs(rawLogs)

	fmt.Printf("Length: %d, Capacity: %d\n", len(result), cap(result))
	fmt.Println("Filtered Logs:", result)
}