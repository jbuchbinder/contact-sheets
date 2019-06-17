package main

import (
	"image"
	"log"

	"github.com/disintegration/imaging"
)

// ThumbnailFromFile returns a ThumbX*ThumbX bounded image thumbnail
func ThumbnailFromFile(fn string) (image.Image, error) {
	log.Printf("ThumbnailFromFile(%s)", fn)
	srcImage, err := imaging.Open(fn, imaging.AutoOrientation(true))
	if err != nil {
		return srcImage, err
	}
	img := imaging.Fit(srcImage, *ThumbX, *ThumbX, imaging.NearestNeighbor)
	return img, nil
}
