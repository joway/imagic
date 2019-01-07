package processor

import "github.com/joway/imagic/pkg/constant"

type Engine interface {
	Compress(input []byte, ratio float32) ([]byte, error)
}

type Processor struct {
	Engine
	Format string
}

type Option struct {
	Ratio  float32
	Width  uint32
	Height uint32
}

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
