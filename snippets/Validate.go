package main

import (
	"fmt"
	"github.com/aspose-pdf/aspose-pdf-go-cpp"
	"log"
)

func main() {
	// Open(filename string) opens a PDF-document with filename
	pdf, err := asposepdf.Open("sample.pdf")
	if err != nil {
		log.Fatal(err)
	}
	// Close() releases allocated resources for PDF-document
	defer pdf.Close()

	// Validate(pdfFormat PdfFormat) validates PDF-document for PDF/A compliance
	ok, logStr, err := pdf.Validate(asposepdf.PDF_A_2A)
	if err != nil {
		log.Fatal("Validate PDF/A error:", err)
	}

	// Print validation result and full log
	fmt.Printf("Validate PDF/A result: %v\n", ok)
	fmt.Printf("Validate PDF/A log:\n%s\n", logStr)
}
