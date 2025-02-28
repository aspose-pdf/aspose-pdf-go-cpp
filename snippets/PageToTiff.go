package main

import "github.com/aspose-pdf/aspose-pdf-go-cpp"
import "log"

func main() {
	// Open(filename string) opens a PDF-document with filename
	pdf, err := asposepdf.Open("sample.pdf")
	if err != nil {
		log.Fatal(err)
	}
	// PageToTiff(num int32, resolution_dpi int32, filename string) saves the specified page as Tiff-image file
	err = pdf.PageToTiff(1, 100, "sample_page1.tiff")
	if err != nil {
		log.Fatal(err)
	}
	// Close() releases allocated resources for PDF-document
	defer pdf.Close()
}
