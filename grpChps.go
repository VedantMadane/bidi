package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type Entry struct {
	Text    map[string]string `json:"text"`
	Book    int               `json:"book"`
	Chapter int               `json:"chapter"`
	Verse   int               `json:"verse"`
}

func groupUniqueChapters(directory string) error {
	// Create output directory for grouped chapters
	outputDir := filepath.Join(directory, "grouped_chapters")
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return err
	}

	// Read JSON files
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		return err
	}

	// Map to store chapters grouped by chapter number
	chapterGroups := make(map[int]map[string]Entry)

	for _, file := range files {
		filename := file.Name()
		if strings.HasPrefix(filename, "MBh") && strings.HasSuffix(filename, ".json") {
			filepath := filepath.Join(directory, filename)

			// Read file
			data, err := ioutil.ReadFile(filepath)
			if err != nil {
				fmt.Printf("Error reading %s: %v\n", filename, err)
				continue
			}

			// Unmarshal JSON
			fileData := make(map[string]Entry)
			if err := json.Unmarshal(data, &fileData); err != nil {
				fmt.Printf("Error parsing %s: %v\n", filename, err)
				continue
			}

			// Group entries by chapter
			for key, entry := range fileData {
				// Ensure chapter group exists
				if _, exists := chapterGroups[entry.Chapter]; !exists {
					chapterGroups[entry.Chapter] = make(map[string]Entry)
				}

				// Add entry to chapter group
				chapterGroups[entry.Chapter][key] = entry
			}
		}
	}

	// Write grouped chapters
	for chapterNumber, chapterEntries := range chapterGroups {
		outputFilepath := filepath.Join(outputDir, fmt.Sprintf("Chapter_%03d.json", chapterNumber))

		// Write JSON file
		jsonData, err := json.MarshalIndent(chapterEntries, "", "  ")
		if err != nil {
			fmt.Printf("Error marshaling JSON for Chapter %d: %v\n", chapterNumber, err)
			continue
		}

		if err := ioutil.WriteFile(outputFilepath, jsonData, 0644); err != nil {
			fmt.Printf("Error writing file for Chapter %d: %v\n", chapterNumber, err)
		}
	}

	fmt.Printf("Grouped %d unique chapters\n", len(chapterGroups))
	return nil
}

func main() {
	directory := `D:\Projects\oper\decomposed`
	if err := groupUniqueChapters(directory); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
