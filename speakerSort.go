package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func extractQuotes(baseDirectory string) error {
	return filepath.Walk(baseDirectory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Process only UD text files
		if !info.IsDir() && strings.HasSuffix(path, "_ud.txt") {
			// Read file content
			content, err := ioutil.ReadFile(path)
			if err != nil {
				return fmt.Errorf("error reading file %s: %v", path, err)
			}

			// Prepare quotes map
			quotes := make(map[string]string)

			// Speakers to look for
			speakers := []string{
				"सूत उवाच",
				"तमृषय ऊचुः",
				"शौनक उवाच",
				"अग्निरुवाच",
				// Add more speakers as needed
			}

			// Read file line by line
			scanner := bufio.NewScanner(strings.NewReader(string(content)))
			var currentQuote strings.Builder
			var speakerCount = make(map[string]int)

			for scanner.Scan() {
				line := scanner.Text()

				// Check for speaker lines
				for _, speaker := range speakers {
					if strings.Contains(line, speaker) {
						// If previous quote exists, save it
						if currentQuote.Len() > 0 {
							speakerKey := speaker
							count := speakerCount[speaker]
							speakerCount[speaker]++

							if count > 0 {
								speakerKey = fmt.Sprintf("%s_%d", speaker, count)
							}

							quotes[speakerKey] = strings.TrimSpace(currentQuote.String())
							currentQuote.Reset()
						}

						// Start new quote with speaker line
						currentQuote.WriteString(line + "\n")
						break
					}
				}

				// Append line to current quote
				if currentQuote.Len() > 0 {
					currentQuote.WriteString(line + "\n")
				}
			}

			// Save last quote if exists
			if currentQuote.Len() > 0 {
				for _, speaker := range speakers {
					if strings.Contains(currentQuote.String(), speaker) {
						speakerKey := speaker
						count := speakerCount[speaker]
						speakerCount[speaker]++

						if count > 0 {
							speakerKey = fmt.Sprintf("%s_%d", speaker, count)
						}

						quotes[speakerKey] = strings.TrimSpace(currentQuote.String())
						break
					}
				}
			}

			// Write quotes to JSON
			if len(quotes) > 0 {
				quotesPath := strings.Replace(path, "_ud.txt", "_quotes.json", 1)
				quotesJSON, err := json.MarshalIndent(quotes, "", "  ")
				if err != nil {
					return fmt.Errorf("error marshaling quotes: %v", err)
				}

				if err := ioutil.WriteFile(quotesPath, quotesJSON, 0644); err != nil {
					return fmt.Errorf("error writing quotes file: %v", err)
				}

				fmt.Printf("Created quotes file: %s\n", quotesPath)
			}
		}

		return nil
	})
}

func main() {
	baseDirectory := `D:\Projects\oper\decomposed\by_chapters`

	if err := extractQuotes(baseDirectory); err != nil {
		fmt.Printf("Error extracting quotes: %v\n", err)
	}
}
