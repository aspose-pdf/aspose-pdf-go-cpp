// use:
// go test -v

package asposepdf

import (
	"fmt"
	"os"
	"reflect"
	"strings"
	"testing"
)

func assert_check_types(t *testing.T, left any, right any) {
	if reflect.TypeOf(left) != reflect.TypeOf(right) {
		t.Fatalf("Different types: %v (type %v) and %v (type %v)", left, reflect.TypeOf(left), right, reflect.TypeOf(right))
	}
}

func assert_eq(t *testing.T, left any, right any) {
	assert_check_types(t, left, right)
	if !reflect.DeepEqual(left, right) {
		t.Errorf("%v != %v", left, right)
	}
}

func assert_ne(t *testing.T, left any, right any) {
	assert_check_types(t, left, right)
	if reflect.DeepEqual(left, right) {
		t.Errorf("%v == %v", left, right)
	}
}

func TestNewAndSave(t *testing.T) {

	pdf_new_filename := fmt.Sprintf("%s/new.pdf", t.TempDir())

	pdf_new, err := New()
	if err != nil {
		t.Errorf("New(): %v", err)
	} else {
		err = pdf_new.SaveAs(pdf_new_filename)
		if err != nil {
			t.Errorf("SaveAs(%s): %v", pdf_new_filename, err)
		} else {
			pdf_empty, err := Open(pdf_new_filename)
			if err != nil {
				t.Errorf("Open(%s): %v", pdf_new_filename, err)
			} else {
				err = pdf_empty.Save()
				if err != nil {
					t.Errorf("Save(%s): %v", pdf_new_filename, err)
				}
			}
			defer pdf_empty.Close()
		}
	}
	defer pdf_new.Close()

	// Check file size != 0
	fi, err := os.Stat(pdf_new_filename)
	if err != nil {
		t.Errorf("Stat(%s): %v", pdf_new_filename, err)
	}
	assert_ne(t, int64(0), fi.Size())
}

func TestMergeDocuments(t *testing.T) {
	// Create the first PDF-document
	pdf1, err := New()
	if err != nil {
		t.Errorf("New(): %v", err)
	}
	defer pdf1.Close()

	// Add one empty page to the first PDF-document
	_ = pdf1.PageAdd()

	// Create the second PDF-document
	pdf2, err := New()
	if err != nil {
		t.Errorf("New(): %v", err)
	}
	defer pdf2.Close()

	// Add one empty page to the second PDF-document
	_ = pdf2.PageAdd()

	// Create the third PDF-document
	pdf3, err := New()
	if err != nil {
		t.Errorf("New(): %v", err)
	}
	defer pdf3.Close()

	// Add two empty pages to the third PDF-document
	for i := 0; i < 2; i++ {
		_ = pdf3.PageAdd()
	}

	// Merge all three documents
	mergedDoc, err := MergeDocuments([]*Document{pdf1, pdf2, pdf3})
	if err != nil {
		t.Errorf("MergeDocuments(): %v", err)
	}
	defer mergedDoc.Close()

	// Check the page count in the merged document (should be 4 pages)
	mergedPageCount, err := mergedDoc.PageCount()
	if err != nil {
		t.Errorf("PageCount(): %v", err)
	}
	assert_eq(t, mergedPageCount, int32(4))
}

func TestSplitDocument(t *testing.T) {
	// Create a PDF-document with 4 pages
	pdf, err := New()
	if err != nil {
		t.Errorf("New(): %v", err)
	}
	defer pdf.Close()

	for i := 0; i < 4; i++ {
		_ = pdf.PageAdd()
	}

	// Split the PDF-document into 3 parts: pages 1-2, page 3, pages 4 to end
	pdfs, err := SplitDocument(pdf, "1-2;3;4-")
	if err != nil {
		t.Errorf("SplitDocument(): %v", err)
	}

	if len(pdfs) != 3 {
		t.Errorf("expected 3 split documents, got %d", len(pdfs))
	}

	// Defer closing all PDF-documents
	for _, doc := range pdfs {
		defer doc.Close()
	}

	// Check page counts for each resulting PDF-document
	if len(pdfs) == 3 {
		count1, _ := pdfs[0].PageCount()
		assert_eq(t, count1, int32(2))

		count2, _ := pdfs[1].PageCount()
		assert_eq(t, count2, int32(1))

		count3, _ := pdfs[2].PageCount()
		assert_eq(t, count3, int32(1))
	}
}

