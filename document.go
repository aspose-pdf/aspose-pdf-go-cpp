package asposepdf

/*
#cgo CFLAGS: -g -Wall
#include <stdlib.h>
#include "extern_c.h"
*/
import "C"

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"unsafe"
)

// Document represents a PDF-document.
type Document struct {
	pdf unsafe.Pointer
}

// New creates a new PDF-document.
//
// Example:
//
//		pdf_new, err := New()
//		if err != nil {
//			fmt.Errorf("New(): %v", err)
//		} else {
//	 		// working with new PDF-document
//		}
func New() (*Document, error) {
	var err *C.char
	doc := C.PDFDocument_New(&err)
	err_str := C.GoString(err)
	C.c_free_string(err)
	if doc != nil {
		return &Document{doc}, nil
	} else {
		return &Document{nil}, errors.New(err_str)
	}
}

// Open opens a PDF-document with filename.
//
// Example:
//
//		pdf, err := Open("example.pdf")
//		if err != nil {
//			fmt.Errorf("Open(): %v", err)
//		} else {
//	 		// working with open PDF-document
//		}
func Open(filename string) (*Document, error) {
	var err *C.char
	_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(_filename))
	doc := C.PDFDocument_Open(_filename, &err)
	err_str := C.GoString(err)
	C.c_free_string(err)
	if doc != nil {
		return &Document{doc}, nil
	} else {
		return &Document{nil}, errors.New(err_str)
	}
}

// MergeDocuments creates a new PDF-document by merging the provided documents.
//
// Example:
//
//		pdf_merged, err := MergeDocuments([]*Document{pdf1, pdf2})
//		if err != nil {
//			fmt.Errorf("MergeDocuments(): %v", err)
//		} else {
//	 		// working with new merged PDF-document
//		}
func MergeDocuments(documents []*Document) (*Document, error) {
	if len(documents) == 0 {
		return nil, errors.New("MergeDocuments(): no documents to merge")
	}

	// Create a new empty PDF document
	merged, err := New()
	if err != nil {
		return nil, fmt.Errorf("MergeDocuments(): failed to create new document: %w", err)
	}

	// Append each input document to the merged document
	for i, document := range documents {
		if document == nil || document.pdf == nil {
			return nil, fmt.Errorf("MergeDocuments(): document at index %d is nil or invalid", i)
		}
		if err := merged.Append(document); err != nil {
			return nil, fmt.Errorf("MergeDocuments(): failed to append document at index %d: %w", i, err)
		}
	}

	return merged, nil
}

// splitDocument is a helper used by Split and SplitDocument.
// Splits the source document into multiple documents based on the page range string.
//
// Each part of the pagerange string (separated by `;`) defines the page range for a new PDF-document.
func splitDocument(document *Document, pagerange string) ([]*Document, error) {
	if document == nil || document.pdf == nil {
		return nil, errors.New("splitDocument: source document is nil or invalid")
	}

	if pagerange == "" {
		return nil, errors.New("splitDocument: empty page range string")
	}

	parts := strings.Split(pagerange, ";")
	result := make([]*Document, 0, len(parts))

	for i, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			return nil, fmt.Errorf("splitDocument: empty page range at index %d", i)
		}

		newdoc, err := New()
		if err != nil {
			return nil, fmt.Errorf("splitDocument: failed to create new document for range %q: %w", part, err)
		}

		if err := newdoc.AppendPages(document, part); err != nil {
			newdoc.Close()
			return nil, fmt.Errorf("splitDocument: failed to append pages %q: %w", part, err)
		}

		result = append(result, newdoc)
	}

	return result, nil
}

// SplitDocument creates multiple new PDF-documents by extracting pages from the source PDF-document.
//
// Each part of the pagerange string (separated by `;`) defines the page range for a new PDF-document.
//
// Example:
//
//	pdfs, err := asposepdf.SplitDocument(pdf_source, "1-2;3;4-")
//	// pdfs[0] will contain pages 1-2, pdfs[1] page 3, pdfs[2] pages 4 to end.
func SplitDocument(document *Document, pagerange string) ([]*Document, error) {
	return splitDocument(document, pagerange)
}

// splitAtPage is an internal helper used by SplitAtPage and SplitAt.
// Splits the document at the specified page into two new PDF-documents:
// [1..=page] and [page+1..end].
func splitAtPage(document *Document, page int) (*Document, *Document, error) {
	if document == nil || document.pdf == nil {
		return nil, nil, errors.New("splitAtPage: source document is nil or invalid")
	}

	pageCount, err := document.PageCount()
	if err != nil {
		return nil, nil, fmt.Errorf("splitAtPage: failed to get page count: %w", err)
	}

	if page < 1 || page >= int(pageCount) {
		return nil, nil, fmt.Errorf("splitAtPage: page %d is out of valid range (1-%d exclusive)", page, pageCount)
	}

	// Create first document for pages 1 to page
	left, err := New()
	if err != nil {
		return nil, nil, fmt.Errorf("splitAtPage: failed to create first document: %w", err)
	}

	if err := left.AppendPages(document, fmt.Sprintf("1-%d", page)); err != nil {
		left.Close()
		return nil, nil, fmt.Errorf("splitAtPage: failed to append left pages: %w", err)
	}

	// Create second document for pages page+1 to end
	right, err := New()
	if err != nil {
		right.Close()
		return nil, nil, fmt.Errorf("splitAtPage: failed to create second document: %w", err)
	}

	if err := right.AppendPages(document, fmt.Sprintf("%d-", page+1)); err != nil {
		left.Close()
		right.Close()
		return nil, nil, fmt.Errorf("splitAtPage: failed to append right pages: %w", err)
	}

	return left, right, nil
}

