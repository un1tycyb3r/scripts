package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

// Check if a string matches the domain pattern
func isValidDomain(domain string) bool {
	domainRegex := `^([a-zA-Z0-9-]+\.)+([a-zA-Z]{2,63})$`
	match, _ := regexp.MatchString(domainRegex, domain)
	return match
}

// Get the apex domain from a valid domain string
func getApexDomain(domain string) string {
	parts := strings.Split(domain, ".")
	if len(parts) < 2 {
		return domain
	}
	return strings.Join(parts[len(parts)-2:], ".")
}

func main() {
	// Open the file
	file, err := os.Open("domains.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create a new Scanner for the file
	scanner := bufio.NewScanner(file)

	// Loop over all lines in the file
	for scanner.Scan() {
		line := scanner.Text()
		if isValidDomain(line) {
			apex := getApexDomain(line)
			// Print the apex domain
			fmt.Println(apex)
		}
	}

	// Check if there were errors during scanning
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
