# Aspose.PDF for Go via C++

The package asposepdf is a powerful toolkit that allows developers to manipulate PDF files directly and helps do various tasks for PDF.
Contains unique features for converting PDF to other formats.

## Features

### PDF Processing

- **Main core operation:** New, Open, Save, SaveAs, Close, SetLicense, Append, AppendPages, MergeDocuments, SplitDocument, Split, SplitAtPage, SplitAt
- **Other core operation:** WordCount, CharacterCount, Bytes
- **Page main core operation:** Add, Insert, Delete, Count
- **Page other core operation:** WordCount, CharacterCount, IsBlank
- **Organize:** Optimize, OptimizeResource, Grayscale, Rotate, SetBackground, Repair, Flatten
- **Page organize:** Rotate, SetSize, Grayscale, AddText
- **Remove operation:** RemoveAnnotations, RemoveAttachments, RemoveBlankPages, RemoveBookmarks, RemoveHiddenText, RemoveImages, RemoveTables, RemoveJavaScripts
- **Page remove operation:** PageRemoveAnnotations, PageRemoveHiddenText, PageRemoveImages, PageRemoveTables
- **Others:** Get contents as plain text

### PDF converting and saving

- **Microsoft Office:** DOC, DOCX, XLSX, PPTX, DOCX with Enhanced Recognition Mode (fully editable tables and paragraphs)
- **Images:** JPEG, PNG, BMP, TIFF
- **Others:** EPUB, DICOM, SVG, SVG(ZIP), XPS, TEX, TXT, MD, N-UP PDF, BOOKLET PDF
- **Export with AcroForm:** FDF, XFDF, XML

### Metadata

- **Product Info:** JSON with product name, version, release date, and license status

## Platforms

