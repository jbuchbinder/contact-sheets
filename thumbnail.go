package main

import (
	"bytes"
	"image"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/disintegration/imaging"
)

// ThumbnailFromFile returns a ThumbX*ThumbX bounded image thumbnail
func ThumbnailFromFile(fn string) (image.Image, error) {
	log.Printf("ThumbnailFromFile(%s)", fn)

	realFn := fn

	if strings.HasSuffix(fn, ".CR2") ||
		strings.HasSuffix(fn, ".DNG") ||
		strings.HasSuffix(fn, ".NEF") {
		realFn = os.TempDir() + string(os.PathSeparator) + path.Base(fn) + ".jpg"
		err := ExtractThumbnailFromRaw(fn, realFn)
		defer os.Remove(realFn)
		if err != nil {
			return nil, err
		}
	}

	srcImage, err := imaging.Open(realFn, imaging.AutoOrientation(true))
	if err != nil {
		return srcImage, err
	}
	img := imaging.Fit(srcImage, *ThumbX, *ThumbX, imaging.NearestNeighbor)
	return img, nil
}

// ExtractThumbnailFromRaw pulls the embedded thumbnail from a raw image using
// a local copy of the dcraw library
func ExtractThumbnailFromRaw(in, out string) error {
	cmd := exec.Command(*DcRaw, "-c", "-e", in)
	var output bytes.Buffer
	cmd.Stdout = &output
	err := cmd.Run()
	if err != nil {
		return err
	}
	return ioutil.WriteFile(out, output.Bytes(), 0600)
}
