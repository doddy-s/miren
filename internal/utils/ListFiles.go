package utils

import (
	"os"
	"path/filepath"
)

type MirenDirEntry struct {
	os.DirEntry
	FullPath string
}

// ListFiles recursively finds matching files only,
// and only descends into directories that match the pattern (if given).
func ListFiles(dir string, pattern string) ([]MirenDirEntry, error) {
	var result []MirenDirEntry

	err := walk(dir, pattern, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func walk(currentPath string, pattern string, result *[]MirenDirEntry) error {
	entries, err := os.ReadDir(currentPath)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		name := entry.Name()
		fullPath := filepath.Join(currentPath, name)

		// Apply pattern filter to both files and directories
		if pattern != "" {
			match, err := filepath.Match(pattern, name)
			if err != nil {
				return err
			}
			if !match {
				continue
			}
		}

		if entry.IsDir() {
			// Only enter matching directories
			err := walk(fullPath, pattern, result)
			if err != nil {
				return err
			}
		} else {
			// Only collect files
			*result = append(*result, MirenDirEntry{
				DirEntry: entry,
				FullPath: fullPath,
			})
		}
	}

	return nil
}
