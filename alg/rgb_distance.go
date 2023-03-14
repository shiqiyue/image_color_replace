package alg

import (
	"image/color"
	"math"
)

func rgbColorSimilarity(c1, c2 color.Color) float64 {
	distance := rgbColorDistance(c1, c2)
	return 1.0 / (1.0 + distance)
}

func rgbColorDistance(c1, c2 color.Color) float64 {

	r1, g1, b1, _ := c1.RGBA()
	r2, g2, b2, _ := c2.RGBA()
	return math.Sqrt(math.Pow(float64(r1-r2), 2) + math.Pow(float64(g1-g2), 2) + math.Pow(float64(b1-b2), 2))
}

func RgbGetSimilarityColor(c1 color.Color, cs []color.Color) color.Color {
	if len(cs) == 0 {
		return c1
	}
	if len(cs) == 1 {
		return cs[0]
	}
	var similarity float64 = 0
	var similarityColor = cs[0]
	for i, c := range cs {
		currentSimilarity := rgbColorSimilarity(c1, c)
		if currentSimilarity > similarity {
			similarityColor = cs[i]
			similarity = currentSimilarity
		}
	}
	return similarityColor
}
