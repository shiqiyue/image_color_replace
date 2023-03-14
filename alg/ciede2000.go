package alg

import (
	ciede "github.com/mattn/go-ciede2000"
	"image/color"
	"math"
)

func cidedColorSimilarity(c1, c2 color.Color) float64 {
	return ciede.Diff(c1, c2)
}

func Ciede2000GetSimilarityColor(c1 color.Color, cs []color.Color) color.Color {
	if len(cs) == 0 {
		return c1
	}
	if len(cs) == 1 {
		return cs[0]
	}
	var similarity float64 = math.MaxFloat64
	var similarityColor = cs[0]
	for i, c := range cs {
		currentSimilarity := cidedColorSimilarity(c1, c)
		if currentSimilarity < similarity {
			similarityColor = cs[i]
			similarity = currentSimilarity
		}
	}
	return similarityColor
}
