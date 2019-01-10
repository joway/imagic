package cmd

import (
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"testing"
)

func Test_getFilesFromPatterns(t *testing.T) {
	ps := []string{"../testdata/images/jpg/*.jpg", "../testdata/images/png/*.png"}
	files := getFilesFromPatterns(ps)
	assert.Equal(t, len(files), 10)
}

func Test_getOutputFileName(t *testing.T) {
	fn := getOutputFileName("../testdata/images/jpg/1.jpg", ".suf", "./output")
	absOutput, _ := filepath.Abs("./output/1.suf.jpg")
	assert.Equal(t, fn, absOutput)

	fn = getOutputFileName("../testdata/images/jpg/1.jpg", "", "./output")
	absOutput, _ = filepath.Abs("./output/1.jpg")
	assert.Equal(t, fn, absOutput)

	fn = getOutputFileName("../testdata/images/jpg/1.jpg", "", "")
	absOutput, _ = filepath.Abs("../testdata/images/jpg/1.jpg")
	assert.Equal(t, fn, absOutput)

	fn = getOutputFileName("../testdata/images/jpg/1", "", "")
	absOutput, _ = filepath.Abs("../testdata/images/jpg/1")
	assert.Equal(t, fn, absOutput)

	fn = getOutputFileName("../testdata/images/jpg/1", ".suf", "")
	absOutput, _ = filepath.Abs("../testdata/images/jpg/1.suf")
	assert.Equal(t, fn, absOutput)
}
