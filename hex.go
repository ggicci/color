package color

import (
	"fmt"
	"image/color"
	"regexp"
	"strconv"
)

var (
	reHexColor3Digits = regexp.MustCompile("^#[0-9a-fA-F]{3}$")
	reHexColor4Digits = regexp.MustCompile("^#[0-9a-fA-F]{4}$")
	reHexColor6Digits = regexp.MustCompile("^#[0-9a-fA-F]{6}$")
	reHexColor8Digits = regexp.MustCompile("^#[0-9a-fA-F]{8}$")

	ErrMalformedHexColor = fmt.Errorf("hex: %w", ErrMalformedColor)
)

type Hex color.RGBA

func hv(h string) uint8 {
	if len(h) == 1 {
		h = h + h
	}
	v, _ := strconv.ParseUint(h, 16, 8)
	return uint8(v)
}

// FromHex parses hexadecimal form of color values.
// See https://en.wikipedia.org/wiki/Web_colors
// Available forms are:
// - Three digits: #09C
// - Six digits: #0099CC
// - With alpha: #09C2, #0099CC22
func FromHex(value string) (color.Color, error) {
	if reHexColor3Digits.MatchString(value) {
		return color.RGBA{hv(value[1:2]), hv(value[2:3]), hv(value[3:4]), 255}, nil
	}
	if reHexColor4Digits.MatchString(value) {
		return color.RGBA{hv(value[1:2]), hv(value[2:3]), hv(value[3:4]), hv(value[4:5])}, nil
	}
	if reHexColor6Digits.MatchString(value) {
		return color.RGBA{hv(value[1:3]), hv(value[3:5]), hv(value[5:7]), 255}, nil
	}
	if reHexColor8Digits.MatchString(value) {
		return color.RGBA{hv(value[1:3]), hv(value[3:5]), hv(value[5:7]), hv(value[7:9])}, nil
	}
	return color.White, ErrMalformedHexColor
}
