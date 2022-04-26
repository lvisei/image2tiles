package image2tiles_test

import (
	"testing"

	"github.com/lvisei/image2tiles"
)

func TestPrepare(t *testing.T) {
	image2tiles := image2tiles.New()
	if err := image2tiles.Prepare(); err != nil {
		t.Fatal("prepare", err)
	}
}
