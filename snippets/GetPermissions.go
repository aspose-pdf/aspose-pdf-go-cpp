package main

import "github.com/aspose-pdf/aspose-pdf-go-cpp"
import (
	"fmt"
	"log"
	"strings"
)

var permissionNames = map[asposepdf.Permissions]string{
	asposepdf.PrintDocument:                  "Allow printing",
	asposepdf.ModifyContent:                  "Allow modifying content (except forms/annotations)",
	asposepdf.ExtractContent:                 "Allow copying/extracting text and graphics",
	asposepdf.ModifyTextAnnotations:          "Allow adding/modifying text annotations",
	asposepdf.FillForm:                       "Allow filling interactive forms",
	asposepdf.ExtractContentWithDisabilities: "Allow content extraction for accessibility",
	asposepdf.AssembleDocument:               "Allow inserting/rotating/deleting pages or changing structure",
	asposepdf.PrintingQuality:                "Allow high-quality / faithful printing",
}

func PermissionsToString(p asposepdf.Permissions) string {
	var result []string
	for flag, name := range permissionNames {
		if p&flag != 0 {
			result = append(result, name)
		}
	}
	if len(result) == 0 {
		return "None"
	}
	return strings.Join(result, ", ")
}

func main() {
	// OpenWithPassword(filename string, password string) opens a password-protected PDF-document
	pdf, err := asposepdf.OpenWithPassword("sample_with_permissions.pdf", "ownerpass")
	if err != nil {
		log.Fatal(err)
	}
	// Close() releases allocated resources for PDF-document
	defer pdf.Close()
	// GetPermissions() gets current permissions of PDF-document
	permissions, err := pdf.GetPermissions()
	if err != nil {
		log.Fatal(err)
	}
	// Print permissions
	fmt.Println("Permissions:", PermissionsToString(permissions))
}
