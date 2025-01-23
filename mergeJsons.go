package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type MergedJSON map[string]struct {
	Text struct {
		UD    string `json:"ud,omitempty"`
		UR    string `json:"ur,omitempty"`
		ASCII string `json:"ascii,omitempty"`
	} `json:"text"`
}

func mergeJSONFiles(file1, file2, file3 string) (MergedJSON, error) {
	// Read files
	content1, err := ioutil.ReadFile(file1)
	if err != nil {
		return nil, err
	}
	content2, err := ioutil.ReadFile(file2)
	if err != nil {
		return nil, err
	}
	content3, err := ioutil.ReadFile(file3)
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON files
	var json1, json2, json3 map[string]string
	json.Unmarshal(content1, &json1)
	json.Unmarshal(content2, &json2)
	json.Unmarshal(content3, &json3)

	// Merge JSON
	mergedJSON := make(MergedJSON)

	for key := range json1 {
		mergedJSON[key] = struct {
			Text struct {
				UD    string `json:"ud,omitempty"`
				UR    string `json:"ur,omitempty"`
				ASCII string `json:"ascii,omitempty"`
			} `json:"text"`
		}{
			Text: struct {
				UD    string `json:"ud,omitempty"`
				UR    string `json:"ur,omitempty"`
				ASCII string `json:"ascii,omitempty"`
			}{
				UD:    json1[key],
				UR:    json2[key],
				ASCII: json3[key],
			},
		}
	}

	return mergedJSON, nil
}

func main() {

	for i := 17; i <= 18; i++ {
		mergedData, err := mergeJSONFiles(fmt.Sprintf("MBh/MBh%02d.json", i), fmt.Sprintf("MBh/MBh%02dUR.json", i), fmt.Sprintf("MBh/MBh%02dASCII.json", i))
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		// Write merged JSON to file
		outputJSON, _ := json.MarshalIndent(mergedData, "", "  ")
		ioutil.WriteFile(fmt.Sprintf("troika/MBh%02d.json", i), outputJSON, 0644)
	}
	// mergedData, err := mergeJSONFiles("MBh/MBh01.json", "MBh/MBh01UR.json", "MBh/MBh01ASCII.json")
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }

	// // Write merged JSON to file
	// outputJSON, _ := json.MarshalIndent(mergedData, "", "  ")
	// ioutil.WriteFile("troika/MBh01.json", outputJSON, 0644)
}
