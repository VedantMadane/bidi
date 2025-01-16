// Create ascii.json
// Every second time you encounter || in hk-ascii.txt, create a new object in the json file named "1", "2" and so on uptil "30"
// The object should have the following
// First time you encounter || in hk-ascii.txt, make a new string in the object named "1" and so on uptil "30" with the key anulom
// Second time you encounter || in hk-ascii.txt, make a new string in the object named "1" and so on uptil "30" with the key vilom
// The value of anulom and vilom should be the text between the || and the next ||
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func Ascii() {
	// Open the file
	file, err := os.Open("D:\\Projects\\oper\\hk-ascii.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("File opened", &file)
	}
	defer file.Close()
	// Create a new scanner and read the file line by line
	scanner := bufio.NewScanner(file)
	// Create a new map to store the data
	data := make(map[string]map[string]string)
	// Create a counter to keep track of the object number
	counter := 1
	// Create a new object in the map
	data[string(counter)] = make(map[string]string)
	// Create a new string in the object
	data[string(counter)]["anulom"] = ""
	data[string(counter)]["vilom"] = ""
	// Create a new string to keep track of the current key
	currentKey := "anulom"
	// Loop through the file
	for scanner.Scan() {
		// Split the line by ||
		parts := strings.Split(scanner.Text(), "||")
		// If the length of the parts is less than 2, continue
		fmt.Printf("parts: %v\n", parts)
		if len(parts) < 2 {

			continue
		}
		// If the current key is anulom, set the current key to vilom
		if currentKey == "anulom" {
			currentKey = "vilom"
			// Set the value of the current key to the text between the || and the next ||
			data[string(counter)][currentKey] = parts[1]
			continue
		}
		// If the current key is vilom, set the current key to anulom
		if currentKey == "vilom" {
			currentKey = "anulom"
			// Set the value of the current key to the text between the || and the next ||
			data[string(counter)][currentKey] = parts[1]
			// Increment the counter
			counter++
			// Create a new object in the map
			data[string(counter)] = make(map[string]string)
			// Create a new string in the object
			data[string(counter)]["anulom"] = ""
			data[string(counter)]["vilom"] = ""
		}
		// Set the value of the current key to the text between the || and the next ||
		data[string(counter)][currentKey] = parts[1]
	}
	// Create a new file
	jsonFile, err := os.Create("D:\\Projects\\oper\\ascii1.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	// Create a new encoder
	encoder := json.NewEncoder(jsonFile)
	// Encode the data
	encoder.Encode(data)
}
func main() {
	Ascii()
}
