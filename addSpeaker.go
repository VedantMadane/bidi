package main

import (
	"encoding/json"
	"fmt"
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

type VerseEntry struct {
	Text    TextVariant  `json:"text"`
	Vaktā   *TextVariant `json:"vaktā,omitempty"`
	Book    int          `json:"book"`
	Chapter int          `json:"chapter"`
	Verse   int          `json:"verse"`
}

func processSpeaker(inputData map[string]VerseEntry) map[string]VerseEntry {
	keys := make([]string, 0, len(inputData))
	for k := range inputData {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var lastSpeaker *TextVariant
	processedData := make(map[string]VerseEntry)

	for _, key := range keys {
		entry := inputData[key]

		if isNumericKey(key) {
			lastSpeaker = &entry.Text
			processedData[key] = entry
		} else if lastSpeaker != nil {
			entry.Vaktā = lastSpeaker
			processedData[key] = entry
		}
	}

	return processedData
}

func isNumericKey(key string) bool {
	return strings.IndexAny(key, "abcdefghijklmnopqrstuvwxyz") == -1
}

func processJSONFiles(inputFolder, outputFolder string) error {
	// Create output folder if not exists
	if err := os.MkdirAll(outputFolder, os.ModePerm); err != nil {
		return err
	}

	// Read all files in input folder
	files, err := ioutil.ReadDir(inputFolder)
	if err != nil {
		return err
	}

	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".json" {
			inputPath := filepath.Join(inputFolder, file.Name())
			outputPath := filepath.Join(outputFolder, file.Name())

			// Read input JSON file
			data, err := ioutil.ReadFile(inputPath)
			if err != nil {
				log.Printf("Error reading file %s: %v", file.Name(), err)
				continue
			}

			// Unmarshal JSON
			var inputData map[string]VerseEntry
			if err := json.Unmarshal(data, &inputData); err != nil {
				log.Printf("Error unmarshaling JSON from %s: %v", file.Name(), err)
				continue
			}

			// Process JSON
			processedData := processSpeaker(inputData)

			// Marshal processed data
			outputJSON, err := json.MarshalIndent(processedData, "", "  ")
			if err != nil {
				log.Printf("Error marshaling processed data: %v", err)
				continue
			}

			// Write output file
			if err := ioutil.WriteFile(outputPath, outputJSON, 0644); err != nil {
				log.Printf("Error writing output file %s: %v", outputPath, err)
				continue
			}

			fmt.Printf("Processed %s successfully\n", file.Name())
		}
	}

	return nil
}

func main() {
	inputFolder := `D:\Projects\oper\decomposed`
	outputFolder := `D:\Projects\oper\vaktāsaha`

	if err := processJSONFiles(inputFolder, outputFolder); err != nil {
		log.Fatalf("Error processing JSON files: %v", err)
	}
}
