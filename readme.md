# image_color_replace

## Contents

* [Installation](#installation)
* [Golang version](#golang-version)
* [Usage](#usage)
* [Sample](#sample)

# installation
```bash
go get -u github.com/shiqiyue/image_color_replace
```

# Golang version
Golang >= 1.18 is required.

# usage
this package Use three different algorithms to get the most similar colors; 
1. rgb_distance, This algorithm is only to calculate the distance of rgb and get the color of the closest distance
2. hsv_distance, This algorithm first converts rbg to hsv, and then gets the most similar hsv color
3. ciede2000(recommend), you can see more from https://zh.wikipedia.org/wiki/%E9%A2%9C%E8%89%B2%E5%B7%AE%E5%BC%82#CIEDE2000
```golang
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
```


# sample
## source
![sample](./example/test.jpeg)

## replace by cided
![sample](./example/test_by_cided.jpeg)

## replace by hsv distance
![sample](./example/test_by_hsv.jpeg)

## replace by rgb distance
![sample](./example/test_by_rgb.jpeg)