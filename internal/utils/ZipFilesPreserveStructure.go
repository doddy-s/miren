package utils

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

// ProgressFunc defines a callback for zip progress updates
type ProgressFunc func(current, total int, filename string)

// ZipFilesPreserveStructure creates a zip file at zipPath containing the given files,
// preserving their paths relative to baseDir.
func ZipFilesPreserveStructure(zipPath string, files []MirenDirEntry, baseDir string, callback ProgressFunc) error {
	// Create the zip file
	zipFile, err := os.Create(zipPath)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	total := len(files)
	for i, file := range files {
		// Calculate the relative path inside the zip
		relPath, err := filepath.Rel(baseDir, file.FullPath)
		if err != nil {
			return err
		}
		relPath = filepath.ToSlash(relPath) // Normalize for ZIP format

		// Open source file
		file, err := os.Open(file.FullPath)
		if err != nil {
			return err
		}

		info, err := file.Stat()
		if err != nil {
			file.Close()
			return err
		}

		// Create zip header
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			file.Close()
			return err
		}
		header.Name = relPath
		header.Method = zip.Deflate

		// Create writer in zip
		writer, err := zipWriter.CreateHeader(header)
		if err != nil {
			file.Close()
			return err
		}

		// Copy file contents to zip
		_, err = io.Copy(writer, file)
		file.Close()
		if err != nil {
			return err
		}

		if callback != nil {
			callback(i+1, total, file.Name())
		}
	}

	return nil
}
