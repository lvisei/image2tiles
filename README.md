# image2tiles

image2tiles is a tool to cut large image into square tiles, to be used for a interactive tiled viewer.

## Usage

```
image2tiles -i image.png -s 256
```

Flags:
-b the background color to be used for the tiles (default "#FFF")
-i image file name
-o output directory (default "out")
-s the tile height/width (default 256)
-t template filename pattern (default "-%d-%d-%d")
