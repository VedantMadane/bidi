package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
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

func countSpeakers(inputFolder string) (map[string]int, error) {
	speakerCounts := make(map[string]int)

	// Iterate through files MBh01.json to MBh18.json
	for i := 1; i <= 18; i++ {
		filename := fmt.Sprintf("MBh%02d.json", i)
		filepath := filepath.Join(inputFolder, filename)

		// Read file
		data, err := ioutil.ReadFile(filepath)
		if err != nil {
			log.Printf("Error reading file %s: %v", filename, err)
			continue
		}

		// Unmarshal JSON
		var inputData map[string]VerseEntry
		if err := json.Unmarshal(data, &inputData); err != nil {
			log.Printf("Error unmarshaling JSON from %s: %v", filename, err)
			continue
		}

		// Count speakers
		for _, entry := range inputData {
			if entry.Vaktā != nil {
				speaker := entry.Vaktā.Ur
				speakerCounts[speaker]++
			}
		}
	}

	return speakerCounts, nil
}

func main() {
	inputFolder := `D:\Projects\oper\vaktāsaha`
	outputFile := filepath.Join(inputFolder, "counts.json")

	// Count speakers
	speakerCounts, err := countSpeakers(inputFolder)
	if err != nil {
		log.Fatalf("Error counting speakers: %v", err)
	}

	// Write counts to JSON
	countsJSON, err := json.MarshalIndent(speakerCounts, "", "  ")
	if err != nil {
		log.Fatalf("Error marshaling counts: %v", err)
	}

	if err := ioutil.WriteFile(outputFile, countsJSON, 0644); err != nil {
		log.Fatalf("Error writing counts file: %v", err)
	}

	fmt.Println("Speaker counts written to counts.json")
}
