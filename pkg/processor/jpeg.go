package processor

import (
	"bufio"
	"bytes"
	"image/jpeg"
)

type JPEGProcessor struct {
	Processor
}

func (c *JPEGProcessor) Compress(input []byte, ratio float32) ([]byte, error) {
	buf := new(bytes.Buffer)
	output := bufio.NewWriter(buf)
	img, err := jpeg.Decode(bytes.NewReader(input))
	if err != nil {
		return nil, err
	}
	err = jpeg.Encode(output, img, &jpeg.Options{Quality: int(ratio * 100)})
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
