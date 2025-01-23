// Go program to extract all lines ending with | and || from https://www.sadagopan.org/ebook/pdf/Raghava%20Yadaveeyam.pdf
package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	// URL of the PDF file
	url := "https://www.sadagopan.org/ebook/pdf/Raghava%20Yadaveeyam.pdf"

	// Send an HTTP GET request to the URL
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error downloading the PDF:", err)
		return
	}
	defer resp.Body.Close()

	// Create a temporary file to store the PDF content
	tempFile, err := os.CreateTemp("", "temp-pdf-*.pdf")
	if err != nil {
		fmt.Println("Error creating temporary file:", err)
		return
	}
	defer os.Remove(tempFile.Name()) // Remove the temporary file when done

	// Copy the PDF content to the temporary file
	_, err = io.Copy(tempFile, resp.Body)
	if err != nil {
		fmt.Println("Error copying PDF content to temporary file:", err)
		return
	}

	// Open the temporary file for reading
	file, err := os.Open(tempFile.Name())
	if err != nil {
		fmt.Println("Error opening temporary file:", err)
		return
	}
	defer file.Close()

	// Create a new file to store the extracted lines
	outputFile, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()

	// Create a scanner to read the temporary file line by line
	scanner := bufio.NewScanner(file)

	// Iterate through each line in the temporary file
	for scanner.Scan() {
		line := scanner.Text()

		// Check if the line ends with | or ||
		if strings.HasSuffix(line, "|") || strings.HasSuffix(line, "||") {
			// Write the line to the output file
			_, err := outputFile.WriteString(line + "\n")
			if err != nil {
				fmt.Println("Error writing to output file:", err)
				return
			}
		}
	}

	// Check for any errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading temporary file:", err)
		return
	}

	fmt.Println("Lines ending with | and || extracted and saved to output.txt")
}
