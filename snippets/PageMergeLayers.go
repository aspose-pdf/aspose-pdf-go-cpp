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
	// PageMergeLayers(num int32, newLayerName string) merges all layers on the page into a single layer with the specified new layer name
	err = pdf.PageMergeLayers(1, "newLayerName")
	if err != nil {
		log.Fatal(err)
	}
	// SaveAs(filename string) saves previously opened PDF-document with new filename
	err = pdf.SaveAs("sample_PageMergeLayers.pdf")
	if err != nil {
		log.Fatal(err)
	}
}