// SplitAtPage splits the PDF-document into two new PDF-documents.
// The first document includes pages 1 to 'page' (inclusive).
// The second document includes pages from 'page+1' to the end.
//
// Example:
//
//	left, right, err := SplitAtPage(source, 3)
//	// 'left' contains pages 1-3, 'right' contains pages 4 to end
func SplitAtPage(document *Document, page int) (*Document, *Document, error) {
	return splitAtPage(document, page)
}

// Close releases allocated resources for PDF-document.
//
// Example:
//
//	defer pdf.Close()
func (document *Document) Close() error {
	var err *C.char
	C.PDFDocument_Release(document.pdf, &err)
	document.pdf = nil
	err_str := C.GoString(err)
	C.c_free_string(err)
	if err_str != ERR_OK {
		return errors.New(err_str)
	} else {
		return nil
	}
}

// About returns metadata information about the Aspose.PDF for Go via C++.
//
// The metadata is returned as a ProductInfo struct, deserialized from a JSON string.
// It includes product name, version, release date, licensing status, and related details.
//
// See also: product_info.go
//
// Example:
//
//	info, err := pdf.About()
func (document *Document) About() (*ProductInfo, error) {
	var err *C.char
	jsonStr := C.PDFDocument_About(document.pdf, &err)
	err_str := C.GoString(err)
	C.c_free_string(err)
	if err_str != ERR_OK {
		return nil, errors.New(err_str)
	}
	defer C.c_free_string(jsonStr)
	goJSON := C.GoString(jsonStr)
	var info ProductInfo
	if e := json.Unmarshal([]byte(goJSON), &info); e != nil {
		return nil, e
	}
	return &info, nil
}

// Save saves previously opened PDF-document.
//
// Example:
//
//	err := pdf.Save()
func (document *Document) Save() error {
	var err *C.char
	C.PDFDocument_Save(document.pdf, &err)
	err_str := C.GoString(err)
	C.c_free_string(err)
	if err_str != ERR_OK {
		return errors.New(err_str)
	} else {
		return nil
	}
}

// SaveAs saves previously opened PDF-document with new filename.
//
// Example:
//
//	err := pdf.SaveAs("new_filename.pdf")
func (document *Document) SaveAs(filename string) error {
	var err *C.char
	_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(_filename))
	C.PDFDocument_Save_As(document.pdf, _filename, &err)
	err_str := C.GoString(err)
	C.c_free_string(err)
	if err_str != ERR_OK {
		return errors.New(err_str)
	} else {
		return nil
	}
}

// SetLicense licenses with filename.
//
// Example:
//
//	err := pdf.SetLicense("Aspose.PDF.Go.lic")
func (document *Document) SetLicense(filename string) error {
	var err *C.char
	_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(_filename))
	C.PDFDocument_set_License(document.pdf, _filename, &err)
	err_str := C.GoString(err)
	C.c_free_string(err)
	if err_str != ERR_OK {
		return errors.New(err_str)
	} else {
		return nil
	}
}

// Split creates multiple new PDF-documents by extracting pages from the current PDF-document.
//
// Each part of the pagerange string (separated by `;`) defines the page range for a new PDF-document.
//
// Example:
//
//	pdfs, err := pdf.Split("1-2;3;4-")
//	// pdfs[0] will contain pages 1-2, pdfs[1] page 3, pdfs[2] pages 4 to end.
func (document *Document) Split(pagerange string) ([]*Document, error) {
	return splitDocument(document, pagerange)
}

// SplitAt splits the current PDF-document into two new PDF-documents.
// The first document includes pages 1 to 'page' (inclusive).
// The second document includes pages from 'page+1' to the end.
//
// Example:
//
//	left, right, err := SplitAtPage(source, 3)
//	// 'left' contains pages 1-3, 'right' contains pages 4 to end
func (document *Document) SplitAt(page int) (*Document, *Document, error) {
	return splitAtPage(document, page)
}

// ExtractText returns PDF-document contents as plain text.
//
// Example:
//
//	txt, err := pdf.ExtractText()
func (document *Document) ExtractText() (string, error) {
	var err *C.char
	txt := C.PDFDocument_ExtractText(document.pdf, &err)
	err_str := C.GoString(err)
	C.c_free_string(err)
	if err_str != ERR_OK {
		return "", errors.New(err_str)
	} else {
		return C.GoString(txt), nil
	}
}

// Optimize optimizes PDF-document content.
//
// Example:
//
//	err := pdf.Optimize()
func (document *Document) Optimize() error {
	var err *C.char
	C.PDFDocument_Optimize(document.pdf, &err)
	err_str := C.GoString(err)
	C.c_free_string(err)
	if err_str != ERR_OK {
		return errors.New(err_str)
	} else {
		return nil
	}
}

// OptimizeResource optimizes resources of PDF-document.
//
// Example:
//
//	err := pdf.OptimizeResource()
func (document *Document) OptimizeResource() error {
	var err *C.char
	C.PDFDocument_OptimizeResource(document.pdf, &err)
	err_str := C.GoString(err)
	C.c_free_string(err)
	if err_str != ERR_OK {
		return errors.New(err_str)
	} else {
		return nil
	}
}

