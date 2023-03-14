package image_color_replace

import (
	"github.com/disintegration/imaging"
	"image_color_replace/alg"
	"testing"
)

func TestReplaceByCided(t *testing.T) {
	img, err := imaging.Open("./example/test.jpeg")
	if err != nil {
		t.Error(err)
		return
	}
	colorStrs := []string{"white", "black", "red", "green", "yellow"}
	reImg, err := Replace(img, colorStrs, nil, alg.CIDED)
	if err != nil {
		t.Error(err)
		return
	}
	err = imaging.Save(reImg, "./example/test_by_cided.jpeg")
	if err != nil {
		t.Error(err)
		return
	}
}

func TestReplaceByRgb(t *testing.T) {
	img, err := imaging.Open("./example/test.jpeg")
	if err != nil {
		t.Error(err)
		return
	}
	colorStrs := []string{"white", "black", "red", "green", "yellow"}
	reImg, err := Replace(img, colorStrs, nil, alg.RGB)
	if err != nil {
		t.Error(err)
		return
	}
	err = imaging.Save(reImg, "./example/test_by_rgb.jpeg")
	if err != nil {
		t.Error(err)
		return
	}
}

func TestReplaceByHsv(t *testing.T) {
	img, err := imaging.Open("./example/test.jpeg")
	if err != nil {
		t.Error(err)
		return
	}
	colorStrs := []string{"white", "black", "red", "green", "yellow"}
	reImg, err := Replace(img, colorStrs, nil, alg.HSV)
	if err != nil {
		t.Error(err)
		return
	}
	err = imaging.Save(reImg, "./example/test_by_hsv.jpeg")
	if err != nil {
		t.Error(err)
		return
	}
}
