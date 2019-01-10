package processor

import (
	"bufio"
	"bytes"
	"github.com/disintegration/imaging"
	"github.com/noelyahan/mergi"
	"image"
	"image/jpeg"
)

// JPEGProcessor is the processor for JPEG images
type JPEGProcessor struct {
	Processor
}

// Compress the images
func (p *JPEGProcessor) Compress(input image.Image, quality int) (image.Image, error) {
	buf := new(bytes.Buffer)
	output := bufio.NewWriter(buf)
	if err := jpeg.Encode(output, input, &jpeg.Options{Quality: quality}); err != nil {
		return nil, err
	}
	outputImg, err := jpeg.Decode(buf)
	if err != nil {
		return nil, err
	}
	return outputImg, nil
}

// Resize the images
func (p *JPEGProcessor) Resize(input image.Image, width int, height int) (image.Image, error) {
	outputImg := imaging.Resize(input, width, height, imaging.Lanczos)
	return outputImg, nil
}

// WaterMark the images
func (p *JPEGProcessor) WaterMark(input image.Image, texture image.Image, x int, y int) (image.Image, error) {
	outputImg, err := mergi.Watermark(texture, input, image.Pt(x, y))
	if err != nil {
		return nil, err
	}
	return outputImg, nil
}
