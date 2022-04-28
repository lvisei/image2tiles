package image2tiles

import (
	"errors"
	"fmt"
	"github.com/disintegration/imaging"
	"image"
	"image/color"
	"math"
	"os"
	"path"
)

var (
	resampleFilter = imaging.Lanczos
)

type Converter struct {
	Image               image.Image
	ImageWidth          int
	ImageHeight         int
	OriginalImageWidth  int
	OriginalImageHeight int
	MaxZoom             int
	TileSize            [2]int
}

func NewConverter() *Converter {
	return &Converter{}
}

// Prepare a large image for tiling
func (converter *Converter) Prepare(imageName string, bgColor string) error {
	originalImage, err := LoadImage(imageName)
	if err != nil {
		return errors.New("open newImage file failed: " + err.Error())
	}

	originalImageWidth := originalImage.Bounds().Size().X
	originalImageHeight := originalImage.Bounds().Size().Y

	imageWidth := 1
	imageHeight := 1

	for imageWidth < originalImageWidth || imageHeight < originalImageHeight {
		imageWidth *= 2
		imageHeight *= 2
	}

	r, g, b, a := ParseHexColor(bgColor)

	newImage := imaging.New(imageWidth, imageHeight, color.NRGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: uint8(a)})
	originalImage = imaging.Thumbnail(originalImage, imageWidth, imageHeight, resampleFilter)
	newImage = imaging.Paste(newImage, originalImage, image.Pt((imageWidth-originalImageWidth)/2, (imageHeight-originalImageHeight)/2))

	converter.Image = newImage
	converter.ImageWidth = imageWidth
	converter.ImageHeight = imageHeight
	converter.OriginalImageWidth = originalImageWidth
	converter.OriginalImageHeight = originalImageHeight

	return nil
}

// Tile Extract a single tile from a larger image
func (converter *Converter) Tile(level int, size [2]int, quadrant [2]int, efficient bool) (image.Image, error) {
	scale := int(math.Pow(2, float64(level)))

	// crop out the area of interest first, then scale and copy it
	if efficient {
		inverseSize := [2]float64{float64(converter.ImageHeight) / float64(size[0]*scale), float64(converter.ImageHeight) / float64(size[1]*scale)}
		topLeft := [2]int{int(float64(quadrant[0]) * float64(size[0]) * inverseSize[0]), int(float64(quadrant[1]) * float64(size[1]) * inverseSize[1])}
		bottomRight := [2]int{topLeft[0] + int(float64(size[0])*inverseSize[0]), topLeft[1] + int(float64(size[1])*inverseSize[1])}

		if inverseSize[0] < 1 || inverseSize[1] < 1 {
			return nil, fmt.Errorf("requested zoom level (%d) is too high", level)
		}

		fmt.Printf(" crop %s resize %v \n", fmt.Sprintf("%v %v", topLeft, bottomRight), size)

		zoomed := imaging.Crop(converter.Image, image.Rect(topLeft[0], topLeft[1], bottomRight[0], bottomRight[1]))
		zoomed = imaging.Resize(zoomed, size[0], size[1], resampleFilter)

		return zoomed, nil
	}

	// copy the whole image, scale it and then crop out the area of interest
	newSize := [2]int{size[0] * scale, size[0] * scale}
	topLeft := [2]int{quadrant[0] * size[0], quadrant[1] * size[1]}
	bottomRight := [2]int{topLeft[0] + size[0], topLeft[1] + size[1]}

	if newSize[0] > converter.ImageWidth || newSize[1] > converter.ImageHeight {
		return nil, fmt.Errorf("requested zoom level (%d) is too high", level)
	}

	fmt.Printf("crop(%s).resize(%v)", fmt.Sprintf("%v %v", topLeft, bottomRight), newSize)

	zoomed := imaging.Clone(converter.Image)
	zoomed = imaging.Resize(zoomed, newSize[0], newSize[1], resampleFilter)
	zoomed = imaging.Crop(zoomed, image.Rect(topLeft[0], topLeft[1], bottomRight[0], bottomRight[1]))

	return zoomed, nil
}

// subdivide Recursively subdivide a large image into small tiles
func (converter *Converter) subdivide(level int, size [2]int, quadrant [2]int, efficient bool, imageQuality int, output string) (image.Image, error) {
	if converter.ImageWidth <= size[0]*int(math.Pow(2, float64(level))) {
		outImg, err := converter.Tile(level, size, quadrant, efficient)
		if err != nil {
			return nil, err
		}

		filePath := fmt.Sprintf(output, level, quadrant[0], quadrant[1])
		if err := os.MkdirAll(path.Dir(filePath), os.ModePerm); err != nil {
			return nil, fmt.Errorf("create output directory: %v\n", err)
		}
		if err := SaveJPG(filePath, outImg, imageQuality); err != nil {
			return nil, err
		}

		if level > converter.MaxZoom {
			converter.MaxZoom = level
		}

		converter.TileSize = size

		fmt.Printf("level %d quadrant %v filePath %s", level, quadrant, filePath)

		return outImg, nil
	}

	outImg := imaging.New(size[0]*2, size[1]*2, color.NRGBA{})

	if img, err := converter.subdivide(level+1, size, [2]int{quadrant[0]*2 + 0, quadrant[1]*2 + 0}, efficient, imageQuality, output); err != nil {
		return nil, err
	} else {
		outImg = imaging.Paste(outImg, img, image.Pt(0, 0))

	}

	if img, err := converter.subdivide(level+1, size, [2]int{quadrant[0]*2 + 0, quadrant[1]*2 + 1}, efficient, imageQuality, output); err != nil {
		return nil, err
	} else {
		outImg = imaging.Paste(outImg, img, image.Pt(0, size[1]))
	}

	if img, err := converter.subdivide(level+1, size, [2]int{quadrant[0]*2 + 1, quadrant[1]*2 + 0}, efficient, imageQuality, output); err != nil {
		return nil, err
	} else {
		outImg = imaging.Paste(outImg, img, image.Pt(size[0], 0))
	}

	if img, err := converter.subdivide(level+1, size, [2]int{quadrant[0]*2 + 1, quadrant[1]*2 + 1}, efficient, imageQuality, output); err != nil {
		return nil, err
	} else {
		outImg = imaging.Paste(outImg, img, image.Pt(size[0], size[1]))
	}

	outImg = imaging.Resize(outImg, size[0], size[1], resampleFilter)

	filePath := fmt.Sprintf(output, level, quadrant[0], quadrant[1])
	if err := os.MkdirAll(path.Dir(filePath), os.ModePerm); err != nil {
		return nil, fmt.Errorf("create output directory: %v\n", err)
	}
	if err := SaveJPG(filePath, outImg, imageQuality); err != nil {
		return nil, err
	}

	return outImg, nil
}

// Execute a large image into small tiles
func (converter *Converter) Execute(size [2]int, efficient bool, imageQuality int, output string) error {
	_, err := converter.subdivide(0, size, [2]int{0, 0}, efficient, imageQuality, output)
	if err != nil {
		return err
	}

	return nil
}