func TestSplitAtPage(t *testing.T) {
	// Create a PDF-document with 4 pages
	pdf, err := New()
	if err != nil {
		t.Errorf("New(): %v", err)
	}
	defer pdf.Close()

	for i := 0; i < 4; i++ {
		_ = pdf.PageAdd()
	}

	// Split the PDF-document into two parts: first 2 pages, and remaining pages
	left, right, err := SplitAtPage(pdf, 2)
	if err != nil {
		t.Errorf("SplitAtPage(): %v", err)
	}

	// Defer closing both PDF-documents
	defer left.Close()
	defer right.Close()

	// Check page counts for each resulting PDF-document
	countLeft, _ := left.PageCount()
	assert_eq(t, countLeft, int32(2))

	countRight, _ := right.PageCount()
	assert_eq(t, countRight, int32(2))
}

func TestPages(t *testing.T) {

	pdf, _ := New()
	defer pdf.Close()

	// Add page
	_ = pdf.PageAdd()
	// Set page size
	_ = pdf.PageSetSize(1, PageSizeA1)
	// Insert page at first position
	_ = pdf.PageInsert(1)
	// Delete first page
	_ = pdf.PageDelete(1)
	// Gets number of pages
	page_count, _ := pdf.PageCount()
	// The number of pages must be equal to 1
	assert_eq(t, page_count, int32(1))
}

func TestStats(t *testing.T) {

	pdf, _ := New()
	defer pdf.Close()

	// Text with stamp: "Evaluation Only. Created with Aspose.PDF ..."

	// Add page
	_ = pdf.PageAdd()
	// Save
	_ = pdf.SaveAs(fmt.Sprintf("%s/stat.pdf", t.TempDir()))

	// Word count
	word_count, _ := pdf.WordCount()
	// The number of word count must be equal to 12
	assert_eq(t, word_count, int32(12))
	// Character count
	character_count, _ := pdf.CharacterCount()
	// The number of character count must be equal to 74
	assert_eq(t, character_count, int32(74))

	// Word count on first page
	page_word_count, _ := pdf.PageWordCount(1)
	// The number of word count must be equal to 12
	assert_eq(t, page_word_count, int32(12))
	// Character count on first page
	page_character_count, _ := pdf.PageCharacterCount(1)
	// The number of character count must be equal to 74
	assert_eq(t, page_character_count, int32(74))

	// First page is blank? No. Evaluation stamp
	page_is_blank, _ := pdf.PageIsBlank(1)
	assert_eq(t, page_is_blank, false)
}

func TestAppend(t *testing.T) {
	// Create the first PDF-document
	pdf1, err := New()
	if err != nil {
		t.Errorf("New(): %v", err)
	}
	defer pdf1.Close()

	// Add two empty pages to the first PDF-document
	for i := 0; i < 2; i++ {
		_ = pdf1.PageAdd()
	}

	// Check the page count in the first PDF-document (should be 2 pages)
	page_count1, _ := pdf1.PageCount()
	assert_eq(t, page_count1, int32(2))

	// Create the second PDF-document
	pdf2, err := New()
	if err != nil {
		t.Errorf("New(): %v", err)
	}
	defer pdf2.Close()

	// Add two empty pages to the second PDF-document
	for i := 0; i < 2; i++ {
		_ = pdf2.PageAdd()
	}

	// Check the page count in the second PDF-document (should be 2 pages)
	page_count2, _ := pdf2.PageCount()
	assert_eq(t, page_count2, int32(2))

	// Append pages from the second PDF-document to the first
	err = pdf1.Append(pdf2)
	if err != nil {
		t.Errorf("Append(): %v", err)
	}

	// Check the page count in the first PDF-document after appending (should be 4 pages)
	page_count1_after_append, _ := pdf1.PageCount()
	assert_eq(t, page_count1_after_append, int32(4))
}

