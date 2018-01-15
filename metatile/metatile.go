// Package metatile provides Metatile struct with coordinates and metatile file encoder/decoder.
//
// Metatile format description:
// https://github.com/openstreetmap/mod_tile/blob/master/src/metatile.cpp
package metatile

import (
	"fmt"
	"path"
	"regexp"
	"strconv"

	"github.com/tierpod/go-osm/tile"
)

const (
	// MaxSize is the maximum metatile size.
	MaxSize int = 8
	// Area is the area of metatile.
	Area int = MaxSize * MaxSize
	// Ext is the metatile file extension.
	Ext string = ".meta"
)

// Metatile describes metatile coordinates.
type Metatile struct {
	Zoom   int
	Style  string
	Hashes hashes
	X, Y   int
}

func (m Metatile) String() string {
	return fmt.Sprintf("Metatile{Zoom:%v Hashes:%v Style:%v Ext:%v X:%v Y:%v}", m.Zoom, m.Hashes, m.Style, Ext, m.X, m.Y)
}

// Filepath returns metatile file path, based on basedir and coordinates.
func (m Metatile) Filepath(basedir string) string {
	zoom := strconv.Itoa(m.Zoom)
	h0 := strconv.Itoa(m.Hashes[0]) + Ext
	h1 := strconv.Itoa(m.Hashes[1])
	h2 := strconv.Itoa(m.Hashes[2])
	h3 := strconv.Itoa(m.Hashes[3])
	h4 := strconv.Itoa(m.Hashes[4])
	return path.Join(basedir, m.Style, zoom, h4, h3, h2, h1, h0)
}

// Size return metatile size for current zoom level.
func (m Metatile) Size() int {
	n := int(uint(1) << uint(m.Zoom))
	if n < MaxSize {
		return n
	}
	return MaxSize
}

// Data is the array of tile data with size Area.
type Data [Area]tile.Data

// XYOffset returns offset of tile data inside metatile.
func XYOffset(x, y int) int {
	mask := MaxSize - 1
	return (x&mask)*MaxSize + (y & mask)
}

var reMetatile = regexp.MustCompile(`(\w+)/(\d+)/(\d+)/(\d+)/(\d+)/(\d+)/(\d+)\.meta`)

// NewFromURL creates Metatile from url.
func NewFromURL(url string) (Metatile, error) {
	items := reMetatile.FindStringSubmatch(url)
	if len(items) == 0 {
		return Metatile{}, fmt.Errorf("could not parse url string to Metatile struct")
	}

	zoom, _ := strconv.Atoi(items[2])
	h4, _ := strconv.Atoi(items[3])
	h3, _ := strconv.Atoi(items[4])
	h2, _ := strconv.Atoi(items[5])
	h1, _ := strconv.Atoi(items[6])
	h0, _ := strconv.Atoi(items[7])
	h := hashes{h0, h1, h2, h3, h4}

	x, y := h.XY()

	return Metatile{
		Style:  items[1],
		Zoom:   zoom,
		Hashes: h,
		X:      x,
		Y:      y,
	}, nil
}

// NewFromTile creates Metatile from Tile.
func NewFromTile(t tile.Tile) Metatile {
	h := xyToHashes(t.X, t.Y)
	x, y := h.XY()
	return Metatile{
		Style:  t.Style,
		Zoom:   t.Zoom,
		Hashes: h,
		X:      x,
		Y:      y,
	}
}
