package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type TextVariant struct {
	Ud    string `json:"ud"`
	Ur    string `json:"ur"`
	Ascii string `json:"ascii"`
}

type Entry struct {
	Text  TextVariant  `json:"text"`
	Vaktā *TextVariant `json:"vaktā,omitempty"`
}

func processJSONWithSpeaker(inputFolder, outputFolder string) error {
	// Ensure output folder exists
	if err := os.MkdirAll(outputFolder, os.ModePerm); err != nil {
		return err
	}

	// Read all JSON files in input folder
	files, err := ioutil.ReadDir(inputFolder)
	if err != nil {
		return err
	}

	for _, file := range files {
		if !strings.HasSuffix(file.Name(), ".json") {
			continue
		}

		// Full file paths
		inputPath := filepath.Join(inputFolder, file.Name())
		outputPath := filepath.Join(outputFolder, file.Name())

		// Read input file
		data, err := ioutil.ReadFile(inputPath)
		if err != nil {
			log.Printf("Error reading file %s: %v", file.Name(), err)
			continue
		}

		// Unmarshal JSON
		var jsonData map[string]Entry
		if err := json.Unmarshal(data, &jsonData); err != nil {
			log.Printf("Error unmarshaling JSON: %v", err)
			continue
		}

		// Sort keys
		var keys []string
		for k := range jsonData {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		var lastSpeaker *TextVariant

		// Process keys
		for _, key := range keys {
			entry := jsonData[key]

			// If numeric key, update last speaker
			if isNumeric(key) {
				lastSpeaker = &entry.Text
			}

			// If alphabetic key and last speaker exists
			if !isNumeric(key) && lastSpeaker != nil {
				entry.Vaktā = lastSpeaker
				jsonData[key] = entry
			}
		}

		// Marshal processed data
		processedJSON, err := json.MarshalIndent(jsonData, "", "  ")
		if err != nil {
			log.Printf("Error marshaling JSON: %v", err)
			continue
		}

		// Write to output file
		if err := ioutil.WriteFile(outputPath, processedJSON, 0644); err != nil {
			log.Printf("Error writing file %s: %v", outputPath, err)
		}
	}

	return nil
}

func isNumeric(s string) bool {
	for _, c := range s {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}

func main() {
	inputFolder := `D:\Projects\oper\decomposed`
	outputFolder := `D:\Projects\oper\vaktāsaha`

	if err := processJSONWithSpeaker(inputFolder, outputFolder); err != nil {
		log.Fatal(err)
	}
}