// Repair repaires PDF-document.
//
// Example:
//
//	err := pdf.Repair()
func (document *Document) Repair() error {
	var err *C.char
	C.PDFDocument_Repair(document.pdf, &err)
	err_str := C.GoString(err)
	C.c_free_string(err)
	if err_str != ERR_OK {
		return errors.New(err_str)
	} else {
		return nil
	}
}

// Grayscale converts PDF-document to black and white.
//
// Example:
//
//	err := pdf.Grayscale()
func (document *Document) Grayscale() error {
	var err *C.char
	C.PDFDocument_Grayscale(document.pdf, &err)
	err_str := C.GoString(err)
	C.c_free_string(err)
	if err_str != ERR_OK {
		return errors.New(err_str)
	} else {
		return nil
	}
}

// Flatten flattens PDF-document.
//
// Example:
//
//	err := pdf.Flatten()
func (document *Document) Flatten() error {
	var err *C.char
	C.PDFDocument_Flatten(document.pdf, &err)
	err_str := C.GoString(err)
	C.c_free_string(err)
	if err_str != ERR_OK {
		return errors.New(err_str)
	} else {
		return nil
	}
}

// RemoveAnnotations removes annotations from PDF-document.
//
// Example:
//
//	err := pdf.RemoveAnnotations()
func (document *Document) RemoveAnnotations() error {
	var err *C.char
	C.PDFDocument_RemoveAnnotations(document.pdf, &err)
	err_str := C.GoString(err)
	C.c_free_string(err)
	if err_str != ERR_OK {
		return errors.New(err_str)
	} else {
		return nil
	}
}

// RemoveAttachments removes attachments from PDF-document.
//
// Example:
//
//	err := pdf.RemoveAttachments()
func (document *Document) RemoveAttachments() error {
	var err *C.char
	C.PDFDocument_RemoveAttachments(document.pdf, &err)
	err_str := C.GoString(err)
	C.c_free_string(err)
	if err_str != ERR_OK {
		return errors.New(err_str)
	} else {
		return nil
	}
}

// RemoveBlankPages removes blank pages from PDF-document.
//
// Example:
//
//	err := pdf.RemoveBlankPages()
func (document *Document) RemoveBlankPages() error {
	var err *C.char
	C.PDFDocument_RemoveBlankPages(document.pdf, &err)
	err_str := C.GoString(err)
	C.c_free_string(err)
	if err_str != ERR_OK {
		return errors.New(err_str)
	} else {
		return nil
	}
}

// RemoveBookmarks removes bookmarks from PDF-document.
//
// Example:
//
//	err := pdf.RemoveBookmarks()
func (document *Document) RemoveBookmarks() error {
	var err *C.char
	C.PDFDocument_RemoveBookmarks(document.pdf, &err)
	err_str := C.GoString(err)
	C.c_free_string(err)
	if err_str != ERR_OK {
		return errors.New(err_str)
	} else {
		return nil
	}
}

// RemoveHiddenText removes hidden text from PDF-document.
//
// Example:
//
//	err := pdf.RemoveHiddenText()
func (document *Document) RemoveHiddenText() error {
	var err *C.char
	C.PDFDocument_RemoveHiddenText(document.pdf, &err)
	err_str := C.GoString(err)
	C.c_free_string(err)
	if err_str != ERR_OK {
		return errors.New(err_str)
	} else {
		return nil
	}
}

// RemoveImages removes images from PDF-document.
//
// Example:
//
//	err := pdf.RemoveImages()
func (document *Document) RemoveImages() error {
	var err *C.char
	C.PDFDocument_RemoveImages(document.pdf, &err)
	err_str := C.GoString(err)
	C.c_free_string(err)
	if err_str != ERR_OK {
		return errors.New(err_str)
	} else {
		return nil
	}
}

// RemoveJavaScripts removes java scripts from PDF-document.
//
// Example:
//
//	err := pdf.RemoveJavaScripts()
func (document *Document) RemoveJavaScripts() error {
	var err *C.char
	C.PDFDocument_RemoveJavaScripts(document.pdf, &err)
	err_str := C.GoString(err)
	C.c_free_string(err)
	if err_str != ERR_OK {
		return errors.New(err_str)
	} else {
		return nil
	}
}

// RemoveTables removes tables from PDF-document.
//
// Example:
//
//	err := pdf.RemoveTables()
func (document *Document) RemoveTables() error {
	var err *C.char
	C.PDFDocument_RemoveTables(document.pdf, &err)
	err_str := C.GoString(err)
	C.c_free_string(err)
	if err_str != ERR_OK {
		return errors.New(err_str)
	} else {
		return nil
	}
}

// SetBackground sets PDF-document background color.
//
// Example:
//
//	err := pdf.SetBackground(200, 100, 101)
func (document *Document) SetBackground(r, g, b int32) error {
	var err *C.char
	C.PDFDocument_set_Background(document.pdf, C.int(r), C.int(g), C.int(b), &err)
	err_str := C.GoString(err)
	C.c_free_string(err)
	if err_str != ERR_OK {
		return errors.New(err_str)
	} else {
		return nil
	}
}

// Rotate rotates PDF-document.
//
// Example:
//
//	err := pdf.Rotate(asposepdf.RotationOn180)
func (document *Document) Rotate(rotation int32) error {
	var err *C.char
	C.PDFDocument_Rotate(document.pdf, C.int(rotation), &err)
	err_str := C.GoString(err)
	C.c_free_string(err)
	if err_str != ERR_OK {
		return errors.New(err_str)
	} else {
		return nil
	}
}

