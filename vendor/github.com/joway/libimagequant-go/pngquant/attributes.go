package pngquant

import "github.com/pkg/errors"

/*
#include "libimagequant.h"
*/
import "C"

const (
	SPEED_SLOWEST = 1
	SPEED_DEFAULT = 3
	SPEED_FASTEST = 10
)

type Attributes struct {
	p *C.struct_liq_attr
}

// Callers MUST call Release() on the returned object to free memory.
func NewAttributes() (*Attributes, error) {
	pAttr := C.liq_attr_create()
	if pAttr == nil {
		return nil, errors.New("Unsupported platform")
	}

	return &Attributes{p: pAttr}, nil
}

// Free memory. Callers must not use this object after Release has been called.
func (a *Attributes) Release() {
	C.liq_attr_destroy(a.p)
}

func (a *Attributes) SetQuality(minQuality int, maxQuality int) error {
	return translateError(C.liq_set_quality(a.p, C.int(minQuality), C.int(maxQuality)))
}

func (a *Attributes) SetSpeed(speed int) error {
	return translateError(C.liq_set_speed(a.p, C.int(speed)))
}
