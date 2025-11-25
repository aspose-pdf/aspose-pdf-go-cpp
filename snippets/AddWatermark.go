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
	// AddWatermark(text string, fontName string, fontSize float64, foregroundColor string, xPosition int32, yPosition int32, rotation int32, isBackground bool, opacity float64) adds watermark to PDF-document
	err = pdf.AddWatermark("Watermark", "Arial", 16, "#010101", 100, 100, 45, true, 0.5)
	if err != nil {
		log.Fatal(err)
	}
	// SaveAs(filename string) saves previously opened PDF-document with new filename
	err = pdf.SaveAs("sample_AddWatermark.pdf")
	if err != nil {
		log.Fatal(err)
	}
}
