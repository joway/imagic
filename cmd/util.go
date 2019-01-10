package cmd

import (
	"github.com/joway/imagic/pkg/util"
	"github.com/thoas/go-funk"
	"path/filepath"
)

func getFilesFromPatterns(patterns []string) []string {
	result := make([]string, 0)
	for _, pattern := range patterns {
		files := getFilesFromPattern(pattern)
		result = append(result, files...)
	}
	return result
}

func getFilesFromPattern(pattern string) []string {
	files, err := filepath.Glob(pattern)
	if err != nil {
		util.LogFatal(err)
	}
	return files
}

func getOutputFileName(filename string, suffix string, outputDir string) string {
	var outputPath string
	if outputDir != "" {
		absOutput, _ := filepath.Abs(outputDir)
		outputPath = absOutput
	} else {
		absOutput, _ := filepath.Abs(filepath.Dir(filename))
		outputPath = absOutput
	}
	baseFilename := filepath.Base(filename)
	dotPos := funk.LastIndexOf(baseFilename, ".")
	if dotPos == -1 {
		return outputPath + "/" + baseFilename + suffix
	}
	return outputPath + "/" + baseFilename[:dotPos] + suffix + baseFilename[dotPos:]
}
