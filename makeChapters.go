package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type ChapterData struct {
	// Define your JSON structure here
	// Adjust fields based on actual JSON structure
	Chapter string `json:"chapter"`
	// Add other relevant fields
}

func main() {
	sourceDir := `D:\Projects\oper\MBh`
	outputDir := filepath.Join(sourceDir, "processed_chapters")

	// Create output directory if not exists
	os.MkdirAll(outputDir, 0755)

	// Organized chapters map
	chapterGroups := make(map[string][]ChapterData)

	// Iterate through MBh01.json to MBh18.json
	for i := 1; i <= 18; i++ {
		filename := fmt.Sprintf("MBh%02d.json", i)
		filePath := filepath.Join(sourceDir, filename)

		// Read JSON file
		data, err := ioutil.ReadFile(filePath)
		if err != nil {
			fmt.Printf("Error reading %s: %v\n", filename, err)
			continue
		}

		var chapters []ChapterData
		err = json.Unmarshal(data, &chapters)
		if err != nil {
			fmt.Printf("Error parsing %s: %v\n", filename, err)
			continue
		}

		// Determ

		// Group chapters
		for _, chapter := range chapters {
			chapterKey := extractChapterKey(chapter.Chapter)
			chapterGroups[chapterKey] = append(chapterGroups[chapterKey], chapter)
		}

		// Create book-specific directory
		bookDir := filepath.Join(outputDir, fmt.Sprintf("Book_%02d", i))
		os.MkdirAll(bookDir, 0755)

		// Write chapter files
		for chapterKey, chapterData := range chapterGroups {
			chapterFile := filepath.Join(bookDir, chapterKey+".json")
			chapterJSON, _ := json.MarshalIndent(chapterData, "", "  ")
			ioutil.WriteFile(chapterFile, chapterJSON, 0644)
		}

		// Clear for next iteration
		chapterGroups = make(map[string][]ChapterData)
	}
}

func extractChapterKey(chapter string) string {
	// Extract chapter number from key format
	if len(chapter) >= 8 {
		return chapter[2:8]
	}
	return chapter
}
