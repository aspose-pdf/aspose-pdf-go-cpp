package main

import (
	"github.com/aspose-pdf/aspose-pdf-go-cpp"
	"log"
	"os"
)

func main() {
	// Open(filename string) opens a PDF-document with filename
	pdf, err := asposepdf.Open("sample.pdf")
	if err != nil {
		log.Fatal(err)
	}
	// Close() releases allocated resources for PDF-document
	defer pdf.Close()

	// Load stamp image data
	imageBytes, err := os.ReadFile("stamp.png")
	if err != nil {
		log.Fatal(err)
	}

	// Add stamp to the PDF-document
	err = pdf.AddStamp("Stamp", imageBytes, true, 100, 100, 50, 50, asposepdf.RotationOn180, 0.5)
	if err != nil {
		log.Fatal(err)
	}

	// Save the modified PDF-document
	err = pdf.SaveAs("sample_AddStamp.pdf")
	if err != nil {
		log.Fatal(err)
	}
}
