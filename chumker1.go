package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strconv"
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

func sortAndSplitJSON() error {
	// Read input file
	inputPath := `D:\Projects\oper\vaktāsaha\go_by_speaker.json`
	data, err := ioutil.ReadFile(inputPath)
	if err != nil {
		return fmt.Errorf("error reading input file: %v", err)
	}

	// Unmarshal JSON
	var entries map[string]Entry
	if err := json.Unmarshal(data, &entries); err != nil {
		return fmt.Errorf("error parsing JSON: %v", err)
	}

	// Sort keys
	sortedKeys := make([]string, 0, len(entries))
	for k := range entries {
		sortedKeys = append(sortedKeys, k)
	}
	sort.Slice(sortedKeys, func(i, j int) bool {
		// Convert keys to integers for proper numerical sorting
		numI, _ := strconv.Atoi(sortedKeys[i])
		numJ, _ := strconv.Atoi(sortedKeys[j])
		return numI < numJ
	})

	// Prepare output directory
	outputDir := `D:\Projects\oper\by_speaker1`
	os.MkdirAll(outputDir, 0755)

	// Prepare text type directories
	asciiDir := filepath.Join(outputDir, "ascii")
	dvDir := filepath.Join(outputDir, "dv")
	drDir := filepath.Join(outputDir, "dr")
	os.MkdirAll(asciiDir, 0755)
	os.MkdirAll(dvDir, 0755)
	os.MkdirAll(drDir, 0755)

	// Split into files
	const maxFileSize = 25 * 1024 * 1024 // 25 MB
	fileCounter := 0

	asciiEntries := make(map[string]Entry)
	dvEntries := make(map[string]Entry)
	drEntries := make(map[string]Entry)

	currentSize := 0

	for _, key := range sortedKeys {
		entry := entries[key]

		// Add to respective type entries
		asciiEntries[key] = entry
		dvEntries[key] = entry
		drEntries[key] = entry

		// Check file size
		jsonData, _ := json.Marshal(map[string]Entry{key: entry})
		currentSize += len(jsonData)

		// Write files if size exceeds limit
		if currentSize >= maxFileSize {
			fileCounter++

			// Write ASCII file
			asciiPath := filepath.Join(asciiDir, fmt.Sprintf("part_%03d.json", fileCounter))
			if err := writeJSON(asciiPath, asciiEntries); err != nil {
				return err
			}

			// Write DV file
			dvPath := filepath.Join(dvDir, fmt.Sprintf("part_%03d.json", fileCounter))
			if err := writeJSON(dvPath, dvEntries); err != nil {
				return err
			}

			// Write DR file
			drPath := filepath.Join(drDir, fmt.Sprintf("part_%03d.json", fileCounter))
			if err := writeJSON(drPath, drEntries); err != nil {
				return err
			}

			// Reset
			asciiEntries = make(map[string]Entry)
			dvEntries = make(map[string]Entry)
			drEntries = make(map[string]Entry)
			currentSize = 0
		}
	}

	// Write final files if not empty
	if len(asciiEntries) > 0 {
		fileCounter++

		asciiPath := filepath.Join(asciiDir, fmt.Sprintf("part_%03d.json", fileCounter))
		if err := writeJSON(asciiPath, asciiEntries); err != nil {
			return err
		}

		dvPath := filepath.Join(dvDir, fmt.Sprintf("part_%03d.json", fileCounter))
		if err := writeJSON(dvPath, dvEntries); err != nil {
			return err
		}

		drPath := filepath.Join(drDir, fmt.Sprintf("part_%03d.json", fileCounter))
		if err := writeJSON(drPath, drEntries); err != nil {
			return err
		}
	}

	fmt.Printf("Processed %d files\n", fileCounter)
	return nil
}

func writeJSON(path string, data map[string]Entry) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(data)
}

func main() {
	if err := sortAndSplitJSON(); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
