package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

const (
	inputFile   = `D:\Projects\oper\vaktāsaha\go_by_speaker.json`
	outputDir   = `D:\Projects\oper\by_speaker\`
	maxFileSize = 25 * 1024 * 1024 // 25 MB
)

type Entry struct {
	Speaker struct {
		ASCII string `json:"ascii"`
		UD    string `json:"ud"`
		UR    string `json:"ur"`
	} `json:"speaker"`
	Text struct {
		ASCII string `json:"ascii"`
		UD    string `json:"ud"`
		UR    string `json:"ur"`
	} `json:"text"`
}

func splitLargeJSON() error {
	// Read entire JSON file
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		return fmt.Errorf("error reading input file: %v", err)
	}

	// Unmarshal JSON
	var fullData map[string]Entry
	if err := json.Unmarshal(data, &fullData); err != nil {
		return fmt.Errorf("error parsing JSON: %v", err)
	}

	// Ensure output directory exists
	os.MkdirAll(outputDir, 0755)

	// Split into files
	currentFile := make(map[string]Entry)
	fileCounter := 1
	currentSize := 0

	for key, value := range fullData {
		// Add entry to current file
		currentFile[key] = value

		// Check file size
		jsonData, _ := json.Marshal(currentFile)
		currentSize = len(jsonData)

		// If file exceeds max size, write and reset
		if currentSize >= maxFileSize {
			outputPath := filepath.Join(outputDir, fmt.Sprintf("part_%03d.json", fileCounter))

			if err := ioutil.WriteFile(outputPath, jsonData, 0644); err != nil {
				return fmt.Errorf("error writing file %s: %v", outputPath, err)
			}

			fmt.Printf("Created file: %s (Size: %d bytes)\n", outputPath, currentSize)

			// Reset for next file
			currentFile = make(map[string]Entry)
			fileCounter++
			currentSize = 0
		}
	}

	// Write last file if not empty
	if len(currentFile) > 0 {
		outputPath := filepath.Join(outputDir, fmt.Sprintf("part_%03d.json", fileCounter))
		jsonData, _ := json.Marshal(currentFile)

		if err := ioutil.WriteFile(outputPath, jsonData, 0644); err != nil {
			return fmt.Errorf("error writing final file %s: %v", outputPath, err)
		}

		fmt.Printf("Created final file: %s (Size: %d bytes)\n", outputPath, len(jsonData))
	}

	return nil
}

func main() {
	if err := splitLargeJSON(); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
