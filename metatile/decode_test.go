package metatile

import (
	"fmt"
	"os"
	"testing"

	"github.com/tierpod/go-osm/point"
)

func TestDecoder(t *testing.T) {
	f, err := os.Open("testdata/0.meta")
	if err != nil {
		t.Fatalf("got error %v", err)
	}
	defer f.Close()

	d, err := NewDecoder(f)
	if err != nil {
		t.Fatalf("got error %v", err)
	}

	// check header
	l := d.Layout()
	if l.Z != 0 || l.X != 0 || l.Y != 1 || l.Count != 64 {
		t.Errorf("got wrong header: z:%v(0) x:%v(0) y:%v(0) count:%v(64)", l.Z, l.X, l.Y, l.Count)
	}

	// check entries table
	e := d.Entries()
	validEntries := []struct {
		index int
		entry Entry
	}{
		{0, Entry{532, 25093}},
		{1, Entry{25625, 11330}},
		{2, Entry{36955, 0}},
		{8, Entry{36955, 26298}},
		{9, Entry{63253, 10439}},
		{10, Entry{73692, 0}},
		{20, Entry{73692, 0}},
		{40, Entry{73692, 0}},
		{63, Entry{73692, 0}},
	}

	for _, tt := range validEntries {
		if tt.entry != e[tt.index] {
			t.Errorf("got wrong Entry value")
		}
	}

	data, err := d.Tile(1, 1)
	if err != nil {
		t.Errorf("got error %v", err)
	}

	if len(data) != 10439 {
		t.Errorf("got wrong data length: %v", len(data))
	}
}

func ExampleDecoder_Tile() {
	f, err := os.Open("testdata/0.meta")
	if err != nil {
		fmt.Printf("got error %v\n", err)
	}
	defer f.Close()

	d, err := NewDecoder(f)
	if err != nil {
		fmt.Printf("got error: %v\n", err)
		return
	}

	// print header and size
	l := d.Layout()
	size := d.Size()
	fmt.Println(l.Z, l.X, l.Y, l.Count, size)

	// tile exist in metatile and has data
	data, err := d.Tile(1, 1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(len(data))

	// tile exist in metatile and has no data
	_, err = d.Tile(6, 6)
	if err != nil {
		fmt.Println(err)
	}

	// tile does not exist in metatile
	_, err = d.Tile(7, 8)
	if err != nil {
		fmt.Println(err)
	}

	// Output:
	// 0 0 1 64 8
	// 10439
	// decoder: empty data
	// decoder: invalid index
}

func ExampleDecoder_Tiles() {
	f, err := os.Open("testdata/0.meta")
	if err != nil {
		fmt.Printf("got error %v\n", err)
	}
	defer f.Close()

	d, err := NewDecoder(f)
	if err != nil {
		fmt.Printf("got error: %v\n", err)
		return
	}

	data, err := d.Tiles()
	if err != nil {
		fmt.Println(err)
		return
	}

	for k, v := range data {
		if len(v) != 0 {
			fmt.Println(k, len(v))
		}
	}

	// Output:
	// 0 25093
	// 1 11330
	// 8 26298
	// 9 10439
}

func ExampleDecoder_TilesMap() {
	f, err := os.Open("testdata/0.meta")
	if err != nil {
		fmt.Printf("got error %v\n", err)
	}
	defer f.Close()

	d, err := NewDecoder(f)
	if err != nil {
		fmt.Printf("got error: %v\n", err)
		return
	}

	data, err := d.TilesMap()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(len(data))

	points := []point.ZXY{
		point.ZXY{Z: 1, X: 0, Y: 0},
		point.ZXY{Z: 1, X: 0, Y: 1},
		point.ZXY{Z: 1, X: 1, Y: 0},
		point.ZXY{Z: 1, X: 1, Y: 1},
		point.ZXY{Z: 99, X: 99, Y: 99},
	}
	for _, p := range points {
		fmt.Printf("%v %v\n", p, len(data[p]))
	}

	// Output:
	// 4
	// {Z:1 X:0 Y:0} 25093
	// {Z:1 X:0 Y:1} 11330
	// {Z:1 X:1 Y:0} 26298
	// {Z:1 X:1 Y:1} 10439
	// {Z:99 X:99 Y:99} 0
}
