package processor

import (
	"bytes"
	"github.com/disintegration/imaging"
	"github.com/joway/libimagequant-go/pngquant"
	"image/png"
)

type PNGProcessor struct {
	Processor
}

func (p *PNGProcessor) Compress(input []byte, quality int) ([]byte, error) {
	decoded, err := png.Decode(bytes.NewReader(input))
	if err != nil {
		return nil, err
	}
	outputImg, err := pngquant.Compress(decoded, quality, pngquant.SPEED_DEFAULT)
	if err != nil {
		return nil, err
	}
	buf := new(bytes.Buffer)
	if err := png.Encode(buf, outputImg); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (p *PNGProcessor) Resize(input []byte, width int, height int) ([]byte, error) {
	img, err := png.Decode(bytes.NewReader(input))
	if err != nil {
		return nil, err
	}
	outputImg := imaging.Resize(img, width, height, imaging.Lanczos)
	buf := new(bytes.Buffer)
	if err := png.Encode(buf, outputImg); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
