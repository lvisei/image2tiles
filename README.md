# image2tiles

[![GoDoc](https://godoc.org/github.com/lvisei/image2tiles?status.svg)](https://pkg.go.dev/github.com/lvisei/image2tiles)
[![Go Report Card](https://goreportcard.com/badge/github.com/lvisei/image2tiles)](https://goreportcard.com/report/github.com/lvisei/image2tiles)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

image2tiles is a tool to cut large image into square tiles, to be used for a interactive tiled viewer.

There are many high-performant viewing of large image available，but cropping and slicing images may be trouble. image2tiles support custom tile size and creates all tiles to be used by map tool like [leaflet](https://lvisei.github.io/image2tiles/leaflet.html), [openlayers](https://lvisei.github.io/image2tiles/openlayers.html) or Other.

## Command Line

### How to get

#### Download

You can download from GitHub [releases](https://github.com/lvisei/image2tiles/releases).

For example download file:

- windows: `**_windows_x86_64.zip`
- maxOS x86: `**_darwin_x86_64.tar.gz`
- maxOS M1: `**_darwin_arm64.tar.gz`

#### Build from source

```
git clone https://github.com/lvisei/image2tiles
cd cmd/image2tiles && go install
```

### Usage

```bash
image2tiles -f image.png -s 512
```

Options flags:

```
image2tiles

Usage: image2tiles -f <filename> [-s] [-t] [-b] [-o]

Options:
  -b string
        The background color to be used for the tiles (default "#ffffff00")
  -f string
        Image filename to be convert
  -o string
        Output file pattern (default "out/%d/%d-%d.jpg")
  -s int
        The tile height/width (default 512)
```

## Library

### How to get

```bash
go get github.com/lvisei/image2tiles
```

### Usage

Image into small single tile

```go
package main

import (
	"fmt"
	"github.com/lvisei/image2tiles"
)

func main() {
  converter := image2tiles.NewConverter()

  if err := converter.Prepare("image.png", "#fff"); err != nil {
      fmt.Println(err)
  }

  if img, err := converter.Tile(0, [2]int{256, 256}, [2]int{0, 0}, true); err!=nil {
  	fmt.Println(err)
  }else {
  	image2tiles.SaveJPG("out/0-0-0.jpg", img, 75)
  }
}

```

Image into multiple small tiles

```go
package main

import (
	"fmt"
	"github.com/lvisei/image2tiles"
)

func main() {
  converter := image2tiles.NewConverter()

  if err := converter.Prepare("image.png", "#00000000"); err != nil {
      fmt.Println(err)
  }

  if err := converter.Execute([2]int{256, 256}, true, 75, "out/level-%d/%d-%d.jpg"); err != nil {
  	fmt.Println(err)
  }

  fmt.Println(converter.MaxZoom, converter.TileSize)
}

```

## LICENSE

[MIT](./LICENSE)
