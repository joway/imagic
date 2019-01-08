package pngquant

import "image"

/*
#include "libimagequant.h"
*/
import "C"

func Compress(img image.Image, quality int, speed int) (image.Image, error) {
	width := img.Bounds().Max.X
	height := img.Bounds().Max.Y

	// init libimagequant attributes holder
	attr, err := NewAttributes()
	if err != nil {
		return nil, err
	}
	if err := attr.SetSpeed(speed); err != nil {
		return nil, err
	}
	if err := attr.SetQuality(0, quality); err != nil {
		return nil, err
	}
	defer attr.Release()

	// new image from RGBA32 data (not Go Image)
	rgba32data := GoImageToRgba32(img)
	iqm, err := NewImage(attr, rgba32data, width, height, 0)
	if err != nil {
		return nil, err
	}
	defer iqm.Release()

	// compress image
	res, err := iqm.Quantize(attr)
	if err != nil {
		return nil, err
	}
	defer res.Release()

	// get RBG8 data from compressed image
	rgb8data, err := res.WriteRemappedImage()
	if err != nil {
		return nil, err
	}

	// convert RBG8 to Go Image
	outImg := Rgb8PaletteToGoImage(res.GetImageWidth(), res.GetImageHeight(),
		rgb8data, res.GetPalette())

	return outImg, nil
}

func LibVersion() string {
	return C.LIQ_VERSION_STRING
}
