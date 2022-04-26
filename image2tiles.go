package image2tiles

import (
	"image"
)

type Image2tiles struct {
	Image *image.Image
}

func New() *Image2tiles {
	return &Image2tiles{}
}

func (image2tiles *Image2tiles) Prepare(imageName string) error {
	img, err := loadImage(imageName)
	if err != nil {
		return err
	}

	image2tiles.Image = &img

	return nil
}

func (image2tiles *Image2tiles) Tile() {

}

func (image2tiles *Image2tiles) Subdivide() {

}
