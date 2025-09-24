#pragma once

#if defined(_MSC_VER)
    #define ASPOSE_PDF_RUST_SHARED_API __declspec(dllexport)
#elif defined(__GNUC__)
    #define ASPOSE_PDF_RUST_SHARED_API __attribute__((visibility("default")))
#else
    #define ASPOSE_PDF_RUST_SHARED_API
#endif

#if defined(_MSC_VER)
    #define strdup _strdup
#endif

#ifdef __cplusplus
extern "C" {
#endif
    ASPOSE_PDF_RUST_SHARED_API void* PDFDocument_New(const char** error);
    ASPOSE_PDF_RUST_SHARED_API void* PDFDocument_Open(const char* filename, const char** error);
    ASPOSE_PDF_RUST_SHARED_API void PDFDocument_Release(void* pdfdocumentclass, const char** error);
    ASPOSE_PDF_RUST_SHARED_API const char* PDFDocument_About(void* pdfdocumentclass, const char** error);
    ASPOSE_PDF_RUST_SHARED_API void PDFDocument_set_License(void* pdfdocumentclass, const char* filename, const char** error);
    ASPOSE_PDF_RUST_SHARED_API void PDFDocument_Save(void* pdfdocumentclass, const char** error);
    ASPOSE_PDF_RUST_SHARED_API void PDFDocument_Save_As(void* pdfdocumentclass, const char* filename, const char** error);
    ASPOSE_PDF_RUST_SHARED_API const char* PDFDocument_ExtractText(void* pdfdocumentclass, const char** error);
    ASPOSE_PDF_RUST_SHARED_API void PDFDocument_Optimize(void* pdfdocumentclass, const char** error);
    ASPOSE_PDF_RUST_SHARED_API void PDFDocument_Repair(void* pdfdocumentclass, const char** error);
    ASPOSE_PDF_RUST_SHARED_API void PDFDocument_OptimizeResource(void* pdfdocumentclass, const char** error);
    ASPOSE_PDF_RUST_SHARED_API void PDFDocument_Grayscale(void* pdfdocumentclass, const char** error);
    ASPOSE_PDF_RUST_SHARED_API void PDFDocument_Flatten(void* pdfdocumentclass, const char** error);
    ASPOSE_PDF_RUST_SHARED_API void PDFDocument_RemoveAnnotations(void* pdfdocumentclass, const char** error);
    ASPOSE_PDF_RUST_SHARED_API void PDFDocument_RemoveAttachments(void* pdfdocumentclass, const char** error);
    ASPOSE_PDF_RUST_SHARED_API void PDFDocument_RemoveBlankPages(void* pdfdocumentclass, const char** error);
    ASPOSE_PDF_RUST_SHARED_API void PDFDocument_RemoveBookmarks(void* pdfdocumentclass, const char** error);
    ASPOSE_PDF_RUST_SHARED_API void PDFDocument_RemoveHiddenText(void* pdfdocumentclass, const char** error);
    ASPOSE_PDF_RUST_SHARED_API void PDFDocument_RemoveImages(void* pdfdocumentclass, const char** error);
    ASPOSE_PDF_RUST_SHARED_API void PDFDocument_RemoveJavaScripts(void* pdfdocumentclass, const char** error);
    ASPOSE_PDF_RUST_SHARED_API void PDFDocument_set_Background(void* pdfdocumentclass, int r, int g, int b, const char** error);
    ASPOSE_PDF_RUST_SHARED_API void PDFDocument_Rotate(void* pdfdocumentclass, int rotation, const char** error);
    ASPOSE_PDF_RUST_SHARED_API int PDFDocument_get_WordCount(void* pdfdocumentclass, const char** error);
    ASPOSE_PDF_RUST_SHARED_API int PDFDocument_get_CharacterCount(void* pdfdocumentclass, const char** error);
    ASPOSE_PDF_RUST_SHARED_API void PDFDocument_ReplaceText(void* pdfdocumentclass, const char* findText, const char* replaceText, const char** error);
    ASPOSE_PDF_RUST_SHARED_API void PDFDocument_AddPageNum(void* pdfdocumentclass, const char** error);
    ASPOSE_PDF_RUST_SHARED_API void PDFDocument_AddTextHeader(void* pdfdocumentclass, const char* header, const char** error);
    ASPOSE_PDF_RUST_SHARED_API void PDFDocument_AddTextFooter(void* pdfdocumentclass, const char* footer, const char** error);
    ASPOSE_PDF_RUST_SHARED_API void PDFDocument_Save_DocX(void* pdfdocumentclass, const char* filename, const char** error);
    ASPOSE_PDF_RUST_SHARED_API void PDFDocument_Save_Doc(void* pdfdocumentclass, const char* filename, const char** error);
    ASPOSE_PDF_RUST_SHARED_API void PDFDocument_Save_XlsX(void* pdfdocumentclass, const char* filename, const char** error);
    ASPOSE_PDF_RUST_SHARED_API void PDFDocument_Save_PptX(void* pdfdocumentclass, const char* filename, const char** error);
    ASPOSE_PDF_RUST_SHARED_API void PDFDocument_Save_Xps(void* pdfdocumentclass, const char* filename, const char** error);
    ASPOSE_PDF_RUST_SHARED_API void PDFDocument_Save_Txt(void* pdfdocumentclass, const char* filename, const char** error);
    ASPOSE_PDF_RUST_SHARED_API void PDFDocument_Save_Epub(void* pdfdocumentclass, const char* filename, const char** error);
    ASPOSE_PDF_RUST_SHARED_API void PDFDocument_Save_TeX(void* pdfdocumentclass, const char* filename, const char** error);
    ASPOSE_PDF_RUST_SHARED_API void PDFDocument_Save_Markdown(void* pdfdocumentclass, const char* filename, const char** error);
    ASPOSE_PDF_RUST_SHARED_API void PDFDocument_Save_Booklet(void* pdfdocumentclass, const char* filename, const char** error);
    ASPOSE_PDF_RUST_SHARED_API void PDFDocument_Save_NUp(void* pdfdocumentclass, const char* filename, int columns, int rows, const char** error);
    ASPOSE_PDF_RUST_SHARED_API void PDFDocument_Save_Tiff(void* pdfdocumentclass, int resolutionDPI, const char* filename, const char** error);
    ASPOSE_PDF_RUST_SHARED_API void PDFDocument_Save_DocXEnhanced(void* pdfdocumentclass, const char* filename, const char** error);
    ASPOSE_PDF_RUST_SHARED_API void PDFDocument_Save_SvgZip(void* pdfdocumentclass, const char* filename, const char** error);
    ASPOSE_PDF_RUST_SHARED_API void PDFDocument_Export_Fdf(void* pdfdocumentclass, const char* filename, const char** error);
    ASPOSE_PDF_RUST_SHARED_API void PDFDocument_Export_Xfdf(void* pdfdocumentclass, const char* filename, const char** error);
    ASPOSE_PDF_RUST_SHARED_API void PDFDocument_Export_Xml(void* pdfdocumentclass, const char* filename, const char** error);
    ASPOSE_PDF_RUST_SHARED_API void PDFDocument_Append(void* pdfdocumentclass, const void* otherpdfdocumentclass, const char** error);
    ASPOSE_PDF_RUST_SHARED_API void PDFDocument_AppendPages(void* pdfdocumentclass, const void* otherpdfdocumentclass, const char* pagerange, const char** error);
    ASPOSE_PDF_RUST_SHARED_API void PDFDocument_Save_Memory(void* pdfdocumentclass, unsigned char** bufferOut, int* sizeOut, const char** error);
    ASPOSE_PDF_RUST_SHARED_API int PDFDocument_Page_get_Count(void* pdfdocumentclass, const char** error);
    ASPOSE_PDF_RUST_SHARED_API void PDFDocument_Page_Add(void* pdfdocumentclass, const char** error);
    ASPOSE_PDF_RUST_SHARED_API void PDFDocument_Page_Insert(void* pdfdocumentclass, int num, const char** error);
    ASPOSE_PDF_RUST_SHARED_API void PDFDocument_Page_Delete(void* pdfdocumentclass, int num, const char** error);
    ASPOSE_PDF_RUST_SHARED_API void PDFDocument_Page_to_Jpg(void* pdfdocumentclass, int num, int resolutionDPI, const char* filename, const char** error);
    ASPOSE_PDF_RUST_SHARED_API void PDFDocument_Page_to_Png(void* pdfdocumentclass, int num, int resolutionDPI, const char* filename, const char** error);
    ASPOSE_PDF_RUST_SHARED_API void PDFDocument_Page_to_Bmp(void* pdfdocumentclass, int num, int resolutionDPI, const char* filename, const char** error);
    ASPOSE_PDF_RUST_SHARED_API void PDFDocument_Page_to_Tiff(void* pdfdocumentclass, int num, int resolutionDPI, const char* filename, const char** error);
    ASPOSE_PDF_RUST_SHARED_API void PDFDocument_Page_to_DICOM(void* pdfdocumentclass, int num, int resolutionDPI, const char* filename, const char** error);
    ASPOSE_PDF_RUST_SHARED_API void PDFDocument_Page_to_Svg(void* pdfdocumentclass, int num, const char* filename, const char** error);
    ASPOSE_PDF_RUST_SHARED_API void PDFDocument_Page_to_Pdf(void* pdfdocumentclass, int num, const char* filename, const char** error);
    ASPOSE_PDF_RUST_SHARED_API void PDFDocument_Page_Grayscale(void* pdfdocumentclass, int num, const char** error);
    ASPOSE_PDF_RUST_SHARED_API void PDFDocument_Page_Rotate(void* pdfdocumentclass, int num, int rotation, const char** error);
    ASPOSE_PDF_RUST_SHARED_API void PDFDocument_Page_ReplaceText(void* pdfdocumentclass, int num, const char* findText, const char* replaceText, const char** error);
    ASPOSE_PDF_RUST_SHARED_API void PDFDocument_Page_AddText(void* pdfdocumentclass, int num, const char* addText, const char** error);
    ASPOSE_PDF_RUST_SHARED_API void PDFDocument_Page_set_Size(void* pdfdocumentclass, int num, int pageSize, const char** error);
    ASPOSE_PDF_RUST_SHARED_API int PDFDocument_Page_get_WordCount(void* pdfdocumentclass, int num, const char** error);
    ASPOSE_PDF_RUST_SHARED_API int PDFDocument_Page_get_CharacterCount(void* pdfdocumentclass, int num, const char** error);
    ASPOSE_PDF_RUST_SHARED_API int PDFDocument_Page_is_Blank(void* pdfdocumentclass, int num, const char** error);
    ASPOSE_PDF_RUST_SHARED_API void PDFDocument_Page_AddPageNum(void* pdfdocumentclass, int num, const char** error);
    ASPOSE_PDF_RUST_SHARED_API void PDFDocument_Page_AddTextHeader(void* pdfdocumentclass, int num, const char* header, const char** error);
    ASPOSE_PDF_RUST_SHARED_API void PDFDocument_Page_AddTextFooter(void* pdfdocumentclass, int num, const char* footer, const char** error);
    ASPOSE_PDF_RUST_SHARED_API void PDFDocument_Page_RemoveAnnotations(void* pdfdocumentclass, int num, const char** error);
    ASPOSE_PDF_RUST_SHARED_API void PDFDocument_Page_RemoveHiddenText(void* pdfdocumentclass, int num, const char** error);
    ASPOSE_PDF_RUST_SHARED_API void PDFDocument_Page_RemoveImages(void* pdfdocumentclass, int num, const char** error);

#ifdef __cplusplus
}
#endif

#ifdef __cplusplus
extern "C" {
#endif
    ASPOSE_PDF_RUST_SHARED_API void c_free_string(char* str);
    ASPOSE_PDF_RUST_SHARED_API void c_free_buffer(void* buffer);
#ifdef __cplusplus
}
#endif
