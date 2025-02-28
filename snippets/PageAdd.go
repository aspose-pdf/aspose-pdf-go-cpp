package main

import "github.com/aspose-pdf/aspose-pdf-go-cpp"
import "log"

func main() {
	// Open(filename string) opens a PDF-document with filename
	pdf, err := asposepdf.Open("sample.pdf")
	if err != nil {
		log.Fatal(err)
	}
	// PageAdd() adds new page in PDF-document
	err = pdf.PageAdd()
	if err != nil {
		log.Fatal(err)
	}
	// Save() saves previously opened PDF-document
	err = pdf.Save()
	if err != nil {
		log.Fatal(err)
	}
	// Close() releases allocated resources for PDF-document
	defer pdf.Close()
}
