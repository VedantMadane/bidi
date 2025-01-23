package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type Speakers map[string]int

func extractQuotes(baseDir string) error {
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

	// Prepare list of speaker keys
	speakerKeys := make([]string, 0, len(speakers))
	for speaker := range speakers {
		speakerKeys = append(speakerKeys, strings.TrimSpace(speaker))
	}

	// Process UR text files
	bookChaptersDir := filepath.Join(baseDir, "decomposed", "by_chapters")
	return filepath.Walk(bookChaptersDir, func(path string, info os.FileInfo, err error) error {
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
			quotes := make(map[string]string)
			lines := strings.Split(string(content), "\n")
			var currentQuote strings.Builder
			var quoteCounter = make(map[string]int)

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
						count := quoteCounter[matchedSpeaker]
						quoteKey := fmt.Sprintf("%s_%d", matchedSpeaker, count)
						quotes[quoteKey] = strings.TrimSpace(currentQuote.String())
						quoteCounter[matchedSpeaker]++
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
				for _, speaker := range speakerKeys {
					if strings.Contains(currentQuote.String(), speaker) {
						count := quoteCounter[speaker]
						quoteKey := fmt.Sprintf("%s_%d", speaker, count)
						quotes[quoteKey] = strings.TrimSpace(currentQuote.String())
						break
					}
				}
			}

			// Write quotes to JSON
			if len(quotes) > 0 {
				// Determine output path
				relPath, _ := filepath.Rel(bookChaptersDir, path)
				bookDir := filepath.Dir(filepath.Dir(relPath))
				outputDir := filepath.Join(baseDir, "by_speaker", bookDir)
				os.MkdirAll(outputDir, 0755)

				outputPath := filepath.Join(outputDir, "quotes.json")
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
	if err := extractQuotes(baseDir); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
