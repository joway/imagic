package util

import (
	"os"
	"path/filepath"
)

func EnsureDir(fn string) error {
	dirName := filepath.Dir(fn)
	if _, err := os.Stat(dirName); err != nil {
		err := os.MkdirAll(dirName, os.ModePerm)
		return err
	}
	return nil
}
