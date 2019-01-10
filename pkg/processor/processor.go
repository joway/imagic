package processor

import (
	"github.com/joway/imagic/pkg/constant"
)

// Engine is the interface of Processor
type Engine interface {
	Compress(input []byte, quality int) ([]byte, error)
	Resize(input []byte, width int, height int) ([]byte, error)
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
