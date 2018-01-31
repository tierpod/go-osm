package metatile

import (
	"fmt"

	"github.com/tierpod/go-osm/tile"
)

func ExampleMetatile_Filepath() {
	hashes := [5]int{1, 2, 3, 4, 5}
	mt := Metatile{Zoom: 10, hashes: hashes, Style: "mapname", X: 0, Y: 0}
	fmt.Println(mt.Filepath(""))
	fmt.Println(mt.Filepath("/var/lib/mod_tile"))

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
	// Metatile{Zoom:10 X:696 Y:320 Style:map Ext:.meta}
	// Metatile{Zoom:10 X:696 Y:320 Style:map Ext:.meta}
	// Metatile{Zoom:10 X:696 Y:320 Style:map Ext:.meta}
	// error: could not parse url string to Metatile struct
}

func ExampleNew() {
	mt := New(10, 696, 320, "mapname")
	fmt.Println(mt)
	fmt.Println(mt.Filepath("/var/lib/mod_tile"))

	// Output:
	// Metatile{Zoom:10 X:696 Y:320 Style:mapname Ext:.meta}
	// /var/lib/mod_tile/mapname/10/0/0/33/180/128.meta
}

func ExampleNewFromTile() {
	t := tile.Tile{Zoom: 10, X: 697, Y: 321, Ext: ".png"}
	mt := NewFromTile(t)
	fmt.Println(mt)
	fmt.Println(mt.Filepath("/var/lib/mod_tile"))

	// Output:
	// Metatile{Zoom:10 X:696 Y:320 Style: Ext:.meta}
	// /var/lib/mod_tile/10/0/0/33/180/128.meta
}

func ExampleMetatile_XYBox() {
	mt := New(1, 1, 1, "")
	fmt.Println(mt.XYBox())

	mt = New(10, 697, 321, "")
	fmt.Println(mt.XYBox())

	// Output:
	// [0 1] [0 1]
	// [696 697 698 699 700 701 702 703] [320 321 322 323 324 325 326 327]
}
