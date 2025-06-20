package utils

import "fmt"

func FormatBytes(size int64) string {
	const (
		_           = iota
		KiB float64 = 1 << (10 * iota)
		MiB
		GiB
		TiB
	)

	s := float64(size)

	switch {
	case s >= TiB:
		return fmt.Sprintf("%.2f TiB", s/TiB)
	case s >= GiB:
		return fmt.Sprintf("%.2f GiB", s/GiB)
	case s >= MiB:
		return fmt.Sprintf("%.2f MiB", s/MiB)
	case s >= KiB:
		return fmt.Sprintf("%.2f KiB", s/KiB)
	default:
		return fmt.Sprintf("%d B", size)
	}
}