// WordCount returns word count in PDF-document.
//
// Example:
//
//	word_count, err := pdf.WordCount()
func (document *Document) WordCount() (int32, error) {
	var err *C.char
	cnt_int := C.PDFDocument_get_WordCount(document.pdf, &err)
	err_str := C.GoString(err)
	C.c_free_string(err)
	if err_str != ERR_OK {
		return -1, errors.New(err_str)
	} else {
		return int32(cnt_int), nil
	}
}

// CharacterCount returns character count in PDF-document.
//
// Example:
//
//	character_count, err := pdf.CharacterCount()
func (document *Document) CharacterCount() (int32, error) {
	var err *C.char
	cnt_int := C.PDFDocument_get_CharacterCount(document.pdf, &err)
	err_str := C.GoString(err)
	C.c_free_string(err)
	if err_str != ERR_OK {
		return -1, errors.New(err_str)
	} else {
		return int32(cnt_int), nil
	}
}

// ReplaceText replaces text in PDF-document.
//
// Example:
//
//	err := pdf.ReplaceText("old text", "new text")
func (document *Document) ReplaceText(findText, replaceText string) error {
	var err *C.char
	_findText := C.CString(findText)
	defer C.free(unsafe.Pointer(_findText))
	_replaceText := C.CString(replaceText)
	defer C.free(unsafe.Pointer(_replaceText))
	C.PDFDocument_ReplaceText(document.pdf, _findText, _replaceText, &err)
	err_str := C.GoString(err)
	C.c_free_string(err)
	if err_str != ERR_OK {
		return errors.New(err_str)
	} else {
		return nil
	}
}

// AddPageNum adds page number to a PDF-document.
//
// Example:
//
//	err := pdf.AddPageNum()
func (document *Document) AddPageNum() error {
	var err *C.char
	C.PDFDocument_AddPageNum(document.pdf, &err)
	err_str := C.GoString(err)
	C.c_free_string(err)
	if err_str != ERR_OK {
		return errors.New(err_str)
	} else {
		return nil
	}
}

// AddTextHeader adds text in Header of a PDF-document.
//
// Example:
//
//	err := pdf.AddTextHeader("Aspose")
func (document *Document) AddTextHeader(header string) error {
	var err *C.char
	_header := C.CString(header)
	defer C.free(unsafe.Pointer(_header))
	C.PDFDocument_AddTextHeader(document.pdf, _header, &err)
	err_str := C.GoString(err)
	C.c_free_string(err)
	if err_str != ERR_OK {
		return errors.New(err_str)
	} else {
		return nil
	}
}

// AddTextFooter adds text in Footer of a PDF-document.
//
// Example:
//
//	err := pdf.AddTextFooter("Footer")
func (document *Document) AddTextFooter(footer string) error {
	var err *C.char
	_footer := C.CString(footer)
	defer C.free(unsafe.Pointer(_footer))
	C.PDFDocument_AddTextFooter(document.pdf, _footer, &err)
	err_str := C.GoString(err)
	C.c_free_string(err)
	if err_str != ERR_OK {
		return errors.New(err_str)
	} else {
		return nil
	}
}

// SaveDocX saves previously opened PDF-document as DocX-document with filename.
//
// Example:
//
//	err := pdf.SaveDocX("filename.docx")
func (document *Document) SaveDocX(filename string) error {
	var err *C.char
	_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(_filename))
	C.PDFDocument_Save_DocX(document.pdf, _filename, &err)
	err_str := C.GoString(err)
	C.c_free_string(err)
	if err_str != ERR_OK {
		return errors.New(err_str)
	} else {
		return nil
	}
}

// SaveDocXEnhanced saves previously opened PDF-document as Enhanced Recognition Mode DocX-document with filename.
//
// Example:
//
//	err := pdf.SaveDocXEnhanced("filename.docx")
func (document *Document) SaveDocXEnhanced(filename string) error {
	var err *C.char
	_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(_filename))
	C.PDFDocument_Save_DocXEnhanced(document.pdf, _filename, &err)
	err_str := C.GoString(err)
	C.c_free_string(err)
	if err_str != ERR_OK {
		return errors.New(err_str)
	} else {
		return nil
	}
}

// SaveDoc saves previously opened PDF-document as Doc-document with filename.
//
// Example:
//
//	err := pdf.SaveDoc("filename.doc")
func (document *Document) SaveDoc(filename string) error {
	var err *C.char
	_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(_filename))
	C.PDFDocument_Save_Doc(document.pdf, _filename, &err)
	err_str := C.GoString(err)
	C.c_free_string(err)
	if err_str != ERR_OK {
		return errors.New(err_str)
	} else {
		return nil
	}
}

// SaveXlsX saves previously opened PDF-document as XlsX-document with filename.
//
// Example:
//
//	err := pdf.SaveXlsX("filename.xlsx")
func (document *Document) SaveXlsX(filename string) error {
	var err *C.char
	_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(_filename))
	C.PDFDocument_Save_XlsX(document.pdf, _filename, &err)
	err_str := C.GoString(err)
	C.c_free_string(err)
	if err_str != ERR_OK {
		return errors.New(err_str)
	} else {
		return nil
	}
}

// SavePptX saves previously opened PDF-document as PptX-document with filename.
//
// Example:
//
//	err := pdf.SavePptX("filename.pptx")
func (document *Document) SavePptX(filename string) error {
	var err *C.char
	_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(_filename))
	C.PDFDocument_Save_PptX(document.pdf, _filename, &err)
	err_str := C.GoString(err)
	C.c_free_string(err)
	if err_str != ERR_OK {
		return errors.New(err_str)
	} else {
		return nil
	}
}

