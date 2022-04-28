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

	if _, err := converter.Tile(0, [2]int{512, 512}, [2]int{0, 0}, true); err != nil {
		t.Fatal("subdivide", err)
	}
}

func TestExecute(t *testing.T) {
	converter := image2tiles.NewConverter()
	imageFilename := "testdata/earth_5568 × 3712.jpg"

	if err := converter.Prepare(imageFilename, "#fff"); err != nil {
		t.Fatal("prepare", err)
	}
	if err := converter.Execute([2]int{512, 512}, "out/%d/%d-%d.jpg"); err != nil {
		t.Fatal("subdivide", err)
	}
}
