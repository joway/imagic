package image

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var png, _ = NewImageFromPath("../../testdata/images/png/3.png")
var jpg, _ = NewImageFromPath("../../testdata/images/jpg/1.jpg")

func TestImage_Compress(t *testing.T) {
	outputPng, err := png.Compress(50)
	assert.NoError(t, err)
	assert.True(t, len(outputPng.Data) < len(png.Data))
	err = outputPng.Write("../../testdata/images/png/1.output.png")
	assert.NoError(t, err)

	outputJpg, err := jpg.Compress(50)
	assert.NoError(t, err)
	assert.True(t, len(outputJpg.Data) < len(jpg.Data))
	err = outputJpg.Write("../../testdata/images/jpg/1.output.jpg")
	assert.NoError(t, err)
}
