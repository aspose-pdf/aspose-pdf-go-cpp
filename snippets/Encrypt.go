package main

import "github.com/aspose-pdf/aspose-pdf-go-cpp"
import "log"

func main() {
	// New creates a new PDF-document
	pdf, err := asposepdf.New()
	if err != nil {
		log.Fatal(err)
	}
	// Close() releases allocated resources for PDF-document
	defer pdf.Close()
	// Encrypt(userPassword, ownerPassword, permissions, cryptoAlgorithm, usePdf20) encrypts PDF-document
	err = pdf.Encrypt(
		"userpass",
		"ownerpass",
		asposepdf.PrintDocument|asposepdf.ModifyContent|asposepdf.FillForm,
		asposepdf.AESx128,
		true,
	)
	if err != nil {
		log.Fatal(err)
	}
	// SaveAs(filename string) saves previously opened PDF-document with new filename
	err = pdf.SaveAs("sample_with_password.pdf")
	if err != nil {
		log.Fatal(err)
	}
}
