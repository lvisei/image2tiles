# image2tiles

image2tiles is a tool to cut large image into square tiles, to be used for a interactive tiled viewer.

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
        The background color to be used for the tiles (default "#FFF")
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

Image into small single tiles

```go
package main

import (
	"fmt"
	"github.com/lvisei/image2tiles"
)

func main() {
  converter := image2tiles.NewConverter()

  if err := converter.Prepare("image.png", "#ffffff"); err != nil {
      fmt.Println(err)
  }

  if img, err := converter.Tile(0, [2]int{256, 256}, [2]int{0, 0}, true); err!=nil {
  	fmt.Println(err)
  }else {
  	image2tiles.SaveJPG("out/0-0-0.jpg", img, 75)
  }
}

```

Image into small multiple tiles

```go
package main

import (
	"fmt"
	"github.com/lvisei/image2tiles"
)

func main() {
  converter := image2tiles.NewConverter()

  if err := converter.Prepare("image.png", "#ffffff"); err != nil {
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
