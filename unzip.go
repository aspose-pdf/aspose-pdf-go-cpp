// To ensure package consistency, the following directive is required:

//go:build ignore

// This program decompresses dynamic libraries. It can be invoked by running
// go generate

package main

import (
	"compress/bzip2"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
)

func main() {

	// Determine the current platform
	platform := fmt.Sprintf("%s_%s", runtime.GOOS, runtime.GOARCH)

	// Map the platform to the corresponding file name
	fileMap := map[string]string{
		"darwin_amd64":  "libAsposePDFforGo_darwin_amd64.dylib.bz2",
		"darwin_arm64":  "libAsposePDFforGo_darwin_arm64.dylib.bz2",
		"linux_amd64":   "libAsposePDFforGo_linux_amd64.so.bz2",
		"windows_amd64": "AsposePDFforGo_windows_amd64.dll.bz2",
	}

	// Get the file name for the current platform
	archiveFileName, ok := fileMap[platform]
	if !ok {
		fmt.Printf("Unsupported platform: %s\n", platform)
		return
	}

	// Define the directory for archives and extracted files
	libDir := "lib"
	archiveFilePath := filepath.Join(libDir, archiveFileName)

	// Output informational message
	fmt.Printf("Decompressing file: %s\n", archiveFilePath)

	// Open the .bz2 file
	file, err := os.Open(archiveFilePath)
	if err != nil {
		fmt.Printf("Error opening file %s: %v\n", archiveFilePath, err)
		return
	}
	defer file.Close()

	// Create a bzip2 reader
	bz2Reader := bzip2.NewReader(file)

	// Determine the output file name based on the archive file name
	outputFileName := archiveFileName[:len(archiveFileName)-4] // remove ".bz2" suffix
	outputFilePath := filepath.Join(libDir, outputFileName)

	// Create the output file
	outFile, err := os.Create(outputFilePath)
	if err != nil {
		fmt.Printf("Error creating output file %s: %v\n", outputFilePath, err)
		return
	}
	defer outFile.Close()

	// Copy the decompressed data to the output file
	_, err = io.Copy(outFile, bz2Reader)
	if err != nil {
		fmt.Printf("Error decompressing file %s: %v\n", archiveFilePath, err)
		return
	}

	fmt.Printf("File decompressed successfully to %s\n", outputFilePath)
}
