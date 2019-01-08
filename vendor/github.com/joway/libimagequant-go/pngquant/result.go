package pngquant

/*
#include "libimagequant.h"
*/
import "C"
import (
	"image/color"
	"unsafe"
)

// Callers must not use i object once Release has been called on the parent
// Image struct.
type Result struct {
	p  *C.struct_liq_result
	im *Image
}

// Free memory. Callers must not use i object after Release has been called.
func (r *Result) Release() {
	C.liq_result_destroy(r.p)
}

func (r *Result) WriteRemappedImage() ([]byte, error) {
	if r.im.released {
		return nil, ErrUseAfterFree
	}

	buffSize := r.im.w * r.im.h
	buff := make([]byte, buffSize)

	iqe := C.liq_write_remapped_image(r.p, r.im.p, unsafe.Pointer(&buff[0]), C.size_t(buffSize))
	if iqe != C.LIQ_OK {
		return nil, translateError(iqe)
	}

	return buff, nil
}

func (r *Result) GetPalette() color.Palette {
	ptr := C.liq_get_palette(r.p) // copy struct content
	max := int(ptr.count)

	ret := make([]color.Color, max)
	for i := 0; i < max; i += 1 {
		ret[i] = color.RGBA{
			R: uint8(ptr.entries[i].r),
			G: uint8(ptr.entries[i].g),
			B: uint8(ptr.entries[i].b),
			A: uint8(ptr.entries[i].a),
		}
	}

	return ret
}

func (r *Result) GetImageWidth() int {
	// C.liq_image_get_width
	return r.im.w
}

func (r *Result) GetImageHeight() int {
	// C.liq_image_get_height
	return r.im.h
}
