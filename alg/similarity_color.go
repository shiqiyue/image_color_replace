package alg

import (
	"github.com/mazznoer/csscolorparser"
	cmap "github.com/orcaman/concurrent-map/v2"
	"image/color"
)

type ALGO_TYPE int

var (
	HSV       ALGO_TYPE = 1
	RGB       ALGO_TYPE = 2
	CIEDE2000 ALGO_TYPE = 3
)

func GetSimilarityColor(c1 color.Color, cs []color.Color, algoType ALGO_TYPE, matchCache *cmap.ConcurrentMap[string, color.Color]) color.Color {
	if len(cs) == 0 {
		return c1
	}
	if len(cs) == 1 {
		return cs[0]
	}
	// Determine whether it can be obtained from the cache of the matching result
	c1Hex := ""
	if matchCache != nil {
		r, g, b, _ := c1.RGBA()
		rf := float64(r) / 65535.0
		gf := float64(g) / 65535.0
		bf := float64(b) / 65535.0
		c1Hex = csscolorparser.Color{
			R: rf,
			G: gf,
			B: bf,
			A: 1,
		}.HexString()
		matchResult, ok := matchCache.Get(c1Hex)
		if ok {
			return matchResult
		}
	}

	var similarityColor color.Color
	switch algoType {
	case HSV:
		similarityColor = HsvGetSimilarityColor(c1, cs)
	case CIEDE2000:
		similarityColor = Ciede2000GetSimilarityColor(c1, cs)
	case RGB:
		similarityColor = RgbGetSimilarityColor(c1, cs)
	}

	if matchCache != nil {
		matchCache.Set(c1Hex, similarityColor)
	}
	return similarityColor
}
