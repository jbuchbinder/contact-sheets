package main

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"os"
	"path"
	"strings"

	"github.com/disintegration/imaging"
	"github.com/schollz/progressbar/v3"
)

// SheetFromFiles creates a contact sheet for a list of files
func SheetFromFiles(files []string, pb *progressbar.ProgressBar) (image.Image, error) {
	dst := imaging.New(
		((*ThumbX+*Gap)*7)+*Gap,                         // gap on side
		((*ThumbX + *Gap + *TextHeight + *Gap) * *Rows), // gap on top, label, gap on bottom
		color.NRGBA{100, 100, 100, 0})

	if len(files) > (*Rows * *Cols) {
		return dst, fmt.Errorf("SheetFromFiles was passed %d files", len(files))
	}
	for i := 0; i < len(files); i++ {
		row := int(i / *Cols)
		col := i % *Cols
		img, err := ThumbnailFromFile(files[i])
		if err != nil {
			log.Printf("WARN: %s", err.Error())
			continue
		}

		// Calculate height/width offset based on difference between size of
		// image and the maximum thumbnail size
		widthOffset := 0
		if *ThumbX-img.Bounds().Dx() > 0 {
			widthOffset = (*ThumbX - img.Bounds().Dx()) / 2
		}
		heightOffset := 0
		if *ThumbX-img.Bounds().Dy() > 0 {
			heightOffset = (*ThumbX - img.Bounds().Dy()) / 2
		}

		// Composite image in position
		if !*Progress {
			log.Printf("Pasting image size %d x %d at %d:%d with offsets %d/%d at pos %d/%d",
				img.Bounds().Dx(), img.Bounds().Dy(),
				row, col,
				widthOffset, heightOffset,
				*Gap+(*ThumbX*col)+(col**Gap)+widthOffset,
				*Gap+(*ThumbX*row)+(row*(*Gap+*TextHeight+*Gap))+heightOffset)
		}
		dst = imaging.Paste(dst,
			img,
			image.Pt(
				*Gap+(*ThumbX*col)+(col**Gap)+widthOffset,
				*Gap+(*ThumbX*row)+(row*(*Gap+*TextHeight+*Gap))+heightOffset,
			))

		// Add text
		bounds := image.Rect(
			*Gap+(*ThumbX*col)+(col**Gap),
			*Gap+(*ThumbX*row)+(row*(*Gap+*TextHeight+*Gap)),
			*Gap+(*ThumbX*col)+(col**Gap)+*ThumbX+*Gap,
			*Gap+(*ThumbX*row)+(row*(*Gap+*TextHeight+*Gap))+*ThumbX+*Gap,
		)
		if !*Progress {
			log.Printf("Text bounds: %s : %#v", path.Base(files[i]), bounds)
		}

		fn := path.Base(files[i])
		if os.PathSeparator == '\\' {
			lastSlash := strings.LastIndex(fn, string(os.PathSeparator))
			fn = fn[lastSlash+1:]
		}

		addLabel(dst,
			*Gap+(*ThumbX*col)+(col**Gap)+(*Gap*2),
			*Gap+(*ThumbX*row)+(row*(*Gap+*TextHeight+*Gap))+*ThumbX+*Gap,
			fn)

		if pb != nil {
			pb.Add(1)
		}
	}

	return dst, nil
}
