package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

func convertToJSON(filename string) (map[string]string, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// Regular expression to match keys and their text
	re := regexp.MustCompile(`(0\d{6}\w+)\s*(.*?)(?=0\d{6}\w+|$)`)
	matches := re.FindAllStringSubmatch(string(content), -1)

	result := make(map[string]string)
	for _, match := range matches {
		result[match[1]] = strings.TrimSpace(match[2])
	}

	return result, nil
}

func fetchAndPrintURLs() string {
	baseURL := "https://bombay.indology.info/mahabharata/text/UD/MBh%02d.txt"
	for i := 1; i <= 18; i++ {
		url := fmt.Sprintf(baseURL, i)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("Failed to fetch URL %s: %v\n", url, err)
			continue
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("Failed to read response body for URL %s: %v\n", url, err)
			continue
		}

		fmt.Printf("Content of %s:\n%s\n", url, body)
	}
	return "MBh18.txt"
}

// Return the content of the URL
func fetchContent(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Failed to fetch URL %s: %v\n", url, err)
		return string(resp.Status)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	// body, err := os.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Failed to read response body for URL %s: %v\n", url, err)

	}

	fmt.Printf("Content of %s:\n%s\n", url, body)
	return string(body)
}

// Save the content of the URL to a file
func saveContentToFile(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Failed to fetch URL %s: %v\n", url, err)
		return "MBh18.txt"
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Failed to read response body for URL %s: %v\n", url, err)

		return ""
	}
	fileAddress := "MBh18.txt"
	err = ioutil.WriteFile(fileAddress, body, 0644)
	if err != nil {
		fmt.Printf("Failed to save content to file: %v\n", err)
		return ""
	}

	fmt.Printf("Content of %s saved to MBh18.txt\n", url, fileAddress)
	return fileAddress
}

func main() {
	// Fetch and print URLs
	// fetchAndPrintURLs()
	// Call fetchContent function from bomIndo.go
	// content := fetchContent("https://bombay.indology.info/mahabharata/text/UD/MBh01.txt")
	// Save the content to a file
	fileAdd := saveContentToFile("https://bombay.indology.info/mahabharata/text/UD/MBh01.txt")
	// Convert the content to JSON
	fmt.Print(fileAdd)
	jsonData, err := convertToJSON(fileAdd)
	if err != nil {
		fmt.Println("Error:", err)
		return
	} else {
		fmt.Println("JSON Data:", jsonData)
	}

	output, _ := json.MarshalIndent(jsonData, "", "  ")
	fmt.Println(string(output))
}
