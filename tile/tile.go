// Package tile provides Tile struct with coordinates based on OSM file naming convention:
// https://wiki.openstreetmap.org/wiki/Slippy_map_tilenames.
package tile

import (
	"fmt"
	"math"
	"path"
	"regexp"
	"strconv"
)

// Tile describes osm tile coordinates.
type Tile struct {
	Zoom  int
	X, Y  int
	Ext   string
	Style string
}

// Data is the Tile data.
type Data []byte

func (t Tile) String() string {
	return fmt.Sprintf("Tile{Zoom:%v X:%v Y:%v Ext:%v Style:%v}", t.Zoom, t.X, t.Y, t.Ext, t.Style)
}

// Filepath returns tile file path, based on basedir and coordinates.
func (t Tile) Filepath(basedir string) string {
	zoom := strconv.Itoa(t.Zoom)
	x := strconv.Itoa(t.X)
	y := strconv.Itoa(t.Y)
	return path.Join(basedir, t.Style, zoom, x, y+t.Ext)
}

// New creates new Tile.
func New(zoom, x, y int, ext, style string) Tile {
	return Tile{
		Zoom:  zoom,
		X:     x,
		Y:     y,
		Ext:   ext,
		Style: style,
	}
}

var reTile = regexp.MustCompile(`(\w+)/(\d+)/(\d+)/(\d+)(\.\w+)`)

// NewFromURL creates Tile from url.
func NewFromURL(url string) (Tile, error) {
	items := reTile.FindStringSubmatch(url)
	if len(items) == 0 {
		return Tile{}, fmt.Errorf("could not parse url string to Tile struct")
	}

	// we can ignore errors because regexp contains `\d+`
	zoom, _ := strconv.Atoi(items[2])
	x, _ := strconv.Atoi(items[3])
	y, _ := strconv.Atoi(items[4])

	return Tile{
		Style: items[1],
		Zoom:  zoom,
		X:     x,
		Y:     y,
		Ext:   items[5],
	}, nil
}

// NewFromLatLong creates Tile from latitude and longtitude coordinates.
func NewFromLatLong(lat, long float64, zoom int) Tile {
	var x, y int
	x = int(math.Floor((long + 180.0) / 360.0 * (math.Exp2(float64(zoom)))))
	y = int(math.Floor((1.0 - math.Log(math.Tan(lat*math.Pi/180.0)+1.0/math.Cos(lat*math.Pi/180.0))/math.Pi) / 2.0 * (math.Exp2(float64(zoom)))))

	return Tile{
		Zoom: zoom,
		X:    x,
		Y:    y,
	}
}