// SaveXps saves previously opened PDF-document as Xps-document with filename.
//
// Example:
//
//	err := pdf.SaveXps("filename.xps")
func (document *Document) SaveXps(filename string) error {
	var err *C.char
	_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(_filename))
	C.PDFDocument_Save_Xps(document.pdf, _filename, &err)
	err_str := C.GoString(err)
	C.c_free_string(err)
	if err_str != ERR_OK {
		return errors.New(err_str)
	} else {
		return nil
	}
}

// SaveTxt saves previously opened PDF-document as Txt-document with filename.
//
// Example:
//
//	err := pdf.SaveTxt("filename.txt")
func (document *Document) SaveTxt(filename string) error {
	var err *C.char
	_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(_filename))
	C.PDFDocument_Save_Txt(document.pdf, _filename, &err)
	err_str := C.GoString(err)
	C.c_free_string(err)
	if err_str != ERR_OK {
		return errors.New(err_str)
	} else {
		return nil
	}
}

// SaveEpub saves previously opened PDF-document as Epub-document with filename.
//
// Example:
//
//	err := pdf.SaveEpub("filename.epub")
func (document *Document) SaveEpub(filename string) error {
	var err *C.char
	_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(_filename))
	C.PDFDocument_Save_Epub(document.pdf, _filename, &err)
	err_str := C.GoString(err)
	C.c_free_string(err)
	if err_str != ERR_OK {
		return errors.New(err_str)
	} else {
		return nil
	}
}

// SaveTeX saves previously opened PDF-document as TeX-document with filename.
//
// Example:
//
//	err := pdf.SaveTeX("filename.tex")
func (document *Document) SaveTeX(filename string) error {
	var err *C.char
	_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(_filename))
	C.PDFDocument_Save_TeX(document.pdf, _filename, &err)
	err_str := C.GoString(err)
	C.c_free_string(err)
	if err_str != ERR_OK {
		return errors.New(err_str)
	} else {
		return nil
	}
}

// SaveMarkdown saves previously opened PDF-document as Markdown-document with filename.
//
// Example:
//
//	err := pdf.SaveMarkdown("filename.md")
func (document *Document) SaveMarkdown(filename string) error {
	var err *C.char
	_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(_filename))
	C.PDFDocument_Save_Markdown(document.pdf, _filename, &err)
	err_str := C.GoString(err)
	C.c_free_string(err)
	if err_str != ERR_OK {
		return errors.New(err_str)
	} else {
		return nil
	}
}

// SaveBooklet saves previously opened PDF-document as booklet PDF-document with filename.
//
// Example:
//
//	err := pdf.SaveBooklet("filename.pdf")
func (document *Document) SaveBooklet(filename string) error {
	var err *C.char
	_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(_filename))
	C.PDFDocument_Save_Booklet(document.pdf, _filename, &err)
	err_str := C.GoString(err)
	C.c_free_string(err)
	if err_str != ERR_OK {
		return errors.New(err_str)
	} else {
		return nil
	}
}

// SaveNUp saves previously opened PDF-document as N-Up PDF-document with filename.
//
// Example:
//
//	err := pdf.SaveNUp("filename.pdf", 2, 2)
func (document *Document) SaveNUp(filename string, columns int32, rows int32) error {
	var err *C.char
	_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(_filename))
	C.PDFDocument_Save_NUp(document.pdf, _filename, C.int(columns), C.int(rows), &err)
	err_str := C.GoString(err)
	C.c_free_string(err)
	if err_str != ERR_OK {
		return errors.New(err_str)
	} else {
		return nil
	}
}

// SaveTiff saves previously opened PDF-document as Tiff-document with filename.
//
// Example:
//
//	err := pdf.SaveTiff("filename.tiff")
func (document *Document) SaveTiff(filename string, resolution_dpi ...int32) error {
	var err *C.char
	_filename := C.CString(filename)
	_resolution_dpi := C.int(100)
	defer C.free(unsafe.Pointer(_filename))
	if len(resolution_dpi) > 0 {
		_resolution_dpi = C.int(resolution_dpi[0])
	}
	C.PDFDocument_Save_Tiff(document.pdf, _resolution_dpi, _filename, &err)
	err_str := C.GoString(err)
	C.c_free_string(err)
	if err_str != ERR_OK {
		return errors.New(err_str)
	} else {
		return nil
	}
}

// SaveSvgZip saves previously opened PDF-document as SVG-archive with filename.
//
// Example:
//
//	err := pdf.SaveSvgZip("filename.zip")
func (document *Document) SaveSvgZip(filename string) error {
	var err *C.char
	_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(_filename))
	C.PDFDocument_Save_SvgZip(document.pdf, _filename, &err)
	err_str := C.GoString(err)
	C.c_free_string(err)
	if err_str != ERR_OK {
		return errors.New(err_str)
	} else {
		return nil
	}
}

// ExportFdf exports from previously opened PDF-document with AcroForm to FDF-document with filename.
//
// Example:
//
//	err := pdf.ExportFdf("filename.fdf")
func (document *Document) ExportFdf(filename string) error {
	var err *C.char
	_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(_filename))
	C.PDFDocument_Export_Fdf(document.pdf, _filename, &err)
	err_str := C.GoString(err)
	C.c_free_string(err)
	if err_str != ERR_OK {
		return errors.New(err_str)
	} else {
		return nil
	}
}

