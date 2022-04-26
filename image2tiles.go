package image2tiles

import (
	"image"
	"log"
)

type Image2tiles struct {
	Image    image.Image
	OrigSize [2]int
	NewSize  [2]int
	MaxZoom  int
	TileSize [2]int
}

func NewConverter() *Image2tiles {
	return &Image2tiles{}
}

func (image2tiles *Image2tiles) Prepare(imageName string) error {
	img, err := LoadImage(imageName)
	if err != nil {
		log.Fatal("open image failed : " + err.Error())
	}

	image2tiles.OrigSize = [2]int{img.Bounds().Size().X, img.Bounds().Size().Y}
	image2tiles.NewSize = [2]int{1, 1}

	image2tiles.Image = image.NewRGBA(image.Rect(0, 0, image2tiles.NewSize[0], image2tiles.NewSize[1]))

	return nil
}

func (image2tiles *Image2tiles) Tile() {

}

func (image2tiles *Image2tiles) Subdivide() {

}
