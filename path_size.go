package code

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"strconv"
)

func GetPathSize(path string, recursive, human, all bool) (string, error) {
	var size int64

	fi, err := os.Lstat(path)
	if err != nil {
		return "", err
	}

	if fi.Mode().IsRegular() {
		size = fi.Size()
	}

	if fi.IsDir() {
		entries, err := os.ReadDir(path)
		if err != nil {
			return "", err
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
				subSize, err := GetPathSize(subPath, recursive, false, all)
				if err != nil {
					continue
				}

				convertedSize, err := strconv.ParseInt(subSize, 10, 64)
				if err != nil {
					continue
				}
				size += convertedSize
			}

			if info.Mode().IsRegular() {
				size += info.Size()
			}
		}
	}
	strSize := fmt.Sprintf("%d", size)
	
	return FormatSize(strSize, human), nil
}

func FormatSize(size string, human bool) string {
	units := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}
	floatSize, err := strconv.ParseFloat(size, 64)
	if err != nil {
		return ""
	}

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
	return fmt.Sprintf("%sB", size)
}
