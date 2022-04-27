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

	imageFilename   = flag.String("f", "", "image file name")
	tileSize        = flag.Int("s", 256, "the tile height/width")
	template        = flag.String("t", "%d-%d-%d.jpg", "template filename pattern")
	outputDirectory = flag.String("o", "out", "output directory")
	backgroundColor = flag.String("b", "#FFF", "the background color to be used for the tiles")
)

func main() {
	flag.Parse()

	if err := parse(); err == nil {
		converter := image2tiles.NewConverter()
		if err := converter.Prepare(*imageFilename, *backgroundColor); err != nil {
			fmt.Println(err)
		}
		if _, err := converter.Subdivide(0, [2]int{*tileSize, *tileSize}, [2]int{0, 0}, *template, *outputDirectory); err != nil {
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

	if err := os.MkdirAll(*outputDirectory, os.ModePerm); err != nil {
		return fmt.Errorf("create output directory: %v\n", err)
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
