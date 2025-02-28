package main

import "github.com/aspose-pdf/aspose-pdf-go-cpp"
import "log"

func main() {
	// Open(filename string) opens a PDF-document with filename
	pdf, err := asposepdf.Open("sample.pdf")
	if err != nil {
		log.Fatal(err)
	}
	// SetLicense(filename string) licenses with filename
	err = pdf.SetLicense("Aspose.PDF.GoViaCPP.lic")
	if err != nil {
		log.Fatal(err)
	}
	// Working with PDF-document
	// ...
	// Close() releases allocated resources for PDF-document
	defer pdf.Close()
}
