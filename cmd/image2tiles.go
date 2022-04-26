package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	imagefileName   = flag.String("i", "", "image file name")
	tileSize        = flag.Int("s", 256, "the tile height/width")
	template        = flag.String("t", "-%d-%d-%d", "template filename pattern")
	outputDirectory = flag.String("o", "out", "output directory")
	backgroundColor = flag.String("b", "#FFF", "the background color to be used for the tiles")
)

func main() {
	flag.Parse()

	if *imagefileName == "" {
		fmt.Println("image file is empty")
		return
	}

	if _, err := os.Stat(*imagefileName); err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("image file %s is not exist\n", *imagefileName)
			return
		}
		fmt.Printf("image file: %v\n", err)
		return
	}

	if err := os.MkdirAll(*outputDirectory, os.ModePerm); err != nil {
		fmt.Printf("create output directory: %v\n", err)
		return
	}

}
