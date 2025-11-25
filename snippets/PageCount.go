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
	// PageCount() returns page count in PDF-document
	count, err := pdf.PageCount()
	if err != nil {
		log.Fatal(err)
	}
	// Print
	fmt.Println("Count:", count)
}
