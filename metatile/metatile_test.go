package metatile

import (
	"fmt"

	"github.com/tierpod/go-osm/tile"
)

func ExampleMetatile_Filepath() {
	hashes := [5]int{1, 2, 3, 4, 5}
	mt := Metatile{Zoom: 10, Hashes: hashes, Style: "mapname", X: 0, Y: 0}

	filepath := mt.Filepath("")
	fmt.Println(filepath)

	filepath = mt.Filepath("/var/lib/mod_tile")
	fmt.Println(filepath)

	// Output:
	// mapname/10/5/4/3/2/1.meta
	// /var/lib/mod_tile/mapname/10/5/4/3/2/1.meta
}

func ExampleMetatile_Size() {
	zooms := []int{1, 2, 3, 8}
	for _, zoom := range zooms {
		mt := Metatile{Zoom: zoom}
		fmt.Println(zoom, mt.Size())
	}

	// Output:
	// 1 2
	// 2 4
	// 3 8
	// 8 8
}

func ExampleXYToIndex() {
	xx := []int{0, 1}
	yy := []int{0, 1}

	for x := range xx {
		for y := range yy {
			offset := XYToIndex(x, y)
			fmt.Printf("(%v, %v): %v\n", x, y, offset)
		}
	}

	// Output:
	// (0, 0): 0
	// (0, 1): 1
	// (1, 0): 8
	// (1, 1): 9
}

func ExampleIndexToXY() {
	ii := []int{0, 1, 8, 9}

	for _, i := range ii {
		x, y := IndexToXY(i)
		fmt.Printf("%v: (%v, %v)\n", i, x, y)
	}

	// Output:
	// 0: (0, 0)
	// 1: (0, 1)
	// 8: (1, 0)
	// 9: (1, 1)
}

func ExampleNewFromURL() {
	urls := []string{
		"map/10/0/0/33/180/128.meta",
		"/var/lib/mod_tile/map/10/0/0/33/180/128.meta",
		"http://localhost:8080/maps/map/10/0/0/33/180/128.meta",
		"map/ZOOM/0/0/33/180/128.meta",
	}

	for _, url := range urls {
		mt, err := NewFromURL(url)
		if err != nil {
			fmt.Printf("error: %v\n", err)
			continue
		}
		fmt.Println(mt)
	}

	// Output:
	// Metatile{Zoom:10 Hashes:[128 180 33 0 0] Style:map Ext:.meta X:696 Y:320}
	// Metatile{Zoom:10 Hashes:[128 180 33 0 0] Style:map Ext:.meta X:696 Y:320}
	// Metatile{Zoom:10 Hashes:[128 180 33 0 0] Style:map Ext:.meta X:696 Y:320}
	// error: could not parse url string to Metatile struct
}

func ExampleNewFromTile() {
	t := tile.Tile{Zoom: 10, X: 697, Y: 321, Ext: ".png"}
	mt := NewFromTile(t)
	fmt.Println(mt)
	fmt.Println(mt.Filepath("/var/lib/mod_tile"))

	// Output:
	// Metatile{Zoom:10 Hashes:[128 180 33 0 0] Style: Ext:.meta X:696 Y:320}
	// /var/lib/mod_tile/10/0/0/33/180/128.meta
}
