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
	// WordCount() returns word count in PDF-document
	word_count, err := pdf.WordCount()
	if err != nil {
		log.Fatal(err)
	}
	// Print
	fmt.Println("Word count:", word_count)
	// Close() releases allocated resources for PDF-document
	defer pdf.Close()
}
