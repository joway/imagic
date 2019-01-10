package processor

import (
	"github.com/joway/imagic/pkg/constant"
	"image"
)

// Engine is the interface of Processor
type Engine interface {
	Compress(input image.Image, quality int) (image.Image, error)
	Resize(input image.Image, width int, height int) (image.Image, error)
	WaterMark(input image.Image, texture image.Image, x int, y int) (image.Image, error)
}

// Processor is the abstract struct of processor
type Processor struct {
	Engine
	Format string
}

// NewProcessor return the Processor instance
func NewProcessor(format string) Engine {
	var proc = Processor{
		Format: format,
	}
	switch format {
	case constant.FormatPng:
		return &PNGProcessor{proc}
	case constant.FormatJpeg:
		return &JPEGProcessor{proc}
	default:
		return nil
	}
}
