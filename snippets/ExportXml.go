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
	// ExportXml(filename string) exports from previously opened PDF-document with AcroForm to XML-document with filename
	err = pdf.ExportXml("sample.xml")
	if err != nil {
		log.Fatal(err)
	}
}
