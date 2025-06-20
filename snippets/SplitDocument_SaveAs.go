package main

import (
	"fmt"
	"log"
	"github.com/aspose-pdf/aspose-pdf-go-cpp"
)

func main() {
	// Open(filename string) opens a PDF-document with filename
	pdf_split, err := asposepdf.Open("sample.pdf")
	if err != nil {
		log.Fatal(err)
	}
	// Close() releases allocated resources for PDF-document
	defer pdf_split.Close()

	// SplitDocument(document *Document, pagerange string) creates multiple new PDF-documents
	pdfs, err := asposepdf.SplitDocument(pdf_split, "1-2;3;4-")
	if err != nil {
		log.Fatal(err)
	}

	// Save each split PDF-document as a separate file
	for i, pdf := range pdfs {
		defer pdf.Close()
		filename := fmt.Sprintf("sample_SplitDocument_part%d.pdf", i+1)
		// SaveAs(filename string) saves previously opened PDF-document with new filename
		err := pdf.SaveAs(filename)
		if err != nil {
			log.Fatal(err)
		}
	}
}
