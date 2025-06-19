package main

import (
	"github.com/aspose-pdf/aspose-pdf-go-cpp"
	"log"
	"os"
)

func main() {
	// New creates a new PDF-document
	pdf, err := asposepdf.New()
	if err != nil {
		log.Fatal(err)
	}
	defer pdf.Close()

	// Bytes returns the contents of the PDF-document as a byte slice
	bytes, err := pdf.Bytes()
	if err != nil {
		log.Fatal(err)
	}

	// Save the byte slice to a file.
	err = os.WriteFile("sample_Bytes.pdf", bytes, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
