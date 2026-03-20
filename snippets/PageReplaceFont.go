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
	// PageReplaceFont(num int32, findFontName, replaceFontName string) replaces font in page
	err = pdf.PageReplaceFont(1, "Times-BoldItalic", "Helvetica-Bold")
	if err != nil {
		log.Fatal(err)
	}
	// SaveAs(filename string) saves previously opened PDF-document with new filename
	err = pdf.SaveAs("sample_page1_ReplaceFont.pdf")
	if err != nil {
		log.Fatal(err)
	}
}
