package asposepdf

// Meaning no error.
const (
	ERR_OK = ""
)

// Enumeration of possible rotation values.
const (
	RotationNone  int32 = 0 // Non-rotated.
	RotationOn90  int32 = 1 // Rotated on 90 degrees clockwise.
	RotationOn180 int32 = 2 // Rotated on 180 degrees.
	RotationOn270 int32 = 3 // Rotated on 270 degrees clockwise.
	RotationOn360 int32 = 4 // Rotated on 360 degrees clockwise.
)

// Enumeration of possible page size values.
const (
	PageSizeA0         int32 = 0  // A0 size.
	PageSizeA1         int32 = 1  // A1 size.
	PageSizeA2         int32 = 2  // A2 size.
	PageSizeA3         int32 = 3  // A3 size.
	PageSizeA4         int32 = 4  // A4 size.
	PageSizeA5         int32 = 5  // A5 size.
	PageSizeA6         int32 = 6  // A6 size.
	PageSizeB5         int32 = 7  // B5 size.
	PageSizePageLetter int32 = 8  // PageLetter size.
	PageSizePageLegal  int32 = 9  // PageLegal size.
	PageSizePageLedger int32 = 10 // PageLedger size.
	PageSizeP11x17     int32 = 11 // P11x17 size.
)

// Enumeration of possible crypto algorithms.
type CryptoAlgorithm int32

const (
	RC4x40  CryptoAlgorithm = 0 // RC4 with key length 40.
	RC4x128 CryptoAlgorithm = 1 // RC4 with key length 128.
	AESx128 CryptoAlgorithm = 2 // AES with key length 128.
	AESx256 CryptoAlgorithm = 3 // AES with key length 256.
)

// Enumeration of possible PDF formats.
type PdfFormat int32

const (
	PDF_A_1A      PdfFormat = iota // Pdf/A-1a format.
	PDF_A_1B                       // Pdf/A-1b format.
	PDF_A_2A                       // Pdf/A-2a format.
	PDF_A_3A                       // Pdf/A-3a format.
	PDF_A_2B                       // Pdf/A-2b format.
	PDF_A_2U                       // Pdf/A-2u format.
	PDF_A_3B                       // Pdf/A-3b format.
	PDF_A_3U                       // Pdf/A-3u format.
	V_1_0                          // Adobe version 1.0.
	V_1_1                          // Adobe version 1.1.
	V_1_2                          // Adobe version 1.2.
	V_1_3                          // Adobe version 1.3.
	V_1_4                          // Adobe version 1.4.
	V_1_5                          // Adobe version 1.5.
	V_1_6                          // Adobe version 1.6.
	V_1_7                          // Adobe version 1.7.
	V_2_0                          // ISO Standard PDF 2.0.
	PDF_UA_1                       // PDF/UA-1 format.
	PDF_X_1A_2001                  // PDF/X-1a-2001 format.
	PDF_X_1A                       // PDF/X-1a format.
	PDF_X_3                        // PDF/X-3 format.
	ZUGFeRD                        // ZUGFeRD format.
	PDF_A_4                        // PDF/A-4 format.
	PDF_A_4E                       // PDF/A-4e format.
	PDF_A_4F                       // PDF/A-4f format.
	PDF_X_4                        // PDF/X-4 format.
	PDF_E_1                        // PDF/E-1 (PDF 1.6) format.
)

// Enumeration of possible conversion errors action.
type ConvertErrorAction int32

const (
	Delete ConvertErrorAction = iota // Delete convert errors.
	None                             // Do nothing with convert errors.
)

// Bitflag set representing PDF permission capabilities.
type Permissions int32

const (
	PrintDocument                  Permissions = 1 << 2  // 4
	ModifyContent                  Permissions = 1 << 3  // 8
	ExtractContent                 Permissions = 1 << 4  // 16
	ModifyTextAnnotations          Permissions = 1 << 5  // 32
	FillForm                       Permissions = 1 << 8  // 256
	ExtractContentWithDisabilities Permissions = 1 << 9  // 512
	AssembleDocument               Permissions = 1 << 10 // 1024
	PrintingQuality                Permissions = 1 << 11 // 2048
)
