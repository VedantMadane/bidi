package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
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

func processFiles(baseDir string) error {
	// Ensure output directory exists
	outputDir := filepath.Join(baseDir, "text_types")
	os.MkdirAll(outputDir, 0755)

	// Process 8 files
	for i := 1; i <= 8; i++ {
		filename := fmt.Sprintf("part_%03d.json", i)
		fullPath := filepath.Join(baseDir, filename)

		// Read file
		data, err := ioutil.ReadFile(fullPath)
		if err != nil {
			return fmt.Errorf("error reading %s: %v", filename, err)
		}

		// Unmarshal JSON
		var entries map[string]Entry
		if err := json.Unmarshal(data, &entries); err != nil {
			return fmt.Errorf("error parsing %s: %v", filename, err)
		}

		// Prepare output files
		asciiFile, err := os.Create(filepath.Join(outputDir, fmt.Sprintf("ascii_%d.json", i)))
		if err != nil {
			return err
		}
		defer asciiFile.Close()

		dvFile, err := os.Create(filepath.Join(outputDir, fmt.Sprintf("dv_%d.json", i)))
		if err != nil {
			return err
		}
		defer dvFile.Close()

		drFile, err := os.Create(filepath.Join(outputDir, fmt.Sprintf("dr_%d.json", i)))
		if err != nil {
			return err
		}
		defer drFile.Close()

		// Prepare output maps
		asciiEntries := make(map[string]string)
		dvEntries := make(map[string]string)
		drEntries := make(map[string]string)

		// Populate entries
		for key, entry := range entries {
			asciiEntries[key] = entry.Text.ASCII
			dvEntries[key] = entry.Text.UD
			drEntries[key] = entry.Text.UR
		}

		// Write JSON files
		writeJSON(asciiFile, asciiEntries)
		writeJSON(dvFile, dvEntries)
		writeJSON(drFile, drEntries)

		fmt.Printf("Processed %s\n", filename)
	}

	return nil
}

func writeJSON(file *os.File, data interface{}) error {
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(data)
}

func main() {
	baseDir := `D:\Projects\oper\by_speaker`
	if err := processFiles(baseDir); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
