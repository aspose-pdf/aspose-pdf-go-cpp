package main

import "github.com/aspose-pdf/aspose-pdf-go-cpp"
import "log"
import "os"

func main() {
	cert, _ := os.ReadFile("sign.pfx")
	img, _ := os.ReadFile("sign.png")

	// Open(filename string) opens a PDF-document with filename
	pdf, err := asposepdf.Open("sample.pdf")
	if err != nil {
		log.Fatal(err)
	}
	// Close() releases allocated resources for PDF-document
	defer pdf.Close()

	// SignPKCS7 signs a PDF-document using PKCS#7 digital signatures
	err = pdf.SignPKCS7(1, cert, "Pa$$w0rd2023", 100, 100, 70, 100, "Reason", "Contact", "Location", true, img, "sample_SignPKCS7.pdf")
	if err != nil {
		log.Fatal(err)
	}
}