Implemented support for Linux x64, macOS x86_64, macOS arm64 and Windows x64 platforms. Used [cgo](https://go.dev/wiki/cgo).

The platform-specific version of the dynamic library from the 'lib'-folder in the package's root directory is required for distributing the resulting application:
- *libAsposePDFforGo_linux_amd64.so* for Linux x64 platform
- *libAsposePDFforGo_darwin_arm64.dylib* for macOS arm64 platform
- *libAsposePDFforGo_darwin_amd64.dylib* for macOS x86_64 platform
- *AsposePDFforGo_windows_amd64.dll* for Windows x64 platform.

Windows x64 platform requires [MinGW-W64](https://www.mingw-w64.org/) installed.

## Installation

This package includes a large file which is stored as a bzip2 archive.

1. Add the asposepdf package to Your Project:
    ```sh
    go get github.com/aspose-pdf/aspose-pdf-go-cpp@latest
    ```

2. Generate the large file:

 - **macOS and linux**

  1. Open Terminal

  2. List the folders of the github.com/aspose-pdf within the Go module cache:

        ```sh
        ls $(go env GOMODCACHE)/github.com/aspose-pdf/
        ```

  3. Change curent folder to the specific version folder of the package obtained in the previous step:

      ```sh
      cd $(go env GOMODCACHE)/github.com/aspose-pdf/aspose-pdf-go-cpp@vx.x.x
      ```
      Replace `@vx.x.x` with the actual package version.


  4. Run go generate with superuser privileges:

      ```sh
      sudo go generate
      ```

 - **Windows**

  1. Open Command Prompt
  
  2. List the folders of the github.com/aspose-pdf within the Go module cache:

      ```cmd
      for /f "delims=" %G in ('go env GOMODCACHE') do for /d %a in ("%G\github.com\aspose-pdf\*") do echo %~fa
      ```

  3. Change curent folder to the specific version folder of the package obtained in the previous step:

      ```cmd
      cd <specific version folder of the package>
      ```

  4. Run go generate:

      ```cmd
      go generate
      ```

  5. Add specific version folder of the package to the %PATH% environment variable:

      ```cmd
      setx PATH "%PATH%;<specific version folder of the package>\lib\"
      ```
      Replace `<specific version folder of the package>` with the actual path obtained from step 2.


## Quick Start
All code snippets are contained in the [snippet](https://github.com/aspose-pdf/aspose-pdf-go-cpp).

### Hello World!

```go
package main

import "github.com/aspose-pdf/aspose-pdf-go-cpp"
import "log"

func main() {
	// Create new PDF-document
	pdf, err := asposepdf.New()
	if err != nil {
		log.Fatal(err)
	}
	// Add new page
	err = pdf.PageAdd()
	if err != nil {
		log.Fatal(err)
	}
	// Set page size A4
	err = pdf.PageSetSize(1, asposepdf.PageSizeA4)
	if err != nil {
		log.Fatal(err)
	}
	// Add text on first page
	err = pdf.PageAddText(1, "Hello World!")
	if err != nil {
		log.Fatal(err)
	}
	// Save PDF-document with "hello.pdf" name
	err = pdf.SaveAs("hello.pdf")
	if err != nil {
		log.Fatal(err)
	}
	// Release allocated resources
	defer pdf.Close()
}
```

### Save PDF as Office Formats

One of the most popular features of Aspose.PDF for Go via C++ is to convert PDF documents to other formats without needing to understand the underlying structure of the resultant format.

Give the following snippet a try with your samples:

```go
package main

import "github.com/aspose-pdf/aspose-pdf-go-cpp"
import "log"

func main() {
	// Open(filename string) opens a PDF-document with filename
	pdf, err := asposepdf.Open("sample.pdf")
	if err != nil {
		log.Fatal(err)
	}
	// SaveDocX(filename string) saves previously opened PDF-document as DocX-document with filename
	err = pdf.SaveDocX("sample.docx")
	if err != nil {
		log.Fatal(err)
	}
	// Close() releases allocated resources for PDF-document
	defer pdf.Close()
}
```
### Extract Text From Whole PDF

```go
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
	// ExtractText() returns PDF-document contents as plain text
	txt, err := pdf.ExtractText()
	if err != nil {
		log.Fatal(err)
	}
	// Print
	fmt.Println("Extracted text:\n", txt)
	// Close() releases allocated resources for PDF-document
	defer pdf.Close()
}
```

## Testing

The test run from the root package folder:

```sh
go test -v
```

## License

- The **Go source code** is licensed under the [MIT License](LICENSE).
- The **shared library (`AsposePDFforGo_windows_amd64.dll`, `libAsposePDFforGo_linux_amd64.so`, `libAsposePDFforGo_darwin_amd64.dylib`, `libAsposePDFforGo_darwin_arm64.dylib`)** is proprietary and requires a commercial license.
  To use the full functionality, you must obtain a license.

### Evaluation version

You can use Aspose.PDF for Go via C++ free of cost for evaluation.The evaluation version provides almost all functionality of the product with certain limitations. The same evaluation version becomes licensed when you purchase a license and add a couple of lines of code to apply the license.

>If you want to test Aspose.PDF for Go without the evaluation version limitations, you can also request a 30-day Temporary License. Please refer to [How to get a Temporary License?](https://purchase.aspose.com/temporary-license)

### Limitation of an evaluation version

We want our customers to test our components thoroughly before buying so the evaluation version allows you to use it as you would normally.

- **Documents created with an evaluation watermark.** The evaluation version of Aspose.PDF for Go provides full product functionality, but all pages in the generated files are watermarked with the text "Evaluation Only. Created with Aspose.PDF. Copyright 2002-2025 Aspose Pty Ltd." at the top.
- **Limit the number of pages that can be processed.** In the evaluation version, you can only process the first four pages of a document.

### Use in production

A commercial license key is required in a production environment. Please contact us to <a href="https://purchase.aspose.com/buy">purchase a commercial license</a>.

### Apply license

Applying a license to enable full functionality of the Aspose.PDF for Go using a license file (Aspose.PDF.GoViaCPP.lic).

```go

    package main

    import "github.com/aspose-pdf/aspose-pdf-go-cpp"
    import "log"

    func main() {
        // Open(filename string) opens a PDF-document with filename
        pdf, err := asposepdf.Open("sample.pdf")
        if err != nil {
            log.Fatal(err)
        }
        // SetLicense(filename string) licenses with filename
        err = pdf.SetLicense("Aspose.PDF.GoViaCPP.lic")
        if err != nil {
            log.Fatal(err)
        }
        // Working with PDF-document
        // ...
        // Close() releases allocated resources for PDF-document
        defer pdf.Close()
    }
```

[Home](https://www.aspose.com/) | [Product Page](https://products.aspose.com/pdf/go-cpp/) | [Docs](https://docs.aspose.com/pdf/go-cpp/) | [Demos](https://products.aspose.app/pdf/family) | [API Reference](https://reference.aspose.com/pdf/go-cpp/) | [Examples](https://github.com/aspose-pdf/aspose-pdf-go-cpp/) | [Blog](https://blog.aspose.com/category/pdf/) | [Search](https://search.aspose.com/) | [Free Support](https://forum.aspose.com/c/pdf) | [Temporary License](https://purchase.aspose.com/temporary-license) | [GitHub.com](https://github.com/aspose-pdf/aspose-pdf-go-cpp) | [pkg.go.dev](https://pkg.go.dev/github.com/aspose-pdf/aspose-pdf-go-cpp)