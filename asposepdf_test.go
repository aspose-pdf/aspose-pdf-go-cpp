// use:
// go test -v

package asposepdf

import (
	"fmt"
	"os"
	"reflect"
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
