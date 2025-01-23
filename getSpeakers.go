package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func extractSpeakers(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = '\t' // Tab-separated values
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		return nil
	}

	speakers := make([]string, 0, len(records)-1)
	for i, record := range records[1:] { // Skip header
		if len(record) > 0 {
			speakers = append(speakers, strings.TrimSpace(record[i]))
		}
	}

	return speakers
}

func main() {
	filename := `D:\Projects\oper\Speaker Number of Lines Percentage.yaml`
	speakers := extractSpeakers(filename)

	fmt.Println("Extracted Speakers:")
	for i, speaker := range speakers {
		fmt.Printf("%d. %s\n", i+1, speaker)
	}
}
