package main

import "github.com/aspose-pdf/aspose-pdf-go-cpp"
import "log"

func main() {
	// Open(filename string) opens a PDF-document with filename
	pdf, err := asposepdf.Open("sample.pdf")
	if err != nil {
		log.Fatal(err)
	}
	// SaveTiff(filename string) saves previously opened PDF-document as Tiff-document with filename
	err = pdf.SaveTiff("sample.tiff")
	if err != nil {
		log.Fatal(err)
	}
	// Close() releases allocated resources for PDF-document
	defer pdf.Close()
}
