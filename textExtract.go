// I've a PDF with Sanskrit text in it. I want to extract the text from the PDF and write it to a file.
package main

import (
	"fmt"
	"os"

	"github.com/ledongthuc/pdf"
)

func main() {
	file, err := os.Open("D:/Downloads/dokumen.tips_maitrayani-samhita-of-yajurveda.pdf")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}
	reader, err := pdf.NewReader(file, fileInfo.Size())
	if err != nil {
		fmt.Println(err)
		return
	}

	numPages := reader.NumPage()
	for i := 1; i <= numPages; i++ {
		page := reader.Page(i)
		text := page.Content()
		fmt.Println(text)
	}
}
