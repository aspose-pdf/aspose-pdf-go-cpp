package main

import "github.com/aspose-pdf/aspose-pdf-go-cpp"
import "log"

func main() {
	// Open(filename string) opens a PDF-document with filename
	pdf, err := asposepdf.Open("sample.pdf")
	if err != nil {
		log.Fatal(err)
	}
	// ExportXfdf(filename string) exports from previously opened PDF-document with AcroForm to XFDF-document with filename
	err = pdf.ExportXfdf("sample.xfdf")
	if err != nil {
		log.Fatal(err)
	}
	// Close() releases allocated resources for PDF-document
	defer pdf.Close()
}
