package image2tiles_test

import (
	"testing"

	"github.com/lvisei/image2tiles"
)

func TestPrepare(t *testing.T) {
	converter := image2tiles.NewConverter()
	imageFilename := "testdata/earth_5568_3712.jpg"

	if err := converter.Prepare(imageFilename, "#fff"); err != nil {
		t.Fatal("prepare", err)
	}
}

func TestTile(t *testing.T) {
	converter := image2tiles.NewConverter()
	imageFilename := "testdata/earth_5568_3712.jpg"

	if err := converter.Prepare(imageFilename, "#00000000"); err != nil {
		t.Fatal("prepare", err)
	}

	if _, err := converter.Tile(0, [2]int{512, 512}, [2]int{0, 0}, true); err != nil {
		t.Fatal("subdivide", err)
	}
}

func TestExecute(t *testing.T) {
	converter := image2tiles.NewConverter()
	imageFilename := "testdata/earth_5568_3712.jpg"

	if err := converter.Prepare(imageFilename, "#00000000"); err != nil {
		t.Fatal("prepare", err)
	}
	if err := converter.Execute([2]int{512, 512}, true, 75, "docs/tiles/%d/%d-%d.png"); err != nil {
		t.Fatal("subdivide", err)
	}
}
