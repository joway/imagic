package processor

import (
	"bufio"
	"bytes"
	"github.com/disintegration/imaging"
	"image/jpeg"
)

type JPEGProcessor struct {
	Processor
}

func (p *JPEGProcessor) Compress(input []byte, quality int) ([]byte, error) {
	buf := new(bytes.Buffer)
	output := bufio.NewWriter(buf)
	img, err := jpeg.Decode(bytes.NewReader(input))
	if err != nil {
		return nil, err
	}
	err = jpeg.Encode(output, img, &jpeg.Options{Quality: quality})
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (p *JPEGProcessor) Resize(input []byte, width int, height int) ([]byte, error) {
	img, err := jpeg.Decode(bytes.NewReader(input))
	if err != nil {
		return nil, err
	}
	outputImg := imaging.Resize(img, width, height, imaging.Lanczos)
	buf := new(bytes.Buffer)
	if err := jpeg.Encode(buf, outputImg, nil); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
