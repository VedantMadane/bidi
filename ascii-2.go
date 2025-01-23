// Write every third, fifth, ...odd line to a different key in a map and write the map to a json file. Line delimiter ||
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// "io/ioutil"

func main() {
	// Open the file
	file, err := os.Open("D:\\Projects\\oper\\hk-ascii.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a map to store the lines
	lines := make(map[string]string)

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Read the file line by line
	for scanner.Scan() {
		line := scanner.Text()
		// Split the line by the delimiter
		parts := strings.Split(line, "||")
		// Add the line to the map
		lines[parts[0]] = parts[1]
		// Print the line
		fmt.Println(line)
	}

	// Close the scanner
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Marshal the map to JSON
	jsonData, err := json.Marshal(lines)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	// Write the JSON to a file
	err = os.WriteFile("output1.json", jsonData, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}

	fmt.Println("JSON file created successfully!")
}
