package image2tiles

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"os"
	"strings"
)

func LoadImage(path string) (image.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	im, _, err := image.Decode(file)
	return im, err
}

func LoadPNG(path string) (image.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return png.Decode(file)
}

func SavePNG(path string, im image.Image) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	return png.Encode(file, im)
}

func SaveJPG(path string, im image.Image, quality int) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	var opt jpeg.Options
	opt.Quality = quality

	return jpeg.Encode(file, im, &opt)
}

func ImageToRGBA(src image.Image) *image.RGBA {
	bounds := src.Bounds()
	dst := image.NewRGBA(bounds)
	draw.Draw(dst, bounds, src, bounds.Min, draw.Src)
	return dst
}

func ParseHexColor(x string) color.RGBA {
	x = strings.TrimPrefix(x, "#")
	c := color.RGBA{R: 255, G: 255, B: 255, A: 255}

	if len(x) == 3 {
		format := "%1x%1x%1x"
		fmt.Sscanf(x, format, &c.R, &c.G, &c.B)
		c.R |= c.R << 4
		c.G |= c.G << 4
		c.B |= c.B << 4
	}
	if len(x) == 6 {
		format := "%02x%02x%02x"
		fmt.Sscanf(x, format, &c.R, &c.G, &c.B)
	}
	if len(x) == 8 {
		format := "%02x%02x%02x%02x"
		fmt.Sscanf(x, format, &c.R, &c.G, &c.B, &c.A)
	}

	return c
}
