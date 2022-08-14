package revisiting_arrays_and_slices

import (
	"errors"
	"strconv"
)

type Color struct {
	R, G, B int64
	A       float32
}

//все нуждается в доработке

func ColorMixer(colors ...Color) (Color, error) {
	var ResultColor Color
	lenColors := len(colors)
	for _, color := range colors {
		if ResultColor.R < color.R {
			ResultColor.R += color.R / int64(lenColors)
		}
		if ResultColor.R > color.R {
			ResultColor.R -= color.R / int64(lenColors)
		}
		if ResultColor.G < color.G {
			ResultColor.G += color.G / int64(lenColors)
		}
		if ResultColor.G > color.G {
			ResultColor.G -= color.G / int64(lenColors)
		}
		if ResultColor.B < color.B {
			ResultColor.B += color.B / int64(lenColors)
		}
		if ResultColor.B > color.B {
			ResultColor.B -= color.B / int64(lenColors)
		}
		if ResultColor.A < color.A {
			ResultColor.A += color.A / float32(lenColors)
		}
		if ResultColor.A > color.A {
			ResultColor.A -= color.A / float32(lenColors)
		}
	}
	return ResultColor, nil
}

func NewColor(RGB string, A float32) (Color, error) {
	if RGB[0] == '#' {
		RGB = RGB[1:]
	}
	if len(RGB) != 6 {
		return Color{}, errors.New("RGB should contain 6 symbols")
	}
	var (
		newColor Color
		err      error
	)
	newColor.R, err = strconv.ParseInt(RGB[:2], 16, 64)
	if err != nil {
		return Color{}, err
	}
	newColor.G, err = strconv.ParseInt(RGB[2:4], 16, 64)
	if err != nil {
		return Color{}, err
	}
	newColor.B, err = strconv.ParseInt(RGB[4:], 16, 64)
	if err != nil {
		return Color{}, err
	}
	newColor.A = A
	return newColor, err
}
