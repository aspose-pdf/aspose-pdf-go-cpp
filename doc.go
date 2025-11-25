// Aspose.PDF for Go via C++
//
//	The package asposepdf is a powerful toolkit that allows developers to manipulate PDF files directly and helps do various tasks for PDF.
//
// Features
//
//	PDF Processing
//	 Main core operation: New, Open, Save, SaveAs, Close, SetLicense, Append, AppendPages, MergeDocuments, SplitDocument, Split, SplitAtPage, SplitAt
//	 Other core operation: WordCount, CharacterCount, Bytes
//	 Page main core operation: Add, Insert, Delete, Count
//	 Page other core operation: WordCount, CharacterCount, IsBlank
//	 Organize: Optimize, OptimizeResource, OptimizeFileSize, Grayscale, Rotate, SetBackground, Repair, Flatten, AddPageNum, AddHeader, AddFooter, AddWatermark
//	 Page organize: Rotate, SetSize, Grayscale, AddPageNum, AddText, AddHeader, AddFooter, AddWatermark
//	 Remove operation: RemoveAnnotations, RemoveAttachments, RemoveBlankPages, RemoveBookmarks, RemoveHiddenText, RemoveImages, RemoveTables, RemoveJavaScripts, RemoveWatermarks
//	 Page remove operation: PageRemoveAnnotations, PageRemoveHiddenText, PageRemoveImages, PageRemoveTables, PageRemoveWatermarks
//	 Font embedding: EmbedFonts and UnembedFonts
//	 Others: Get contents as plain text
//
//	PDF converting and saving
//	 Microsoft Office: DOC, DOCX, XLSX, PPTX, DOCX with Enhanced Recognition Mode (fully editable tables and paragraphs)
//	 Images: JPEG, PNG, BMP, TIFF
//	 Others: EPUB, DICOM, SVG, SVG(ZIP), XPS, TEX, TXT, MD, N-UP PDF, BOOKLET PDF
//	 Export with AcroForm: FDF, XFDF, XML
//
//	Metadata
//	 Product Info: JSON with product name, version, release date, and license status
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
//		// Release allocated resources
//		defer pdf.Close()
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
//	}
//
// Testing
//
//	The test run from the root package folder:
//	  go test -v
//
// License
//
//   - The Go source code is licensed under the MIT License.
//
//   - The shared native libraries (DLL, SO, DYLIB) are proprietary and require a commercial license.
//
//     Evaluation version limitations:
//
//   - Documents created with an evaluation watermark.
//
//   - Limit on number of pages processed (first 4 pages).
//
//     For production use, a commercial license is required.
//
// Resources
//
//	Aspose home: https://www.aspose.com
//	Product Page: https://products.aspose.com/pdf/go-cpp/
//	Docs: https://docs.aspose.com/pdf/go-cpp/
//	Demos: https://products.aspose.app/pdf/family
//	API Reference: https://reference.aspose.com/pdf/go-cpp/
//	Examples: https://github.com/aspose-pdf/aspose-pdf-go-cpp/
//	Blog: https://blog.aspose.com/category/pdf/
//	Free Support: https://forum.aspose.com/c/pdf
//	Temporary License: https://purchase.aspose.com/temporary-license
//	GitHub.com: https://github.com/aspose-pdf/aspose-pdf-go-cpp
//	pkg.go.dev: https://pkg.go.dev/github.com/aspose-pdf/aspose-pdf-go-cpp
package asposepdf
