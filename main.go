package main

import (
	"flag"
	"fmt"
	"log"
	"path/filepath"

	"github.com/disintegration/imaging"
	"github.com/schollz/progressbar/v3"
)

var (
	Out        = flag.String("out", "sheets-%03d.jpg", "Output template")
	Gap        = flag.Int("gap", 20, "Gap between thumbnails in px")
	TextHeight = flag.Int("text-height", 10, "Height of text in px")
	Rows       = flag.Int("rows", 6, "Number of rows of thumbnails")
	Cols       = flag.Int("cols", 7, "Number of columns of thumbnails")
	ThumbX     = flag.Int("thumb-x", 300, "Thumbnail X size")
	DcRaw      = flag.String("dcraw", dcraw, "Path to dcraw executable for this syste,")
	Progress   = flag.Bool("progress", false, "Show progress bar instead of text")

	//Font       = flag.String("font", "font.ttf", "Font file name")
	//font *truetype.Font
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

	/*
		// Read the font data.
		fontBytes, err := ioutil.ReadFile(*Font)
		if err != nil {
			panic(err)
		}
		font, err = freetype.ParseFont(fontBytes)
		if err != nil {
			panic(err)
		}
	*/

	// Resolve files
	x := []string{}
	for i := range files {
		matches, _ := filepath.Glob(files[i])
		x = append(x, matches...)
	}
	files = x

	// Split into image batches of 7w x 6h
	batches := int(len(files)/(*Rows**Cols)) + 1

	log.Printf("Found %d images forming %d batches", len(files), batches)

	var pb *progressbar.ProgressBar

	if *Progress {
		pb = progressbar.Default(int64(len(files)))
	}

	//
	for batch := 0; batch < batches; batch++ {
		imglow := (batch * (*Rows * *Cols))
		imghigh := ((batch + 1) * (*Rows * *Cols)) - 1
		if imghigh >= len(files) {
			imghigh = len(files) - 1
		}

		if !*Progress {
			log.Printf("Processing batch #%d [ %d .. %d ] into %s", batch+1, imglow, imghigh, fmt.Sprintf(*Out, batch+1))
		}

		processFiles := files[imglow : imghigh+1]
		img, err := SheetFromFiles(processFiles, pb)
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
