package main

import "github.com/aspose-pdf/aspose-pdf-go-cpp"
import "log"

func main() {
	// Open(filename string) opens a PDF-document with filename
	pdf, err := asposepdf.Open("sample.pdf")
	if err != nil {
		log.Fatal(err)
	}
	// RemoveJavaScripts() removes java scripts from PDF-document
	err = pdf.RemoveJavaScripts()
	if err != nil {
		log.Fatal(err)
	}
	// SaveAs(filename string) saves previously opened PDF-document with new filename
	err = pdf.SaveAs("sample_RemoveJavaScripts.pdf")
	if err != nil {
		log.Fatal(err)
	}
	// Close() releases allocated resources for PDF-document
	defer pdf.Close()
}
