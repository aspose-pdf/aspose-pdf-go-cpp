package main

import "github.com/aspose-pdf/aspose-pdf-go-cpp"
import "log"
import "fmt"

func main() {
	// Open(filename string) opens a PDF-document with filename
	pdf, err := asposepdf.Open("sample.pdf")
	if err != nil {
		log.Fatal(err)
	}
	// Close() releases allocated resources for PDF-document
	defer pdf.Close()
	// IsPdfUaCompliant() gets PDF/UA compliance status of PDF-document
	isPdfua, _ := pdf.IsPdfUaCompliant()
	if isPdfua {
		fmt.Println("IsPdfUaCompliant() is true")
	} else {
		fmt.Println("IsPdfUaCompliant() is false")
	}
}
