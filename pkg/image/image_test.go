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
	outputImg, err := pngImage.Compress(70)
	assert.NoError(t, err)
	err = outputImg.Write("../../output/1.output.png")
	assert.NoError(t, err)

	// JPEG
	// compress
	outputImg, err = jpgImage.Compress(70)
	assert.NoError(t, err)
	err = outputImg.Write("../../output/1.output.jpg")
	assert.NoError(t, err)
}

func TestImage_Resize(t *testing.T) {
	// resize png
	outputImg, err := pngImage.Resize(320, 0)
	assert.NoError(t, err)
	err = outputImg.Write("../../output/1.output.png")
	assert.NoError(t, err)

	// resize jpeg
	outputImg, err = jpgImage.Resize(1000, 0)
	assert.NoError(t, err)
	err = outputImg.Write("../../output/1.output.jpg")
	assert.NoError(t, err)
}

func TestImage_WaterMark(t *testing.T) {
	// watermark png
	outputImg, err := pngImage.WaterMark(textureImage, "30", "30")
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

	// watermark jpeg (center-bottom)
	outputImg, err = jpgImage.Resize(1000, 0)
	assert.NoError(t, err)
	err = outputImg.Write("../../output/1.output.jpg")
	assert.NoError(t, err)
	outputImg, err = outputImg.WaterMark(textureImage, "-600", "-100")
	assert.NoError(t, err)
	err = outputImg.Write("../../output/1.output.jpg")
	assert.NoError(t, err)
}

func TestNewImageFromPath(t *testing.T) {
	nilImage, _ := NewImageFromPath("xxx.png")
	assert.Nil(t, nilImage)
}

func Test_parseXY(t *testing.T) {
	x, _ := parseXY(1000, "+100")
	assert.Equal(t, 100, x)

	x, _ = parseXY(1000, "-100")
	assert.Equal(t, 900, x)

	x, _ = parseXY(1000, "100")
	assert.Equal(t, 100, x)
}