func TestAppendPages(t *testing.T) {
	// Create a PDF document with exactly 4 pages
	pdf4pages, err := New()
	if err != nil {
		t.Fatalf("New(): %v", err)
	}
	defer pdf4pages.Close()

	for i := 0; i < 4; i++ {
		_ = pdf4pages.PageAdd()
	}

	page_count, _ := pdf4pages.PageCount()
	assert_eq(t, page_count, int32(4))

	tests := []struct {
		name      string
		pagerange string
		wantPages int32
	}{
		{"EmptyRangeMeansAll", "", 4},
		{"DashMeansAll", "-", 4},
		{"FirstThreePages", "-3", 3},
		{"SecondToEnd", "2-", 3},
		{"SpecificPages134", "1,3,4", 3},
		{"OnlyPage2", "2", 1},
		{"Range2To3", "2-3", 2},
		{"NonSequential", "1,2,4", 3},
		{"AllPagesExplicit", "1,2,3,4", 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a new empty document for each test case
			pdftest, err := New()
			if err != nil {
				t.Fatalf("New(): %v", err)
			}
			defer pdftest.Close()

			// Append selected pages from pdf4pages
			err = pdftest.AppendPages(pdf4pages, tt.pagerange)
			if err != nil {
				t.Errorf("AppendPages(%q): %v", tt.pagerange, err)
			}

			// Check that the page count matches expected result
			count, _ := pdftest.PageCount()
			assert_eq(t, count, tt.wantPages)
		})
	}
}

func TestExtractText(t *testing.T) {
	// Create a new document
	doc, err := New()
	if err != nil {
		t.Fatalf("New(): %v", err)
	}
	defer doc.Close()

	// Add a page and some text
	if err := doc.PageAdd(); err != nil {
		t.Fatalf("PageAdd(): %v", err)
	}
	expectedText := "This is a test text for extraction"
	if err := doc.PageAddText(1, expectedText); err != nil {
		t.Fatalf("PageAddText(): %v", err)
	}

	// Save to apply changes
	if err := doc.Save(); err != nil {
		t.Fatalf("Save(): %v", err)
	}

	// Extract text
	extracted, err := doc.ExtractText()
	if err != nil {
		t.Fatalf("ExtractText(): %v", err)
	}

	// Validate that extracted text contains the inserted text
	if len(extracted) == 0 {
		t.Errorf("Extracted text is empty")
	}
	if !strings.Contains(extracted, "test text") {
		t.Errorf("Extracted text does not contain expected content.\nGot: %s", extracted)
	}
}

