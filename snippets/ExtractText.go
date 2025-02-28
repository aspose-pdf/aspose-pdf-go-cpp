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
	// ExtractText() returns PDF-document contents as plain text
	txt, err := pdf.ExtractText()
	if err != nil {
		log.Fatal(err)
	}
	// Print
	fmt.Println("Extracted text:\n", txt)
	// Close() releases allocated resources for PDF-document
	defer pdf.Close()
}
