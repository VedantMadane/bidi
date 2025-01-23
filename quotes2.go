package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type Speakers map[string]int

func processChapters(baseDir string) error {
	// Read speakers from counts.json
	speakersFile := filepath.Join(baseDir, "vaktāsaha", "counts.json")
	speakersData, err := ioutil.ReadFile(speakersFile)
	if err != nil {
		return fmt.Errorf("error reading speakers file: %v", err)
	}

	var speakers Speakers
	if err := json.Unmarshal(speakersData, &speakers); err != nil {
		return fmt.Errorf("error parsing speakers JSON: %v", err)
	}

	// Sort speakers by line count (descending)
	speakerKeys := make([]string, 0, len(speakers))
	for speaker := range speakers {
		speakerKeys = append(speakerKeys, strings.TrimSpace(speaker))
	}
	sort.Slice(speakerKeys, func(i, j int) bool {
		return speakers[speakerKeys[i]] > speakers[speakerKeys[j]]
	})

	// Prepare output quotes file
	outputFile := filepath.Join(baseDir, "by_speaker", "quotes.json")

	// Read existing quotes if file exists
	var existingQuotes map[int]string
	existingData, err := ioutil.ReadFile(outputFile)
	if err == nil {
		if err := json.Unmarshal(existingData, &existingQuotes); err != nil {
			return fmt.Errorf("error parsing existing quotes: %v", err)
		}
	} else {
		existingQuotes = make(map[int]string)
	}

	// Determine next quote counter
	var nextQuoteCounter int
	for k := range existingQuotes {
		if k > nextQuoteCounter {
			nextQuoteCounter = k
		}
	}
	nextQuoteCounter++

	// Process UR text files
	byChaptersDir := filepath.Join(baseDir, "decomposed", "by_chapters")
	return filepath.Walk(byChaptersDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Process only UR text files
		if !info.IsDir() && strings.HasSuffix(path, "_ur.txt") {
			// Read file content
			content, err := ioutil.ReadFile(path)
			if err != nil {
				return fmt.Errorf("error reading file %s: %v", path, err)
			}

			lines := strings.Split(string(content), "\n")
			var currentQuote strings.Builder

			for _, line := range lines {
				matchedSpeaker := ""
				for _, speaker := range speakerKeys {
					if strings.Contains(line, speaker) {
						matchedSpeaker = speaker
						break
					}
				}

				if matchedSpeaker != "" {
					// Save previous quote if exists
					if currentQuote.Len() > 0 {
						existingQuotes[nextQuoteCounter] = strings.TrimSpace(currentQuote.String())
						nextQuoteCounter++
						currentQuote.Reset()
					}

					// Start new quote
					currentQuote.WriteString(line + "\n")
				} else if currentQuote.Len() > 0 {
					currentQuote.WriteString(line + "\n")
				}
			}

			// Save last quote
			if currentQuote.Len() > 0 {
				existingQuotes[nextQuoteCounter] = strings.TrimSpace(currentQuote.String())
				nextQuoteCounter++
			}
		}

		return nil
	})

	// Write updated quotes to file
	quotesJSON, err := json.MarshalIndent(existingQuotes, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshaling quotes: %v", err)
	}

	if err := ioutil.WriteFile(outputFile, quotesJSON, 0644); err != nil {
		return fmt.Errorf("error writing quotes file: %v", err)
	}

	fmt.Printf("Updated quotes file: %s\n", outputFile)
	return nil
}

func main() {
	baseDir := `D:\Projects\oper`
	if err := processChapters(baseDir); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
