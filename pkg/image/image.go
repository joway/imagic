package image

import (
	"github.com/joway/imagic/pkg/constant"
	"github.com/joway/imagic/pkg/processor"
	"github.com/joway/imagic/pkg/util"
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
	"path/filepath"
)

var SUPPORTED_FORMATS = []string{constant.FormatPng, constant.FormatJpeg}

type Image struct {
	Data      []byte
	Format    string
	processor processor.Engine
}

func NewImageFromPath(path string) (*Image, error) {
	// read data
	absPath, _ := filepath.Abs(path)
	data, err := ioutil.ReadFile(absPath)
	if err != nil {
		return nil, err
	}

	return NewImageFromBuffer(data)
}

func NewImageFromBuffer(data []byte) (*Image, error) {
	// get format
	format, err := getImageFormat(data)
	if err != nil {
		return nil, err
	}
	if !util.Contains(SUPPORTED_FORMATS, format) {
		return nil, errors.New("Unsupported format")
	}

	return &Image{
		Data:      data,
		Format:    format,
		processor: processor.NewProcessor(format),
	}, nil
}

func (i *Image) Compress(ratio float32) (*Image, error) {
	compressed, err := i.processor.Compress(i.Data, ratio)
	if err != nil {
		return nil, err
	}

	return NewImageFromBuffer(compressed)
}

func (i *Image) Write(fn string) error {
	absPath, _ := filepath.Abs(fn)
	return ioutil.WriteFile(absPath, i.Data, os.ModePerm)
}

func (i *Image) Size() int {
	return len(i.Data)
}

func getImageFormat(data []byte) (string, error) {
	if len(data) < 4 {
		return "", errors.New("Error image format")
	}
	bytes := data[:4]
	if bytes[0] == 0x89 && bytes[1] == 0x50 && bytes[2] == 0x4E && bytes[3] == 0x47 {
		return constant.FormatPng, nil
	}
	if bytes[0] == 0xFF && bytes[1] == 0xD8 {
		return constant.FormatJpeg, nil
	}
	//if bytes[0] == 0x47 && bytes[1] == 0x49 && bytes[2] == 0x46 && bytes[3] == 0x38 {
	//	return FormatGif, nil
	//}
	//if bytes[0] == 0x42 && bytes[1] == 0x4D {
	//	return "bmp", nil
	//}
	return "", errors.New("Error image format")
}