func TestConvertFromPDF(t *testing.T) {
	type conversion struct {
		name string
		fn   func(doc *Document, path string) error
	}

	conversions := []conversion{
		{"SaveDocX", func(doc *Document, path string) error { return doc.SaveDocX(path) }},
		{"SaveDoc", func(doc *Document, path string) error { return doc.SaveDoc(path) }},
		{"SaveXlsX", func(doc *Document, path string) error { return doc.SaveXlsX(path) }},
		{"SaveTxt", func(doc *Document, path string) error { return doc.SaveTxt(path) }},
		{"SavePptX", func(doc *Document, path string) error { return doc.SavePptX(path) }},
		{"SaveXps", func(doc *Document, path string) error { return doc.SaveXps(path) }},
		{"SaveTeX", func(doc *Document, path string) error { return doc.SaveTeX(path) }},
		{"SaveEpub", func(doc *Document, path string) error { return doc.SaveEpub(path) }},
		{"SaveBooklet", func(doc *Document, path string) error { return doc.SaveBooklet(path) }},
		{"SaveNUp", func(doc *Document, path string) error { return doc.SaveNUp(path, 2, 2) }},
		{"SaveMarkdown", func(doc *Document, path string) error { return doc.SaveMarkdown(path) }},
		{"SaveTiff", func(doc *Document, path string) error { return doc.SaveTiff(path) }},
		{"SaveTiffWithDPI", func(doc *Document, path string) error { return doc.SaveTiff(path, 150) }},
		{"ExportFdf", func(doc *Document, path string) error { return doc.ExportFdf(path) }},
		{"ExportXfdf", func(doc *Document, path string) error { return doc.ExportXfdf(path) }},
		{"ExportXml", func(doc *Document, path string) error { return doc.ExportXml(path) }},
		{"PageToJpg", func(doc *Document, path string) error { return doc.PageToJpg(1, 150, path) }},
		{"PageToPng", func(doc *Document, path string) error { return doc.PageToPng(1, 150, path) }},
		{"PageToBmp", func(doc *Document, path string) error { return doc.PageToBmp(1, 150, path) }},
		{"PageToTiff", func(doc *Document, path string) error { return doc.PageToTiff(1, 150, path) }},
		{"PageToSvg", func(doc *Document, path string) error { return doc.PageToSvg(1, path) }},
		{"PageToPdf", func(doc *Document, path string) error { return doc.PageToPdf(1, path) }},
		{"PageToDICOM", func(doc *Document, path string) error { return doc.PageToDICOM(1, 150, path) }},
	}

	for _, conv := range conversions {
		t.Run(conv.name, func(t *testing.T) {
			// Create new document
			doc, err := New()
			if err != nil {
				t.Fatalf("New(): %v", err)
			}
			defer doc.Close()

			// Add one page with text
			if err := doc.PageAdd(); err != nil {
				t.Fatalf("PageAdd(): %v", err)
			}
			if err := doc.PageAddText(1, fmt.Sprintf("Test conversion for %s", conv.name)); err != nil {
				t.Fatalf("PageAddText(): %v", err)
			}

			// Save document before conversion
			if err := doc.Save(); err != nil {
				t.Fatalf("Save(): %v", err)
			}

			// Prepare output path
			outputPath := fmt.Sprintf("%s/output", t.TempDir())

			// Call conversion function
			if err := conv.fn(doc, outputPath); err != nil {
				t.Errorf("%s failed: %v", conv.name, err)
			}

			// Check file was created and is non-zero
			info, err := os.Stat(outputPath)
			if err != nil {
				t.Errorf("Stat(%s): %v", outputPath, err)
			} else {
				assert_ne(t, int64(0), info.Size())
			}
		})
	}
}