// ExportXfdf exports from previously opened PDF-document with AcroForm to XFDF-document with filename.
//
// Example:
//
//	err := pdf.ExportXfdf("filename.xfdf")
func (document *Document) ExportXfdf(filename string) error {
	var err *C.char
	_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(_filename))
	C.PDFDocument_Export_Xfdf(document.pdf, _filename, &err)
	err_str := C.GoString(err)
	C.c_free_string(err)
	if err_str != ERR_OK {
		return errors.New(err_str)
	} else {
		return nil
	}
}

// ExportXml exports from previously opened PDF-document with AcroForm to XML-document with filename.
//
// Example:
//
//	err := pdf.ExportXml("filename.xml")
func (document *Document) ExportXml(filename string) error {
	var err *C.char
	_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(_filename))
	C.PDFDocument_Export_Xml(document.pdf, _filename, &err)
	err_str := C.GoString(err)
	C.c_free_string(err)
	if err_str != ERR_OK {
		return errors.New(err_str)
	} else {
		return nil
	}
}

// Append appends pages from another PDF-document.
//
// Example:
//
//	err := pdf.Append(anotherdoc)
func (document *Document) Append(anotherdocument *Document) error {
	var err *C.char
	C.PDFDocument_Append(document.pdf, anotherdocument.pdf, &err)
	err_str := C.GoString(err)
	C.c_free_string(err)
	if err_str != ERR_OK {
		return errors.New(err_str)
	} else {
		return nil
	}
}

// AppendPages appends selected pages from another PDF-document.
//
// Example:
//
//	err := pdf.AppendPages(anotherdoc, "-2,4,6-8,10-")
func (document *Document) AppendPages(anotherdocument *Document, pagerange string) error {
	var err *C.char
	_pagerange := C.CString(pagerange)
	defer C.free(unsafe.Pointer(_pagerange))
	C.PDFDocument_AppendPages(document.pdf, anotherdocument.pdf, _pagerange, &err)
	err_str := C.GoString(err)
	C.c_free_string(err)
	if err_str != ERR_OK {
		return errors.New(err_str)
	} else {
		return nil
	}
}

// Bytes returns the contents of the PDF-document as a byte slice.
//
// Example:
//
//	bytes, err := pdf.Bytes()
func (document *Document) Bytes() ([]byte, error) {
	var err *C.char
	var buf *C.uchar
	var size C.int

	C.PDFDocument_Save_Memory(document.pdf, &buf, &size, &err)
	defer C.c_free_string(err)

	err_str := C.GoString(err)
	if err_str != "" || buf == nil || size == 0 {
		return nil, fmt.Errorf("failed to get PDF bytes: %s", err_str)
	}

	defer C.c_free_buffer(unsafe.Pointer(buf))
	return C.GoBytes(unsafe.Pointer(buf), size), nil
}

// PageCount returns page count in PDF-document.
//
// Example:
//
//	page_count, err := pdf.PageCount()
func (document *Document) PageCount() (int32, error) {
	var err *C.char
	cnt_int := C.PDFDocument_Page_get_Count(document.pdf, &err)
	err_str := C.GoString(err)
	C.c_free_string(err)
	if err_str != ERR_OK {
		return -1, errors.New(err_str)
	} else {
		return int32(cnt_int), nil
	}
}

// PageAdd adds new page in PDF-document.
//
// Example:
//
//	err := pdf.PageAdd()
func (document *Document) PageAdd() error {
	var err *C.char
	C.PDFDocument_Page_Add(document.pdf, &err)
	err_str := C.GoString(err)
	C.c_free_string(err)
	if err_str != ERR_OK {
		return errors.New(err_str)
	} else {
		return nil
	}
}

// PageInsert inserts new page at the specified position in PDF-document.
//
// Example:
//
//	err := pdf.PageInsert(1)
func (document *Document) PageInsert(num int32) error {
	var err *C.char
	C.PDFDocument_Page_Insert(document.pdf, C.int(num), &err)
	err_str := C.GoString(err)
	C.c_free_string(err)
	if err_str != ERR_OK {
		return errors.New(err_str)
	} else {
		return nil
	}
}

// PageDelete deletes specified page in PDF-document.
//
// Example:
//
//	err := pdf.PageDelete(1)
func (document *Document) PageDelete(num int32) error {
	var err *C.char
	C.PDFDocument_Page_Delete(document.pdf, C.int(num), &err)
	err_str := C.GoString(err)
	C.c_free_string(err)
	if err_str != ERR_OK {
		return errors.New(err_str)
	} else {
		return nil
	}
}

// PageToJpg saves the specified page as Jpg-image file.
//
// Example:
//
//	err := pdf.PageToJpg(1, 300, "page_num_1_with_300_dpi.jpg")
func (document *Document) PageToJpg(num int32, resolution_dpi int32, filename string) error {
	var err *C.char
	_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(_filename))
	C.PDFDocument_Page_to_Jpg(document.pdf, C.int(num), C.int(resolution_dpi), _filename, &err)
	err_str := C.GoString(err)
	C.c_free_string(err)
	if err_str != ERR_OK {
		return errors.New(err_str)
	} else {
		return nil
	}
}

