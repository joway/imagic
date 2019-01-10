package image

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var textureImage, _ = NewImageFromPath("../../testdata/images/texture.png")
var pngImage, _ = NewImageFromPath("../../testdata/images/png/1.png")
var jpgImage, _ = NewImageFromPath("../../testdata/images/jpg/1.jpg")

func TestImage_Compress(t *testing.T) {
	assert.NotNil(t, textureImage)
	assert.NotNil(t, pngImage)
	assert.NotNil(t, jpgImage)

	// PNG
	// compress
	outputImg, err := pngImage.Compress(50)
	assert.NoError(t, err)
	err = outputImg.Write("../../output/1.output.png")
	assert.NoError(t, err)
	// resize
	outputImg, err = pngImage.Resize(320, 0)
	assert.NoError(t, err)
	err = outputImg.Write("../../output/1.output.png")
	assert.NoError(t, err)
	// watermark
	outputImg, err = pngImage.WaterMark(textureImage, "30", "30")
	assert.NoError(t, err)
	err = outputImg.Write("../../output/1.output.png")
	assert.NoError(t, err)
	outputImg, err = pngImage.WaterMark(textureImage, "+30", "+30")
	assert.NoError(t, err)
	err = outputImg.Write("../../output/1.output.png")
	assert.NoError(t, err)
	outputImg, err = pngImage.WaterMark(textureImage, "-200", "-100")
	assert.NoError(t, err)
	err = outputImg.Write("../../output/1.output.png")
	assert.NoError(t, err)

	// JPEG
	// compress
	outputImg, err = jpgImage.Compress(50)
	assert.NoError(t, err)
	err = outputImg.Write("../../output/1.output.jpg")
	assert.NoError(t, err)
	// resize
	outputImg, err = jpgImage.Resize(1000, 0)
	assert.NoError(t, err)
	err = outputImg.Write("../../output/1.output.jpg")
	assert.NoError(t, err)
	// watermark
	outputImg, err = outputImg.WaterMark(textureImage, "-600", "-100")
	assert.NoError(t, err)
	err = outputImg.Write("../../output/1.output.jpg")
	assert.NoError(t, err)
}
