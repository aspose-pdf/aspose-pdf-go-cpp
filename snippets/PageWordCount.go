package main

import "github.com/aspose-pdf/aspose-pdf-go-cpp"
import "log"
import "fmt"

func main() {
	// Open(filename string) opens a PDF-document with filename
	pdf, err := asposepdf.Open("sample.pdf")
	if err != nil {
		log.Fatal(err)
	}
	// Close() releases allocated resources for PDF-document
	defer pdf.Close()
	// PageWordCount(num int32) returns word count on specified page in PDF-document.
	page_word_count, err := pdf.PageWordCount(1)
	if err != nil {
		log.Fatal(err)
	}
	// Print
	fmt.Println("Word count on the first page:", page_word_count)
}
