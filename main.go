package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/disintegration/imaging"
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
)

var (
	Out        = flag.String("out", "sheets-%03d.jpg", "Output template")
	Gap        = flag.Int("gap", 20, "Gap between thumbnails in px")
	Font       = flag.String("font", "font.ttf", "Font file name")
	TextHeight = flag.Int("text-height", 10, "Height of text in px")
	Rows       = flag.Int("rows", 6, "Number of rows of thumbnails")
	Cols       = flag.Int("cols", 7, "Number of columns of thumbnails")
	ThumbX     = flag.Int("thumb-x", 300, "Thumbnail X size")
	DcRaw      = flag.String("dcraw", dcraw, "Path to dcraw executable for this syste,")

	font *truetype.Font
)

func main() {
	flag.Parse()
	files := flag.Args()
	if *Out == "" {
		flag.PrintDefaults()
		return
	}

	// Add add'l files from directory, etc

	if len(files) < 1 {
		panic("No images to process")
	}

	// Read the font data.
	fontBytes, err := ioutil.ReadFile(*Font)
	if err != nil {
		panic(err)
	}
	font, err = freetype.ParseFont(fontBytes)
	if err != nil {
		panic(err)
	}

	// Split into image batches of 7w x 6h
	batches := int(len(files)/(*Rows**Cols)) + 1

	log.Printf("Found %d images forming %d batches", len(files), batches)

	//
	for batch := 0; batch < batches; batch++ {
		imglow := (batch * (*Rows * *Cols))
		imghigh := ((batch + 1) * (*Rows * *Cols)) - 1
		if imghigh >= len(files) {
			imghigh = len(files) - 1
		}
		log.Printf("Processing batch #%d [ %d .. %d ] into %s", batch+1, imglow, imghigh, fmt.Sprintf(*Out, batch+1))

		processFiles := files[imglow : imghigh+1]
		img, err := SheetFromFiles(processFiles)
		if err != nil {
			log.Printf("ERR: %s", err.Error())
			continue
		}
		err = imaging.Save(img, fmt.Sprintf(*Out, batch+1))
		if err != nil {
			log.Printf("ERR: %s", err.Error())
			continue
		}
	}
}
