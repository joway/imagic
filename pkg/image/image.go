package image

import (
	"bytes"
	"github.com/joway/imagic/pkg/constant"
	"github.com/joway/imagic/pkg/processor"
	"github.com/joway/imagic/pkg/util"
	"github.com/pkg/errors"
	"github.com/thoas/go-funk"
	"image"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
)

var supportedFormats = []string{constant.FormatPng, constant.FormatJpeg}

// Image is the internal image struct
type Image struct {
	Raw       image.Image
	Format    string
	processor processor.Engine
}

// NewImageFromPath return the image from path
func NewImageFromPath(path string) (*Image, error) {
	// read data
	absPath, _ := filepath.Abs(path)
	data, err := ioutil.ReadFile(absPath)
	if err != nil {
		return nil, err
	}

	return NewImageFromBuffer(data)
}

// NewImageFromBuffer return the image from buffer
func NewImageFromBuffer(data []byte) (*Image, error) {
	// get format
	format, err := getImageFormat(data)
	if err != nil {
		return nil, err
	}
	if !funk.Contains(supportedFormats, format) {
		return nil, errors.New("Unsupported format")
	}

	var raw image.Image
	switch format {
	case constant.FormatPng:
		raw, err = png.Decode(bytes.NewReader(data))
		if err != nil {
			return nil, err
		}
	case constant.FormatJpeg:
		raw, err = jpeg.Decode(bytes.NewReader(data))
		if err != nil {
			return nil, err
		}
	default:
	}

	return NewImage(raw, format)
}

// NewImage return the image from buffer
func NewImage(raw image.Image, format string) (*Image, error) {
	return &Image{
		Raw:       raw,
		Format:    format,
		processor: processor.NewProcessor(format),
	}, nil
}

// Compress the image
func (i *Image) Compress(quality int) (*Image, error) {
	outputImg, err := i.processor.Compress(i.Raw, quality)
	if err != nil {
		return nil, err
	}
	return NewImage(outputImg, i.Format)
}

// Resize the image
func (i *Image) Resize(width int, height int) (*Image, error) {
	outputImg, err := i.processor.Resize(i.Raw, width, height)
	if err != nil {
		return nil, err
	}
	return NewImage(outputImg, i.Format)
}

// WaterMark the image
func (i *Image) WaterMark(texture *Image, x string, y string) (*Image, error) {
	watermarkX, err := parseXY(i.Raw.Bounds().Max.X, x)
	if err != nil {
		return nil, err
	}
	watermarkY, err := parseXY(i.Raw.Bounds().Max.Y, y)
	if err != nil {
		return nil, err
	}

	outputImg, err := i.processor.WaterMark(i.Raw, texture.Raw, watermarkX, watermarkY)
	if err != nil {
		return nil, err
	}

	return NewImage(outputImg, i.Format)
}

// Write the image into filename
func (i *Image) Write(fn string) error {
	absPath, _ := filepath.Abs(fn)
	if err := util.EnsureDir(fn); err != nil {
		return err
	}

	outputFile, err := os.Create(absPath)
	if err != nil {
		return err
	}
	switch i.Format {
	case constant.FormatPng:
		if err := png.Encode(outputFile, i.Raw); err != nil {
			return err
		}
	case constant.FormatJpeg:
		if err := jpeg.Encode(outputFile, i.Raw, nil); err != nil {
			return err
		}
	default:
	}

	return outputFile.Close()
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
	return "", errors.New("Error image format")
}

func parseXY(max int, v string) (int, error) {
	if len(v) == 0 {
		return 0, nil
	}

	flag := v[0]
	if flag == '+' {
		return strconv.Atoi(v[1:])
	} else if flag == '-' {
		n, err := strconv.Atoi(v[1:])
		if err != nil {
			return n, err
		}
		return max - n, nil
	} else {
		return strconv.Atoi(v)
	}
}
