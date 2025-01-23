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
	// Create main output directory for grouped chapters
	outputDir := filepath.Join(directory, "by_chapters")
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return err
	}

	// Read JSON files
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		return err
	}

	// Map to store chapters grouped by book and chapter number
	bookChapterGroups := make(map[int]map[int]map[string]Entry)

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

			// Group entries by book and chapter
			for key, entry := range fileData {
				// Ensure book group exists
				if _, exists := bookChapterGroups[entry.Book]; !exists {
					bookChapterGroups[entry.Book] = make(map[int]map[string]Entry)
				}

				// Ensure chapter group exists
				if _, exists := bookChapterGroups[entry.Book][entry.Chapter]; !exists {
					bookChapterGroups[entry.Book][entry.Chapter] = make(map[string]Entry)
				}

				// Add entry to book and chapter group
				bookChapterGroups[entry.Book][entry.Chapter][key] = entry
			}
		}
	}

	// Write grouped chapters
	for bookNumber, bookChapters := range bookChapterGroups {
		// Create book-specific directory
		bookOutputDir := filepath.Join(outputDir, fmt.Sprintf("Book_%02d", bookNumber))
		if err := os.MkdirAll(bookOutputDir, 0755); err != nil {
			fmt.Printf("Error creating book directory: %v\n", err)
			continue
		}

		// Write chapters for this book
		for chapterNumber, chapterEntries := range bookChapters {
			outputFilepath := filepath.Join(bookOutputDir, fmt.Sprintf("Chapter_%03d.json", chapterNumber))

			// Write JSON file
			jsonData, err := json.MarshalIndent(chapterEntries, "", "  ")
			if err != nil {
				fmt.Printf("Error marshaling JSON for Book %d, Chapter %d: %v\n", bookNumber, chapterNumber, err)
				continue
			}

			if err := ioutil.WriteFile(outputFilepath, jsonData, 0644); err != nil {
				fmt.Printf("Error writing file for Book %d, Chapter %d: %v\n", bookNumber, chapterNumber, err)
			}
		}
	}

	fmt.Printf("Grouped chapters for %d books\n", len(bookChapterGroups))
	return nil
}

func main() {
	directory := `D:\Projects\oper\decomposed`
	if err := groupUniqueChapters(directory); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
