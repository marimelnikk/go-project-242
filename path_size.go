package pathsize

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func GetPathSize(path string, all, recursive bool) (int64, error) {
	var size int64

	fi, err := os.Lstat(path)
	if err != nil {
		return 0, err
	}

	if fi.Mode().IsRegular() {
		size = fi.Size()
	}

	if fi.IsDir() {
		entries, err := os.ReadDir(path)
		if err != nil {
			return 0, err
		}

		for _, entry := range entries {
			if !all && strings.HasPrefix(entry.Name(), ".") {
				continue
			}

			info, err := entry.Info()
			if err != nil {
				continue
			}

			if recursive && info.Mode().IsDir() {
				subPath := filepath.Join(path, entry.Name())
				subSize, err := GetPathSize(subPath, all, recursive)
				if err != nil {
					continue
				}
				size += subSize
			}

			if info.Mode().IsRegular() {
				size += info.Size()
			}
		}
	}
	return size, nil
}

func FormatSize(size int64, human bool) string {
	units := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}
	floatSize := float64(size)

	if human {
		for i, u := range units {
			if floatSize < 1024 {
				if i == 0 {
					return fmt.Sprintf("%d%s", int(floatSize), u)
				}
				return fmt.Sprintf("%.1f%s", floatSize, u)
			}
			floatSize /= 1024
		}
	}
	return fmt.Sprintf("%dB", size)
}
