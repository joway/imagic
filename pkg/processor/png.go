package processor

import (
	"github.com/disintegration/imaging"
	"github.com/joway/libimagequant-go/pngquant"
	"github.com/noelyahan/mergi"
	"image"
)

// PNGProcessor is the processor for PNG images
type PNGProcessor struct {
	Processor
}

// Compress the images
func (p *PNGProcessor) Compress(input image.Image, quality int) (image.Image, error) {
	outputImg, err := pngquant.Compress(input, quality, pngquant.SPEED_FASTEST)
	if err != nil {
		return nil, err
	}
	return outputImg, nil
}

// Resize the images
func (p *PNGProcessor) Resize(input image.Image, width int, height int) (image.Image, error) {
	outputImg := imaging.Resize(input, width, height, imaging.Lanczos)
	return outputImg, nil
}

// WaterMark the images
func (p *PNGProcessor) WaterMark(input image.Image, texture image.Image, x int, y int) (image.Image, error) {
	outputImg, err := mergi.Watermark(texture, input, image.Pt(x, y))
	if err != nil {
		return nil, err
	}
	return outputImg, nil
}
