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
	// PageToDICOM(num int32, resolution_dpi int32, filename string) saves the specified page as DICOM-image file
	err = pdf.PageToDICOM(1, 100, "sample_page1.dcm")
	if err != nil {
		log.Fatal(err)
	}
}
