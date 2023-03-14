package alg

import (
	"image/color"
	"math"
)

type ColorHSV struct {
	H, S, V float64
}

// Convert RGB color to HSV color space
func RGBToHSV(c color.Color) ColorHSV {
	ta, tg, tb, _ := c.RGBA()
	// Convert an RGB value to a floating point number in the range [0, 1]
	rf := float64(ta) / 65535.0
	gf := float64(tg) / 65535.0
	bf := float64(tb) / 65535.0

	// Compute values in HSV color space
	cmax := math.Max(math.Max(rf, gf), bf)
	cmin := math.Min(math.Min(rf, gf), bf)
	delta := cmax - cmin

	h := 0.0
	if delta == 0 {
		h = 0.0
	} else if cmax == rf {
		h = math.Mod(((gf - bf) / delta), 6.0)
	} else if cmax == gf {
		h = ((bf-rf)/delta + 2.0)
	} else {
		h = ((rf-gf)/delta + 4.0)
	}
	h = h * 60.0

	s := 0.0
	if cmax == 0 {
		s = 0.0
	} else {
		s = delta / cmax
	}

	v := cmax

	return ColorHSV{h, s, v}
}

func hsvColorSimilarity(hsv1, hsv2 ColorHSV) float64 {
	dh := math.Abs(hsv1.H - hsv2.H)
	if dh > 180.0 {
		dh = 360.0 - dh
	}
	ds := math.Abs(hsv1.S - hsv2.S)
	dv := math.Abs(hsv1.V - hsv2.V)
	return 1.0 - (dh/180.0+ds+dv)/3.0
}

func HsvGetSimilarityColor(c1 color.Color, cs []color.Color) color.Color {

	var similarity float64 = 0
	var similarityColor = cs[0]
	for i, c := range cs {
		currentSimilarity := hsvColorSimilarity(RGBToHSV(c1), RGBToHSV(c))
		if currentSimilarity > similarity {
			similarityColor = cs[i]
			similarity = currentSimilarity
		}
	}
	return similarityColor
}
