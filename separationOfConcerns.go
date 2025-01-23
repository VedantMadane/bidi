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
	Text struct {
		ASCII string `json:"ascii"`
		UD    string `json:"ud"`
		UR    string `json:"ur"`
	} `json:"text"`
	Book    int `json:"book"`
	Chapter int `json:"chapter"`
	Verse   int `json:"verse"`
}

func processChapters(baseDirectory string) error {
	// Walk through all book directories
	return filepath.Walk(baseDirectory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Check if it's a chapter JSON file
		if !info.IsDir() && strings.HasSuffix(path, ".json") {
			// Read chapter JSON file
			data, err := ioutil.ReadFile(path)
			if err != nil {
				return fmt.Errorf("error reading file %s: %v", path, err)
			}

			// Unmarshal JSON
			chapterData := make(map[string]Entry)
			if err := json.Unmarshal(data, &chapterData); err != nil {
				return fmt.Errorf("error parsing JSON in %s: %v", path, err)
			}

			// Prepare output files
			dirPath := filepath.Dir(path)
			chapterName := strings.TrimSuffix(filepath.Base(path), ".json")

			asciiFile, err := os.Create(filepath.Join(dirPath, chapterName+"_ascii.txt"))
			if err != nil {
				return err
			}
			defer asciiFile.Close()

			udFile, err := os.Create(filepath.Join(dirPath, chapterName+"_ud.txt"))
			if err != nil {
				return err
			}
			defer udFile.Close()

			urFile, err := os.Create(filepath.Join(dirPath, chapterName+"_ur.txt"))
			if err != nil {
				return err
			}
			defer urFile.Close()

			// Write text to respective files
			for _, entry := range chapterData {
				asciiFile.WriteString(entry.Text.ASCII)
				udFile.WriteString(entry.Text.UD)
				urFile.WriteString(entry.Text.UR)
			}

			fmt.Printf("Processed %s\n", path)
		}

		return nil
	})
}

func main() {
	baseDirectory := `D:\Projects\oper\decomposed\by_chapters`

	if err := processChapters(baseDirectory); err != nil {
		fmt.Printf("Error processing chapters: %v\n", err)
	}
}
