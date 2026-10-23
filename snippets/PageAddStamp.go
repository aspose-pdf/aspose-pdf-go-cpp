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

	// PageAddStamp adds stamp on page.
	err = pdf.PageAddStamp(1, "PageStamp", imageBytes, true, 100, 100, 50, 50, asposepdf.RotationOn90, 0.5)
	if err != nil {
		log.Fatal(err)
	}

	// SaveAs(filename string) saves previously opened PDF-document with new filename
	err = pdf.SaveAs("sample_PageAddStamp.pdf")
	if err != nil {
		log.Fatal(err)
	}
}
