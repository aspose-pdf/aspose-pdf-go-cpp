package main

import (
	"fmt"
	"github.com/aspose-pdf/aspose-pdf-go-cpp"
	"log"
)

func main() {
	// Open(filename string) opens a PDF-document with filename
	pdf, err := asposepdf.Open("sample.pdf")
	if err != nil {
		log.Fatal(err)
	}
	// Close() releases allocated resources for PDF-document
	defer pdf.Close()

	// Convert(pdfFormat PdfFormat, action ConvertErrorAction) converts PDF-document to PDF/A
	ok, logStr, err := pdf.Convert(asposepdf.PDF_A_2A, asposepdf.Delete)
	if err != nil {
		log.Fatal("Convert PDF/A error:", err)
	}

	// Print conversion result and full log
	fmt.Printf("Convert PDF/A result: %v\n", ok)
	fmt.Printf("Convert PDF/A log:\n%s\n", logStr)

	// Save converted PDF-document
	err = pdf.SaveAs("sample_Convert.pdf")
	if err != nil {
		log.Fatal(err)
	}
}
