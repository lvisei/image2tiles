package image2tiles_test

import (
	"testing"

	"github.com/lvisei/image2tiles"
	"github.com/stretchr/testify/assert"
)

// TestLoadImage test open different type image file
func TestLoadImage(t *testing.T) {
	assertions := assert.New(t)
	jpgFilename := "testdata/earth_5568*3712.jpg"
	openedImage, err := image2tiles.LoadImage(jpgFilename)
	assertions.True(err == nil, "jpg image format should be supported")
	assertions.True(openedImage != nil, "opened jpg file should not be nil")

	pngFilename := "testdata/spongebob_698*530.png"
	openedImage, err = image2tiles.LoadImage(pngFilename)
	assertions.True(err == nil, "png image format should be supported")
	assertions.True(openedImage != nil, "opened jpg file should not be nil")

	notSupported := "testdata/not_supported_sample_image"
	openedImage, err = image2tiles.LoadImage(notSupported)
	assertions.True(err != nil, "should not open unsupported image")
	assertions.True(openedImage == nil, "not supported image should be nil")
}

// TestLoadImageNotExistsFile test open a not exists file
func TestLoadImageNotExistsFile(t *testing.T) {
	assertions := assert.New(t)
	_, err := image2tiles.LoadImage("not exists")
	assertions.True(err != nil)
}
