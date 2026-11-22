package asposepdf

// PdfSignStatus contains status of signature in PDF-document.
type PdfSignStatus struct {
	Name   string     `json:"sign_name"`   // name of signature
	Status SignStatus `json:"sign_status"` // status of signature
}
