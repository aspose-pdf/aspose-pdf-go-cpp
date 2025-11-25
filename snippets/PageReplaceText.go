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
	// PageReplaceText(num int32, findText, replaceText string) replaces text on page
	err = pdf.PageReplaceText(1, "PDF", "TXT")
	if err != nil {
		log.Fatal(err)
	}
	// SaveAs(filename string) saves previously opened PDF-document with new filename
	err = pdf.SaveAs("sample_page1_ReplaceText.pdf")
	if err != nil {
		log.Fatal(err)
	}
}
