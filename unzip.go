// To ensure package consistency, the following directive is required:

//go:build ignore

// This program downloads and decompresses dynamic libraries.
// It can be invoked by running `go generate`.

package main

import (
	"compress/bzip2"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

const verLibsTag = "v1.25.5"
const repoLibsURL = "https://raw.githubusercontent.com/aspose-pdf/aspose-pdf-go-cpp-libs/"
const repoLibsDir = "libs/"

var fileMap = map[string]string{
	"darwin_amd64":  "libAsposePDFforGo_darwin_amd64.dylib.bz2",
	"darwin_arm64":  "libAsposePDFforGo_darwin_arm64.dylib.bz2",
	"linux_amd64":   "libAsposePDFforGo_linux_amd64.so.bz2",
	"windows_amd64": "AsposePDFforGo_windows_amd64.dll.bz2",
}

func downloadFile(url, destPath string) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to download %s: %w", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to download %s: HTTP %d", url, resp.StatusCode)
	}

	outFile, err := os.Create(destPath)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %w", destPath, err)
	}
	defer outFile.Close()

	totalSize := resp.ContentLength
	buffer := make([]byte, 1024)
	var downloaded int64
	for {
		n, err := resp.Body.Read(buffer)
		if n > 0 {
			if _, err := outFile.Write(buffer[:n]); err != nil {
				return err
			}
			downloaded += int64(n)

			if totalSize > 0 {
				fmt.Printf("\rDownloading %s: %d%%", destPath, downloaded*100/totalSize)
			} else {
				fmt.Printf("\rDownloading %s: %db", destPath, downloaded)
			}

		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
	}
	fmt.Println(" Done.")
	return nil
}

func verifySHA256(filePath, sha256Path string) error {
	shaBytes, err := os.ReadFile(sha256Path)
	if err != nil {
		return fmt.Errorf("failed to read SHA256 file: %w", err)
	}
	expectedSHA := strings.TrimSpace(string(shaBytes))

	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file for hashing: %w", err)
	}
	defer file.Close()

	hasher := sha256.New()
	if _, err := io.Copy(hasher, file); err != nil {
		return fmt.Errorf("failed to compute hash: %w", err)
	}

	computedSHA := hex.EncodeToString(hasher.Sum(nil))
	if computedSHA != expectedSHA {
		return fmt.Errorf("SHA256 mismatch: expected %s, got %s", expectedSHA, computedSHA)
	}

	fmt.Println("SHA256 verification passed.")
	return nil
}

func extractBZ2(archivePath, outputPath string) error {
	file, err := os.Open(archivePath)
	if err != nil {
		return fmt.Errorf("error opening file %s: %w", archivePath, err)
	}
	defer file.Close()

	bz2Reader := bzip2.NewReader(file)
	outFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("error creating output file %s: %w", outputPath, err)
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, bz2Reader)
	if err != nil {
		return err
	}

	fmt.Println("Extraction completed successfully.")
	return nil
}

func main() {
	fmt.Println("`go generate` stared.")

	platform := fmt.Sprintf("%s_%s", runtime.GOOS, runtime.GOARCH)
	archiveFileName, ok := fileMap[platform]
	if !ok {
		fmt.Printf("Unsupported platform: %s\n", platform)
		return
	}

	libDir := "lib"
	archiveFilePath := filepath.Join(libDir, archiveFileName)
	outputFilePath := archiveFilePath[:len(archiveFilePath)-4]
	sha256FilePath := archiveFilePath + ".sha256"

	os.MkdirAll(libDir, os.ModePerm)

	if _, err := os.Stat(archiveFilePath); os.IsNotExist(err) {
		url := repoLibsURL + verLibsTag + "/" + repoLibsDir + archiveFileName
		fmt.Printf("Downloading %s...\n", archiveFileName)
		if err := downloadFile(url, archiveFilePath); err != nil {
			fmt.Println(err)
			return
		}
	}

	if _, err := os.Stat(sha256FilePath); os.IsNotExist(err) {
		shaURL := repoLibsURL + verLibsTag + "/" + repoLibsDir + archiveFileName + ".sha256"
		fmt.Printf("Downloading SHA256 file...\n")
		if err := downloadFile(shaURL, sha256FilePath); err != nil {
			fmt.Println("Warning: SHA256 file not found, skipping verification.")
		}
	}

	if _, err := os.Stat(sha256FilePath); err == nil {
		if err := verifySHA256(archiveFilePath, sha256FilePath); err != nil {
			fmt.Println("Warning:", err)
		}
	}

	fmt.Printf("Decompressing %s...\n", archiveFilePath)
	if err := extractBZ2(archiveFilePath, outputFilePath); err != nil {
		fmt.Println(err)
	}

	fmt.Println("`go generate` completed successfully.")
}