// PageToPng saves the specified page as Png-image file.
//
// Example:
//
//	err := pdf.PageToPng(1, 100, "page_num_1_with_100_dpi.png")
func (document *Document) PageToPng(num int32, resolution_dpi int32, filename string) error {
	var err *C.char
	_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(_filename))
	C.PDFDocument_Page_to_Png(document.pdf, C.int(num), C.int(resolution_dpi), _filename, &err)
	err_str := C.GoString(err)
	C.c_free_string(err)
	if err_str != ERR_OK {
		return errors.New(err_str)
	} else {
		return nil
	}
}

// PageToBmp saves the specified page as Bmp-image file.
//
// Example:
//
//	err := pdf.PageToBmp(1, 100, "page_num_1_with_100_dpi.bmp")
func (document *Document) PageToBmp(num int32, resolution_dpi int32, filename string) error {
	var err *C.char
	_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(_filename))
	C.PDFDocument_Page_to_Bmp(document.pdf, C.int(num), C.int(resolution_dpi), _filename, &err)
	err_str := C.GoString(err)
	C.c_free_string(err)
	if err_str != ERR_OK {
		return errors.New(err_str)
	} else {
		return nil
	}
}

// PageToTiff saves the specified page as Tiff-image file.
//
// Example:
//
//	err := pdf.PageToTiff(1, 100, "page_num_1_with_100_dpi.tiff")
func (document *Document) PageToTiff(num int32, resolution_dpi int32, filename string) error {
	var err *C.char
	_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(_filename))
	C.PDFDocument_Page_to_Tiff(document.pdf, C.int(num), C.int(resolution_dpi), _filename, &err)
	err_str := C.GoString(err)
	C.c_free_string(err)
	if err_str != ERR_OK {
		return errors.New(err_str)
	} else {
		return nil
	}
}

// PageToDICOM saves the specified page as DICOM-image file.
//
// Example:
//
//	err := pdf.PageToDICOM(1, 100, "page_num_1_with_100_dpi.dcm")
func (document *Document) PageToDICOM(num int32, resolution_dpi int32, filename string) error {
	var err *C.char
	_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(_filename))
	C.PDFDocument_Page_to_DICOM(document.pdf, C.int(num), C.int(resolution_dpi), _filename, &err)
	err_str := C.GoString(err)
	C.c_free_string(err)
	if err_str != ERR_OK {
		return errors.New(err_str)
	} else {
		return nil
	}
}

// PageToSvg saves the specified page as Svg-image file.
//
// Example:
//
//	err := pdf.PageToSvg(1, "page_num_1.svg")
func (document *Document) PageToSvg(num int32, filename string) error {
	var err *C.char
	_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(_filename))
	C.PDFDocument_Page_to_Svg(document.pdf, C.int(num), _filename, &err)
	err_str := C.GoString(err)
	C.c_free_string(err)
	if err_str != ERR_OK {
		return errors.New(err_str)
	} else {
		return nil
	}
}

// PageToPdf saves the specified page as Pdf-file.
//
// Example:
//
//	err := pdf.PageToPdf(1, "page_num_1.pdf")
func (document *Document) PageToPdf(num int32, filename string) error {
	var err *C.char
	_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(_filename))
	C.PDFDocument_Page_to_Pdf(document.pdf, C.int(num), _filename, &err)
	err_str := C.GoString(err)
	C.c_free_string(err)
	if err_str != ERR_OK {
		return errors.New(err_str)
	} else {
		return nil
	}
}

// PageGrayscale converts page to black and white.
//
// Example:
//
//	err := pdf.PageGrayscale(1)
func (document *Document) PageGrayscale(num int32) error {
	var err *C.char
	C.PDFDocument_Page_Grayscale(document.pdf, C.int(num), &err)
	err_str := C.GoString(err)
	C.c_free_string(err)
	if err_str != ERR_OK {
		return errors.New(err_str)
	} else {
		return nil
	}
}

// PageRotate rotates page.
//
// Example:
//
//	err := pdf.PageRotate(1, asposepdf.RotationOn180)
func (document *Document) PageRotate(num int32, rotation int32) error {
	var err *C.char
	C.PDFDocument_Page_Rotate(document.pdf, C.int(num), C.int(rotation), &err)
	err_str := C.GoString(err)
	C.c_free_string(err)
	if err_str != ERR_OK {
		return errors.New(err_str)
	} else {
		return nil
	}
}

// PageReplaceText replaces text on page.
//
// Example:
//
//	err := pdf.PageReplaceText(1, "old text", "new text")
func (document *Document) PageReplaceText(num int32, findText, replaceText string) error {
	var err *C.char
	_findText := C.CString(findText)
	defer C.free(unsafe.Pointer(_findText))
	_replaceText := C.CString(replaceText)
	defer C.free(unsafe.Pointer(_replaceText))
	C.PDFDocument_Page_ReplaceText(document.pdf, C.int(num), _findText, _replaceText, &err)
	err_str := C.GoString(err)
	C.c_free_string(err)
	if err_str != ERR_OK {
		return errors.New(err_str)
	} else {
		return nil
	}
}

// PageAddText adds text on page.
//
// Example:
//
//	err := pdf.PageAddText(1, "text on the first page")
func (document *Document) PageAddText(num int32, addText string) error {
	var err *C.char
	_addText := C.CString(addText)
	defer C.free(unsafe.Pointer(_addText))
	C.PDFDocument_Page_AddText(document.pdf, C.int(num), _addText, &err)
	err_str := C.GoString(err)
	C.c_free_string(err)
	if err_str != ERR_OK {
		return errors.New(err_str)
	} else {
		return nil
	}
}

