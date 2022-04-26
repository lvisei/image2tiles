package image2tiles_test

import (
	"testing"

	"github.com/lvisei/image2tiles"
)

func TestPrepare(t *testing.T) {
	converter := image2tiles.NewConverter()
	imageFilename := "testdata/earth_5568 × 3712.jpg"

	if err := converter.Prepare(imageFilename); err != nil {
		t.Fatal("prepare", err)
	}
}
