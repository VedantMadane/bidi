package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

type Anuloma struct {
	Anulom   string `json:"anulom"`
	Pratilom string `json:"pratilom"`
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
func getAnulomPratilom() (string, string) {
	directory, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	filePath := directory + "\\anulom.json"
	fmt.Println(filePath, ":: getting json from here")
	file, err := os.ReadFile(filePath)
	file = bytes.ReplaceAll(file, []byte("\r"), []byte(""))
	var data map[string]string
	err = json.Unmarshal(file, &data)
	if err != nil {
		log.Fatal(err)
	}
	anulom := data["anulom"]
	pratilom := data["pratilom"]
	return anulom, pratilom
}

// Function to get anulom.json object from the file
func getAnulom() string {
	// Get the present workiing directiry
	directory, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	// Print wd
	filePath := directory + "\\anulom.json"
	// Print filePath
	fmt.Println(filePath)
	// Read the file
	file, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	// file = bytes.ReplaceAll(file, []byte("\r", []byte("")))
	// Unmarshal the json object
	var anulom Anuloma
	err = json.Unmarshal(file, &anulom)
	if err != nil {
		log.Printf("Error unmarshalling JSON: %v", err)
		return ""
	}
	// Return the anulom object
	return anulom.Anulom
}

// anulom.json is of the form:
// Multi-line comment	// {"anulom": "वन्देऽहं देवं तं श्रीतं रन्तारं कालं भासा यः ।
// रामो रामाधीराप्यागो लीलामारायोध्ये वासे ॥",
// Add error handling for file not found
// "pratilom": "सेवाध्येयो रामालाली गोप्याराधी मारामोरा ।
// यस्साभालंकारं तारं तं श्रीतं वन्देहं देवं ॥"
// }

// Function to strip ॥ and । and ऽ from the string in one line
func stripAnulom(s string) string {
	s = strings.Replace(s, "॥", "", -1)
	s = strings.Replace(s, "।", "", -1)
	s = strings.Replace(s, "ऽ", "", -1)
	s = strings.Replace(s, " ", "", -1)
	return s
}

// Separate the shlok into syllables

func main() {
	anulom, pratilom := getAnulomPratilom()
	s := stripAnulom(anulom)
	reversed := reverseString(s)
	fmt.Println(reversed, ":: reversed string\n", pratilom)
	compare(reversed, pratilom)
}

// Compare reversed and pratilom and show where it is different
func compare(reversed string, pratilom string) {
	for i := 0; i < len(reversed); i++ {
		if reversed[i] != pratilom[i] {
			fmt.Println("Difference at index", i, ":", reversed[i], "vs", pratilom[i])
			// wHAT TO DO WITH THESE Difference at index

		}
	}
}

// Write the reversed back to the json
// func writePratilom(reversed string) {
// 	directory, err := os.Getwd()
// 	if err != nil {
// 		log.Printf(("Error getting working directory"))
// 	}
// 	filePath := directory + "\\anulom.json"
// 	// Read the file
// 	_, err := os.ReadFile(filePath)
// 	if err != nil {
// 		log.Printf("error reading file: %v", err)
// 		return
// 	}

// }
