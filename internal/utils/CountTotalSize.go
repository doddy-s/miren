package utils

// CountTotalSize returns the total size in bytes of all files in the slice.
func CountTotalSize(entries []MirenDirEntry) (int64, error) {
	var total int64 = 0

	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			return 0, err
		}

		// Only add size if it's a regular file (optional check)
		if info.Mode().IsRegular() {
			total += info.Size()
		}
	}

	return total, nil
}