// PageSetSize sets size of page.
//
// Example:
//
//	err := pdf.PageSetSize(1, asposepdf.PageSizeA4)
func (document *Document) PageSetSize(num int32, pageSize int32) error {
	var err *C.char
	C.PDFDocument_Page_set_Size(document.pdf, C.int(num), C.int(pageSize), &err)
	err_str := C.GoString(err)
	C.c_free_string(err)
	if err_str != ERR_OK {
		return errors.New(err_str)
	} else {
		return nil
	}
}

// PageWordCount returns word count on specified page in PDF-document.
//
// Example:
//
//	word_count, err := pdf.PageWordCount(1)
func (document *Document) PageWordCount(num int32) (int32, error) {
	var err *C.char
	cnt_int := C.PDFDocument_Page_get_WordCount(document.pdf, C.int(num), &err)
	err_str := C.GoString(err)
	C.c_free_string(err)
	if err_str != ERR_OK {
		return -1, errors.New(err_str)
	} else {
		return int32(cnt_int), nil
	}
}

// PageCharacterCount returns character count on specified page in PDF-document.
//
// Example:
//
//	character_count, err := pdf.PageCharacterCount(1)
func (document *Document) PageCharacterCount(num int32) (int32, error) {
	var err *C.char
	cnt_int := C.PDFDocument_Page_get_CharacterCount(document.pdf, C.int(num), &err)
	err_str := C.GoString(err)
	C.c_free_string(err)
	if err_str != ERR_OK {
		return -1, errors.New(err_str)
	} else {
		return int32(cnt_int), nil
	}
}

// PageIsBlank returns page is blank in PDF-document.
//
// Example:
//
//	is_blank, err := pdf.PageIsBlank(1)
func (document *Document) PageIsBlank(num int32) (bool, error) {
	var err *C.char
	is_blank_int := C.PDFDocument_Page_is_Blank(document.pdf, C.int(num), &err)
	err_str := C.GoString(err)
	C.c_free_string(err)
	if err_str != ERR_OK {
		return false, errors.New(err_str)
	} else {
		return is_blank_int != 0, nil
	}
}

// PageAddPageNum adds page number on page.
//
// Example:
//
//	err := pdf.PageAddPageNum(1)
func (document *Document) PageAddPageNum(num int32) error {
	var err *C.char
	C.PDFDocument_Page_AddPageNum(document.pdf, C.int(num), &err)
	err_str := C.GoString(err)
	C.c_free_string(err)
	if err_str != ERR_OK {
		return errors.New(err_str)
	} else {
		return nil
	}
}

// PageAddTextHeader adds text in page header
//
// Example:
//
//	err := pdf.PageAddTextHeader(1, "Aspose")
func (document *Document) PageAddTextHeader(num int32, header string) error {
	var err *C.char
	_header := C.CString(header)
	defer C.free(unsafe.Pointer(_header))
	C.PDFDocument_Page_AddTextHeader(document.pdf, C.int(num), _header, &err)
	err_str := C.GoString(err)
	C.c_free_string(err)
	if err_str != ERR_OK {
		return errors.New(err_str)
	} else {
		return nil
	}
}

// PageAddTextFooter adds text in page footer
//
// Example:
//
//	err := pdf.PageAddTextFooter("Footer")
func (document *Document) PageAddTextFooter(num int32, footer string) error {
	var err *C.char
	_footer := C.CString(footer)
	defer C.free(unsafe.Pointer(_footer))
	C.PDFDocument_Page_AddTextFooter(document.pdf, C.int(num), _footer, &err)
	err_str := C.GoString(err)
	C.c_free_string(err)
	if err_str != ERR_OK {
		return errors.New(err_str)
	} else {
		return nil
	}
}

// PageRemoveAnnotations removes annotations in page.
//
// Example:
//
//	err := pdf.PageRemoveAnnotations(1)
func (document *Document) PageRemoveAnnotations(num int32) error {
	var err *C.char
	C.PDFDocument_Page_RemoveAnnotations(document.pdf, C.int(num), &err)
	err_str := C.GoString(err)
	C.c_free_string(err)
	if err_str != ERR_OK {
		return errors.New(err_str)
	} else {
		return nil
	}
}

// PageRemoveHiddenText removes hidden text in page.
//
// Example:
//
//	err := pdf.PageRemoveHiddenText(1)
func (document *Document) PageRemoveHiddenText(num int32) error {
	var err *C.char
	C.PDFDocument_Page_RemoveHiddenText(document.pdf, C.int(num), &err)
	err_str := C.GoString(err)
	C.c_free_string(err)
	if err_str != ERR_OK {
		return errors.New(err_str)
	} else {
		return nil
	}
}

// PageRemoveImages removes images in page.
//
// Example:
//
//	err := pdf.PageRemoveImages(1)
func (document *Document) PageRemoveImages(num int32) error {
	var err *C.char
	C.PDFDocument_Page_RemoveImages(document.pdf, C.int(num), &err)
	err_str := C.GoString(err)
	C.c_free_string(err)
	if err_str != ERR_OK {
		return errors.New(err_str)
	} else {
		return nil
	}
}

// PageRemoveTables removes tables in page.
//
// Example:
//
//	err := pdf.PageRemoveTables(1)
func (document *Document) PageRemoveTables(num int32) error {
	var err *C.char
	C.PDFDocument_Page_RemoveTables(document.pdf, C.int(num), &err)
	err_str := C.GoString(err)
	C.c_free_string(err)
	if err_str != ERR_OK {
		return errors.New(err_str)
	} else {
		return nil
	}
}
