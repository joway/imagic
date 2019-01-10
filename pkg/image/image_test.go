package image

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var png, _ = NewImageFromPath("../../testdata/images/png/1.png")
var jpg, _ = NewImageFromPath("../../testdata/images/jpg/1.jpg")

func TestImage_Compress(t *testing.T) {
	// PNG
	// compress
	outputPng, err := png.Compress(50)
	assert.NoError(t, err)
	assert.True(t, len(outputPng.Data) < len(png.Data))
	err = outputPng.Write("../../output/1.output.png")
	assert.NoError(t, err)
	// resize
	outputPng, err = png.Resize(320, 0)
	assert.NoError(t, err)
	err = outputPng.Write("../../output/1.output.png")
	assert.NoError(t, err)

	// JPEG
	// compress
	outputJpg, err := jpg.Compress(50)
	assert.NoError(t, err)
	assert.True(t, len(outputJpg.Data) < len(jpg.Data))
	err = outputJpg.Write("../../output/1.output.jpg")
	assert.NoError(t, err)
	// resize
	outputJpg, err = jpg.Resize(320, 0)
	assert.NoError(t, err)
	err = outputJpg.Write("../../output/1.output.jpg")
	assert.NoError(t, err)
}
