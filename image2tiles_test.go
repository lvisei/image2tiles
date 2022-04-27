package image2tiles_test

import (
	"testing"

	"github.com/lvisei/image2tiles"
)

func TestPrepare(t *testing.T) {
	converter := image2tiles.NewConverter()
	imageFilename := "testdata/earth_5568 × 3712.jpg"

	if err := converter.Prepare(imageFilename, "#fff"); err != nil {
		t.Fatal("prepare", err)
	}
}


func TestTile(t *testing.T) {
	converter := image2tiles.NewConverter()
	imageFilename := "testdata/earth_5568 × 3712.jpg"

	if err := converter.Prepare(imageFilename, "#fff"); err != nil {
		t.Fatal("prepare", err)
	}

	if _, err := converter.Tile(0, [2]int{256, 256}, [2]int{0, 0}, true); err!=nil {
		t.Fatal("subdivide", err)
	}
}

func TestSubdivide(t *testing.T) {
	converter := image2tiles.NewConverter()
	imageFilename := "testdata/earth_5568 × 3712.jpg"

	if err := converter.Prepare(imageFilename, "#fff"); err != nil {
		t.Fatal("prepare", err)
	}
	if _, err := converter.Subdivide(0, [2]int{256, 256}, [2]int{0, 0}, "%d-%d-%d.jpg", "out"); err != nil {
		t.Fatal("subdivide", err)
	}
}
