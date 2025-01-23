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

	// Process UR text files
	byChaptersDir := filepath.Join(baseDir, "decomposed", "by_chapters")
	outputDir := filepath.Join(baseDir, "by_speaker")
	os.MkdirAll(outputDir, 0755)

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

			// Prepare quotes map
			quotes := make(map[int]string)
			lines := strings.Split(string(content), "\n")

			quoteCounter := 1
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
						quotes[quoteCounter] = strings.TrimSpace(currentQuote.String())
						quoteCounter++
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
				quotes[quoteCounter] = strings.TrimSpace(currentQuote.String())
			}

			// Determine output path
			relPath, _ := filepath.Rel(byChaptersDir, path)
			bookDir := filepath.Dir(filepath.Dir(relPath))
			outputPath := filepath.Join(outputDir, bookDir, "quotes.json")

			// Ensure directory exists
			os.MkdirAll(filepath.Dir(outputPath), 0755)

			// Write quotes to JSON
			if len(quotes) > 0 {
				quotesJSON, _ := json.MarshalIndent(quotes, "", "  ")
				if err := ioutil.WriteFile(outputPath, quotesJSON, 0644); err != nil {
					return fmt.Errorf("error writing quotes file: %v", err)
				}

				fmt.Printf("Created quotes file: %s\n", outputPath)
			}
		}

		return nil
	})
}

func main() {
	baseDir := `D:\Projects\oper`
	if err := processChapters(baseDir); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
