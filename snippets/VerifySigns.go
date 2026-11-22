package main

import (
	"fmt"
	"github.com/aspose-pdf/aspose-pdf-go-cpp"
	"log"
)

func main() {
	// Open(filename string) opens a PDF-document with filename
	pdf, err := asposepdf.Open("sample_with_sign.pdf")
	if err != nil {
		log.Fatal(err)
	}
	// Close() releases allocated resources for PDF-document
	defer pdf.Close()

	// VerifySigns() verifies signs from PDF-document
	signs, err := pdf.VerifySigns()
	if err != nil {
		log.Fatal(err)
	}

	// Print signatures with status in the PDF-document
	for _, sign := range signs {
		fmt.Printf("Signature: %s, Status: %s\n", sign.Name, sign.Status)
	}
}
