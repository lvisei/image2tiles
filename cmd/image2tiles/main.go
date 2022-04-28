package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/lvisei/image2tiles"
)

var (
	// these are set in build step
	version = "unversioned"
	//lint:ignore U1000 embedded by goreleaser
	commit = "?"
	//lint:ignore U1000 embedded by goreleaser
	date = "?"

	imageFilename   = flag.String("f", "", "Image file name")
	tileSize        = flag.Int("s", 512, "The tile height/width")
	output          = flag.String("o", "out/%d/%d-%d.jpg", "Output file pattern")
	backgroundColor = flag.String("b", "#FFF", "The background color to be used for the tiles")
)

func main() {
	flag.Parse()

	if err := parse(); err == nil {
		converter := image2tiles.NewConverter()
		if err := converter.Prepare(*imageFilename, *backgroundColor); err != nil {
			fmt.Println(err)
		}
		if err := converter.Execute([2]int{*tileSize, *tileSize}, true, 75, *output); err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Print(err)
		usage()
	}

}

func parse() error {
	if *imageFilename == "" {
		return errors.New("image file should not be empty")
	}

	if _, err := os.Stat(*imageFilename); err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("image file %s is not exist\n", *imageFilename)
		}
		return fmt.Errorf("image file: %v\n", err)
	}

	if *output == "" {
		return errors.New("output file pattern should not be empty")
	}

	return nil
}

func usage() {
	fmt.Fprintf(os.Stderr, `
image2tiles

Version: v%s
HomePage: github.com/lvisei/image2tiles
Issue   : github.com/lvisei/image2tiles/issues
Author  : lvisei

Usage: image2tiles -f <filename> [-s] [-t] [-b] [-o]

Options:
	`, version)
	flag.PrintDefaults()
}
