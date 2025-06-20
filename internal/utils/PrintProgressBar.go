package utils

import (
	"fmt"
	"strings"
)

// PrintProgressBar prints a dynamic progress bar to the terminal
func PrintProgressBar(current, total int, filename string) {
	progress := float64(current) / float64(total)
	barWidth := 40
	filled := int(progress * float64(barWidth))
	empty := barWidth - filled
	bar := strings.Repeat("=", filled) + strings.Repeat(" ", empty)

	// Truncate filename if too long for neatness
	maxNameLen := 30
	if len(filename) > maxNameLen {
		filename = "..." + filename[len(filename)-maxNameLen+3:]
	}

	fmt.Printf("\rZipping files: [%s] %3d%% (%d/%d) %s", bar, int(progress*100), current, total, filename)
	if current == total {
		fmt.Println("")
	}
}
