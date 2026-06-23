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
	// GetMetaInfo(key string) gets meta information value of PDF-document
	value, err := pdf.GetMetaInfo("Author")
	if err != nil {
		log.Fatal(err)
	}
	// Print
	fmt.Println("Author: ", value)
}
