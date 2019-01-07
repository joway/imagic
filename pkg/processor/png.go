package processor

import (
	"github.com/ultimate-guitar/go-imagequant"
)

type PNGProcessor struct {
	Processor
}

func (c *PNGProcessor) Compress(input []byte, ratio float32) ([]byte, error) {
	output, err := imagequant.Crush(input, 9, 0)
	if err != nil {
		return nil, err
	}
	return output, nil
}
