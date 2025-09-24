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
	// About() returns metadata information about the Aspose.PDF for Go via C++
	info, err := pdf.About()
	if err != nil {
		log.Fatal(err)
	}
	// Print
	fmt.Println("Aspose.PDF Product Info:")
	fmt.Println("  Product:     ", info.Product)
	fmt.Println("  Family:      ", info.Family)
	fmt.Println("  Version:     ", info.Version)
	fmt.Println("  ReleaseDate: ", info.ReleaseDate)
	fmt.Println("  Producer:    ", info.Producer)
	fmt.Println("  IsLicensed:  ", info.IsLicensed)
	// Close() releases allocated resources for PDF-document
	defer pdf.Close()
}
