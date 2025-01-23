package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func fetchURLs() {
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
}

// Fetch the content of the URLs
func fetchContent(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Failed to fetch URL %s: %v\n", url, err)
		return
	}
	defer resp.Body.Close()

	body, err := os.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Failed to read response body for URL %s: %v\n", url, err)
		return
	}

	fmt.Printf("Content of %s:\n%s\n", url, body)
}

// We have to convert this text file to json
// The form of the data is as follows:
// 01001000a नारायणं नमस्कृत्य नरं चैव नरोत्तमम्
// 01001000c देवीं सरस्वतीं चैव ततो जयमुदीरयेत्
// 01001001A लोमहर्षणपुत्र उग्रश्रवाः सूतः पौराणिको नैमिषारण्ये शौनकस्य कुलपतेर्द्वादशवार्षिके सत्रे
// 01001002a समासीनानभ्यगच्छद्ब्रह्मर्षीन्संशितव्रतान्
// 01001002c विनयावनतो भूत्वा कदाचित्सूतनन्दनः
// Here the 01001002c should be the key and the value should be the text that follows it until the next key is found
// The json should look like this:
// {
// 	"01001000c": "नारायणं नमस्कृत्य नरं चैव नरोत्तमम्\nदेवीं सरस्वतीं चैव ततो जयमुदीरयेत्",
// 	"01001001A": "लोमहर्षणपुत्र उग्रश्रवाः सूतः पौराणिको नैमिषारण्ये शौनकस्य कुलपतेर्द्वादशवार्षिके सत्रे",
// 	"01001002c": "समासीनानभ्यगच्छद्ब्रह्मर्षीन्संशितव्रतान्\nविनयावनतो भूत्वा कदाचित्सूतनन्दनः"
// }
// The key should be the first line and the value should be the text until the next key is found

func

func main1() {
	fetchURLs()
}