package util

import (
	"image"
	"image/png"
	"os"
	"path/filepath"
)

// EnsureDir ensure dir is existed
func EnsureDir(fn string) error {
	dirName := filepath.Dir(fn)
	if _, err := os.Stat(dirName); err != nil {
		err := os.MkdirAll(dirName, os.ModePerm)
		return err
	}
	return nil
}

// WriteImage
func WriteImage(filename string, img image.Image) error {
	outputFile, err := os.Create(filename)
	if err != nil {
		return err
	}

	if err := png.Encode(outputFile, img); err != nil {
		return err
	}

	return outputFile.Close()
}
