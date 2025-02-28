// Aspose.PDF for Go via C++
//
//	The package asposepdf is a powerful toolkit that allows developers to manipulate PDF files directly and helps do various tasks for PDF.
//
// Features
//
//	PDF Processing
//	 Core operation: New, Open, Save, SaveAs, Close, SetLicense, WordCount, CharacterCount
//	 Page core operation: Add, Insert, Delete, Count, WordCount, CharacterCount, IsBlank
//	 Organize: Optimize, OptimizeResource, Grayscale, Rotate, SetBackground, Repair
//	 Page organize: Rotate, SetSize, Grayscale, AddText
//	 Others: Get contents as plain text
//
//	PDF converting and saving
//	 Microsoft Office: DOC, DOCX, XLSX, PPTX
//	 Images: JPEG, PNG, BMP, TIFF
//	 Others: EPUB, DICOM, SVG, XPS, TEX, TXT
//
// Platforms
//
//	Linux x64, macOS x86_64, macOS arm64 and Windows x64, with using cgo.
//
//	The platform-specific version of the dynamic library from the 'lib'-folder in the package's root directory is required for distributing the resulting application:
//	 libAsposePDFforGo_linux_amd64.so for Linux x64 platform
//	 libAsposePDFforGo_darwin_arm64.dylib for macOS arm64 platform
//	 libAsposePDFforGo_darwin_amd64.dylib for macOS x86_64 platform
//	 AsposePDFforGo_windows_amd64.dll for Windows x64 platform
//
//	Windows x64 platform requires MinGW-W64 installed.
//
// Installation
//
//	This package includes a large file which is stored as a bzip2 archive.
//
//	Add the asposepdf package to Your Project:
//		  go get github.com/aspose-pdf/aspose-pdf-go-cpp@latest
//
//	Generate the large file on macOS and linux
//
//		List the folders of the github.com/aspose-pdf within the Go module cache:
//		  ls $(go env GOMODCACHE)/github.com/aspose-pdf/
//
//		Change curent folder to the specific version folder of the package obtained in the previous step.
//		Replace `@vx.x.x` with the actual package version:
//		  cd $(go env GOMODCACHE)/github.com/aspose-pdf/aspose-pdf-go-cpp@vx.x.x
//
//		Run go generate with superuser privileges:
//		  sudo go generate
//
//	Generate the large file on Windows:
//
//		List the folders of the github.com/aspose-pdf within the Go module cache:
//		  for /f "delims=" %G in ('go env GOMODCACHE') do for /d %a in ("%G\github.com\aspose-pdf\*") do echo %~fa
//
//		Change curent folder to the specific version folder of the package obtained in the previous step:
//		  cd <specific version folder of the package>
//
//		Run go generate:
//		  go generate
//
//		Add specific version folder of the package to the %PATH% environment variable:
//		  setx PATH "%PATH%;<specific version folder of the package>\lib\"
//
// Quick Start
//
//	All code snippets are contained in the snippet folder on https://github.com/aspose-pdf/aspose-pdf-go-cpp.
//
//	Hello World! example:
//
//	package main
//
//	import "github.com/aspose-pdf/aspose-pdf-go-cpp"
//	import "log"
//
//	func main() {
//		// Create new PDF-document
//		pdf, err := asposepdf.New()
//		if err != nil {
//			log.Fatal(err)
//		}
//		// Add new page
//		err = pdf.PageAdd()
//		if err != nil {
//			log.Fatal(err)
//		}
//		// Set page size A4
//		err = pdf.PageSetSize(1, asposepdf.PageSizeA4)
//		if err != nil {
//			log.Fatal(err)
//		}
//		// Add text on first page
//		err = pdf.PageAddText(1, "Hello World!")
//		if err != nil {
//			log.Fatal(err)
//		}
//		// Save PDF-document with "hello.pdf" name
//		err = pdf.SaveAs("hello.pdf")
//		if err != nil {
//			log.Fatal(err)
//		}
//		// Release allocated resources
//		defer pdf.Close()
//	}
//
// Testing
//
//	The test run from the root package folder:
//	  go test -v
//
// Aspose home
//
//	https://www.aspose.com
//
// GitHub.com
//
//	https://github.com/aspose-pdf/aspose-pdf-go-cpp
package asposepdf
