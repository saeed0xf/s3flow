package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/cheggaaa/pb/v3"
)

// GenerateWordlist generates permutations of bucket names based on environments
func GenerateWordlist(commonPrefix string, orgName string, environments []string) []string {
	var list []string
	// Raw permutation
	list = append(list, fmt.Sprintf("%s-%s", orgName, commonPrefix))
	list = append(list, fmt.Sprintf("%s.%s", orgName, commonPrefix))

	// Environment permutations
	for _, env := range environments {
		formats := []string{
			"%s-%s-%s", "%s-%s.%s", "%s-%s%s", "%s.%s-%s", "%s.%s.%s",
		}
		for _, format := range formats {
			list = append(list, fmt.Sprintf(format, orgName, commonPrefix, env))
		}
	}

	// Host permutations
	formats := []string{"%s.%s", "%s-%s", "%s%s"}
	for _, format := range formats {
		list = append(list, fmt.Sprintf(format, orgName, commonPrefix))
		list = append(list, fmt.Sprintf(format, commonPrefix, orgName))
	}

	// Remove duplicates
	uniqueList := make(map[string]struct{})
	for _, item := range list {
		uniqueList[item] = struct{}{}
	}
	var result []string
	for key := range uniqueList {
		result = append(result, key)
	}
	return result
}

// ReadWordlistFromFile reads a wordlist from a file
func ReadWordlistFromFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, strings.TrimSpace(scanner.Text()))
	}
	return lines, scanner.Err()
}

// SaveWordlistToFile saves the generated wordlist to a file with a progress bar
func SaveWordlistToFile(wordlist []string, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	// Create a progress bar
	bar := pb.StartNew(len(wordlist))
	defer bar.Finish()

	for _, bucket := range wordlist {
		_, err := writer.WriteString(bucket + "\n")
		if err != nil {
			return err
		}
		bar.Increment()
		time.Sleep(1 * time.Millisecond) // Progress bar visualization
	}

	writer.Flush()
	return nil
}

func main() {
	// Define command-line arguments
	wordlistFile := flag.String("w", "", "Path to the wordlist file (common prefixes)")
	orgName := flag.String("org", "", "Organization name")
	outputFile := flag.String("o", "generated_wordlist.txt", "Output file name")
	mediumFlag := flag.Bool("medium", false, "Use medium-sized environment list")
	largeFlag := flag.Bool("large", false, "Use large environment list")
	flag.Parse()

	// Validate required arguments
	if *wordlistFile == "" || *orgName == "" {
		fmt.Println("Usage: s3flow -w <wordlist_file> -org <organization_name> [-o <output_file>] [-medium] [-large]")
		return
	}

	// Define environment lists
	smallEnvironments := []string{
		"dev", "development", "stage", "s3", "staging", "prod", "production", "test",
	}
	mediumEnvironments := []string{
		"dev", "development", "devel", "develop", "local", "localhost", "test", "testing", "debug",
		"stage", "staging", "preprod", "pre-production", "uat", "qa", "demo", "sandbox", "mock",
		"prod", "production", "live", "real", "main",
	}
	largeEnvironments := []string{
		"dev", "development", "devel", "develop", "local", "localhost", "test", "testing", "debug",
		"stage", "staging", "preprod", "pre-production", "uat", "qa", "demo", "sandbox", "mock",
		"prod", "production", "live", "real", "main",
		"backup", "archive", "archived", "bkp", "snapshots", "snapshot",
		"aws", "azure", "gcp", "s3", "cloudfront", "cdn", "storage",
		"internal", "intranet", "private", "secure", "confidential", "restricted", "admin", "management", "ops",
		"temp", "temporary", "exp", "experiment", "trial", "beta", "alpha",
		"us-east", "us-west", "eu-central", "ap-southeast", "global", "region1", "region2",
		"api", "rest", "graphql", "rpc", "service", "microservice", "backend", "frontend",
		"logs", "logging", "metrics", "monitoring", "analytics", "reporting", "cache", "cdn", "static", "media", "assets", "files", "uploads", "downloads", "shared", "public", "external",
	}

	// Select environment list based on flags
	var environments []string
	if *largeFlag {
		environments = largeEnvironments
	} else if *mediumFlag {
		environments = mediumEnvironments
	} else {
		environments = smallEnvironments
	}

	// Read wordlist file
	commonPrefixes, err := ReadWordlistFromFile(*wordlistFile)
	if err != nil {
		fmt.Printf("Error reading wordlist file: %v\n", err)
		return
	}

	// Generate wordlist for each common prefix
	var wordlist []string
	for _, commonPrefix := range commonPrefixes {
		wordlist = append(wordlist, GenerateWordlist(commonPrefix, *orgName, environments)...)
	}

	fmt.Printf("Generated wordlist with %d items.\n", len(wordlist))

	// Save the wordlist to a file
	err = SaveWordlistToFile(wordlist, *outputFile)
	if err != nil {
		fmt.Printf("Error saving wordlist to file: %v\n", err)
		return
	}

	fmt.Printf("Wordlist saved to '%s'.\n", *outputFile)
}
