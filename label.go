package main

import (
	"image"
	"image/color"

	fontpkg "golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
)

func addLabel(img *image.NRGBA, x, y int, label string) {
	col := color.White // color.RGBA{200, 100, 0, 255}
	point := fixed.Point26_6{X: fixed.Int26_6(x * 64), Y: fixed.Int26_6(y * 64)}

	d := &fontpkg.Drawer{
		Dst:  img,
		Src:  image.NewUniform(col),
		Face: basicfont.Face7x13,
		Dot:  point,
	}
	d.DrawString(label)
}
