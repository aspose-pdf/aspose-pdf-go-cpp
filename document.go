package asposepdf

/*
#cgo CFLAGS: -g -Wall
#include <stdlib.h>
#include "extern_c.h"
*/
import "C"
import "unsafe"
import "errors"

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
