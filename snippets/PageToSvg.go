package main

import "github.com/aspose-pdf/aspose-pdf-go-cpp"
import "log"

func main() {
	// Open(filename string) opens a PDF-document with filename
	pdf, err := asposepdf.Open("sample.pdf")
	if err != nil {
		log.Fatal(err)
	}
	// Close() releases allocated resources for PDF-document
	defer pdf.Close()
	// PageToSvg(num int32, filename string) saves the specified page as Svg-image file
	err = pdf.PageToSvg(1, "sample_page1.svg")
	if err != nil {
		log.Fatal(err)
	}
}