func TestOrganize(t *testing.T) {
	type organizeFunction struct {
		name string
		fn   func(doc *Document) error
	}

	organizeFunctions := []organizeFunction{
		{"Optimize", (*Document).Optimize},
		{"OptimizeResource", (*Document).OptimizeResource},
		{"Grayscale", (*Document).Grayscale},
		{"Rotate", func(doc *Document) error { return doc.Rotate(RotationOn270) }},
		{"SetBackground", func(doc *Document) error { return doc.SetBackground(255, 255, 200) }},
		{"ReplaceText", func(doc *Document) error {
			_ = doc.PageAddText(1, "Hello World")
			return doc.ReplaceText("Hello", "Hi")
		}},
		{"AddPageNum", (*Document).AddPageNum},
		{"AddTextHeader", func(doc *Document) error { return doc.AddTextHeader("Header") }},
		{"AddTextFooter", func(doc *Document) error { return doc.AddTextFooter("Footer") }},
		{"Flatten", (*Document).Flatten},
		{"RemoveAnnotations", (*Document).RemoveAnnotations},
		{"RemoveAttachments", (*Document).RemoveAttachments},
		{"RemoveBlankPages", (*Document).RemoveBlankPages},
		{"RemoveBookmarks", (*Document).RemoveBookmarks},
		{"RemoveHiddenText", (*Document).RemoveHiddenText},
		{"RemoveImages", (*Document).RemoveImages},
		{"RemoveJavaScripts", (*Document).RemoveJavaScripts},
		{"PageRotate", func(doc *Document) error { return doc.PageRotate(1, RotationOn270) }},
		{"PageSetSize", func(doc *Document) error { return doc.PageSetSize(1, PageSizeA1) }},
		{"PageGrayscale", func(doc *Document) error { return doc.PageGrayscale(1) }},
		{"PageAddText", func(doc *Document) error { return doc.PageAddText(1, "Page-level text") }},
		{"PageReplaceText", func(doc *Document) error {
			_ = doc.PageAddText(1, "Replace me")
			return doc.PageReplaceText(1, "Replace", "Changed")
		}},
		{"PageAddPageNum", func(doc *Document) error { return doc.PageAddPageNum(1) }},
		{"PageAddTextHeader", func(doc *Document) error { return doc.PageAddTextHeader(1, "Page Header") }},
		{"PageAddTextFooter", func(doc *Document) error { return doc.PageAddTextFooter(1, "Page Footer") }},
		{"PageRemoveAnnotations", func(doc *Document) error { return doc.PageRemoveAnnotations(1) }},
		{"PageRemoveHiddenText", func(doc *Document) error { return doc.PageRemoveHiddenText(1) }},
		{"PageRemoveImages", func(doc *Document) error { return doc.PageRemoveImages(1) }},
	}

	for _, test := range organizeFunctions {
		t.Run(test.name, func(t *testing.T) {
			doc, err := New()
			if err != nil {
				t.Fatalf("New(): %v", err)
			}
			defer doc.Close()

			if err := doc.PageAdd(); err != nil {
				t.Fatalf("PageAdd(): %v", err)
			}

			if err := test.fn(doc); err != nil {
				t.Errorf("%s(): %v", test.name, err)
			}

			outputPath := fmt.Sprintf("%s/output.pdf", t.TempDir())
			if err := doc.SaveAs(outputPath); err != nil {
				t.Errorf("SaveAs(): %v", err)
			}

			info, err := os.Stat(outputPath)
			if err != nil {
				t.Errorf("Stat(%s): %v", outputPath, err)
			} else {
				assert_ne(t, int64(0), info.Size())
			}
		})
	}
}

func TestRepair(t *testing.T) {

	tmpDir := t.TempDir()
	filePath := fmt.Sprintf("%s/input.pdf", tmpDir)

	doc, err := New()
	if err != nil {
		t.Fatalf("New(): %v", err)
	}
	defer doc.Close()

	if err := doc.PageAdd(); err != nil {
		t.Fatalf("PageAdd(): %v", err)
	}
	if err := doc.SaveAs(filePath); err != nil {
		t.Fatalf("SaveAs(%s): %v", filePath, err)
	}

	reopenDoc, err := Open(filePath)
	if err != nil {
		t.Fatalf("Open(%s): %v", filePath, err)
	}
	defer reopenDoc.Close()

	if err := reopenDoc.Repair(); err != nil {
		t.Errorf("Repair(): %v", err)
	}
}

func TestBytes(t *testing.T) {

	pdf, err := New()
	if err != nil {
		t.Fatalf("New(): %v", err)
	}
	defer pdf.Close()

	data, err := pdf.Bytes()
	if err != nil {
		t.Fatalf("Bytes(): %v", err)
	}

	// Assert that the byte slice is not empty
	assert_ne(t, int64(0), int64(len(data)))
}
