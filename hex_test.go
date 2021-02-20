package color_test

import (
	"errors"
	"image/color"
	"testing"

	gcolor "github.com/ggicci/color"
)

func isEqualColor(a, b color.Color) bool {
	r1, g1, b1, a1 := a.RGBA()
	r2, g2, b2, a2 := b.RGBA()
	return r1 == r2 && g1 == g2 && b1 == b2 && a1 == a2
}

func TestFromHex(t *testing.T) {
	shouldPass := []map[string]interface{}{
		{"hex": "#09c", "rgba": color.RGBA{0, 153, 204, 255}},
		{"hex": "#0099cc", "rgba": color.RGBA{0, 153, 204, 255}},
		{"hex": "#09c7", "rgba": color.RGBA{0, 153, 204, 119}},
		{"hex": "#0099cc77", "rgba": color.RGBA{0, 153, 204, 119}},
	}
	for _, spec := range shouldPass {
		c, err := gcolor.FromHex(spec["hex"].(string))
		if err != nil {
			t.Errorf("hex [%s] expect nil, got %v", spec["hex"], err)
		}
		if !isEqualColor(c, spec["rgba"].(color.Color)) {
			t.Errorf("hex [%s] expect %v, got %v", spec["hex"], spec["rgba"], c)
		}
	}

	shouldFail := []string{
		"481483",
		"#1G2F34",
		"#3",
		"#32557",
	}

	for _, spec := range shouldFail {
		_, err := gcolor.FromHex(spec)
		if !errors.Is(err, gcolor.ErrMalformedColor) {
			t.Errorf("hex [%s] should be malformed, got %#v", spec, err)
		}
	}

}
