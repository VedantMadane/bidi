package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// Text represents the three formats of text
type Text struct {
	UD    string `json:"ud"`
	UR    string `json:"ur"`
	ASCII string `json:"ascii"`
}

// Speaker represents the speaker information
type Speaker map[string]string

// VerseData represents the structure of each verse in the input files
type VerseData struct {
	Vakta Speaker `json:"vaktā"`
	Text  Text    `json:"text"`
}

// SpeakerInfo represents the final output structure for each speaker
type SpeakerInfo struct {
	Speaker Speaker `json:"speaker"`
	Text    Text    `json:"text"`
}

// createSpeakerKey creates a consistent key for speaker mapping
func createSpeakerKey(speaker Speaker) string {
	// Convert map entries to slice for sorting
	pairs := make([]string, 0, len(speaker))
	for k, v := range speaker {
		pairs = append(pairs, fmt.Sprintf("%s:%s", k, strings.TrimSpace(v)))
	}
	sort.Strings(pairs)
	return strings.Join(pairs, "|")
}

func processMBHFiles() {
	// Initialize speakers map to store accumulated text
	speakers := make(map[string]Text)
	keyMapping := make(map[string]string)
	counter := 1

	// Process files MBh01.json through MBh18.json
	for bookNum := 1; bookNum <= 18; bookNum++ {
		filePath := filepath.Join("D:", "Projects", "oper", "vaktāsaha",
			fmt.Sprintf("MBh%02d.json", bookNum))

		// Read file
		data, err := os.ReadFile(filePath)
		if err != nil {
			fmt.Printf("Warning: Could not read file %s: %v\n", filePath, err)
			continue
		}

		// Parse JSON
		var verses map[string]VerseData
		if err := json.Unmarshal(data, &verses); err != nil {
			fmt.Printf("Error parsing JSON from %s: %v\n", filePath, err)
			continue
		}

		// Process each verse
		for _, verse := range verses {
			// Skip if no speaker information
			if verse.Vakta == nil {
				continue
			}

			// Create speaker key
			speakerKey := createSpeakerKey(verse.Vakta)

			// Assign new number if this is a new speaker
			if _, exists := keyMapping[speakerKey]; !exists {
				keyMapping[speakerKey] = fmt.Sprintf("%d", counter)
				counter++
				// Initialize text for new speaker
				speakers[keyMapping[speakerKey]] = Text{}
			}

			// Get current speaker number
			speakerNum := keyMapping[speakerKey]

			// Update text
			currentText := speakers[speakerNum]
			currentText.UD += verse.Text.UD
			currentText.UR += verse.Text.UR
			currentText.ASCII += verse.Text.ASCII
			speakers[speakerNum] = currentText
		}
	}

	// Create final output structure
	speakerInfo := make(map[string]SpeakerInfo)
	for speakerKey, number := range keyMapping {
		// Convert speaker key back to map
		speakerPairs := strings.Split(speakerKey, "|")
		speaker := make(Speaker)
		for _, pair := range speakerPairs {
			if parts := strings.Split(pair, ":"); len(parts) == 2 {
				speaker[parts[0]] = parts[1]
			}
		}

		speakerInfo[number] = SpeakerInfo{
			Speaker: speaker,
			Text:    speakers[number],
		}
	}

	// Write output file
	outputPath := filepath.Join("D:", "Projects", "oper", "vaktāsaha", "go_by_speaker.json")
	outputData, err := json.MarshalIndent(speakerInfo, "", "  ")
	if err != nil {
		fmt.Printf("Error creating JSON output: %v\n", err)
		return
	}

	if err := os.WriteFile(outputPath, outputData, 0644); err != nil {
		fmt.Printf("Error writing output file: %v\n", err)
		return
	}

	fmt.Printf("Successfully created %s\n", outputPath)
	fmt.Printf("Total speakers processed: %d\n", len(speakerInfo))
}

func main() {
	processMBHFiles()
}
