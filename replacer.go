package image_color_replace

import (
	"github.com/anthonynsimon/bild/parallel"
	"github.com/g4s8/hexcolor"
	"github.com/mazznoer/csscolorparser"
	cmap "github.com/orcaman/concurrent-map/v2"
	"github.com/pkg/errors"
	"image"
	"image/color"
	"image_color_replace/alg"
)

// Replace replace all colors from inputImage to colorStrs, and ignore colors in ignoreColorStrs
func Replace(inputImage image.Image, colorStrs []string, ignoreColorStrs []string, algoType alg.ALGO_TYPE) (image.Image, error) {
	// parse color
	cs := make([]color.Color, 0)
	ignoreCs := make([]color.Color, 0)
	for _, s := range colorStrs {
		c, err := csscolorparser.Parse(s)
		if err != nil {
			return nil, errors.Wrap(err, "parse color error")
		}
		cs = append(cs, c)
	}
	if len(ignoreColorStrs) > 0 {
		for _, s := range ignoreColorStrs {
			c, err := hexcolor.Parse(s)
			if err != nil {
				return nil, errors.Wrap(err, "parse color error")
			}
			ignoreCs = append(ignoreCs, c)
		}
	}

	var xMax = inputImage.Bounds().Dx()
	var yMax = inputImage.Bounds().Dy()
	dst := image.NewRGBA(inputImage.Bounds())
	c := cmap.New[color.Color]()
	parallel.Line(yMax, func(start, end int) {
		for y := start; y < end; y++ {
			for x := 0; x < xMax; x++ {
				ip := inputImage.At(x, y)
				r, g, b, a := ip.RGBA()
				if a == 0 {
					dst.Set(x, y, ip)
					continue
				}
				// Identify whether the current pixel is the ignore color
				ignoreColor := false
				for _, ignoreCs := range ignoreCs {
					ignoreR, ignoreG, ignoreB, _ := ignoreCs.RGBA()
					if r == ignoreR && g == ignoreG && b == ignoreB {
						dst.Set(x, y, ip)
						ignoreColor = true
						continue
					}
				}
				if ignoreColor {
					continue
				}

				// get similar alg of pixels
				similarityColor := alg.GetSimilarityColor(ip, cs, algoType, &c)
				rr, rg, rb, _ := similarityColor.RGBA()
				dst.Set(x, y, color.RGBA{
					R: uint8(rr),
					G: uint8(rg),
					B: uint8(rb),
					A: uint8(a),
				})
			}

		}
	})
	return dst, nil
}
