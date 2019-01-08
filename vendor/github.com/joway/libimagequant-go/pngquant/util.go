package pngquant

import (
	"image"
	"image/color"
)

// convert Go Image to RGBA32 bytes
func GoImageToRgba32(im image.Image) []byte {
	w := im.Bounds().Max.X
	h := im.Bounds().Max.Y
	ret := make([]byte, w*h*4)

	p := 0

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			r16, g16, b16, a16 := im.At(x, y).RGBA()

			ret[p+0] = uint8(r16 >> 8)
			ret[p+1] = uint8(g16 >> 8)
			ret[p+2] = uint8(b16 >> 8)
			ret[p+3] = uint8(a16 >> 8)
			p += 4
		}
	}

	return ret
}

// convert from RBG8 byte to Go Image
func Rgb8PaletteToGoImage(w, h int, rgb8data []byte, pal color.Palette) image.Image {
	rect := image.Rectangle{
		Max: image.Point{
			X: w,
			Y: h,
		},
	}

	ret := image.NewPaletted(rect, pal)

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			ret.SetColorIndex(x, y, rgb8data[y*w+x])
		}
	}

	return ret
}
