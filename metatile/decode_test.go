package metatile

import (
	"fmt"
	"os"
	"testing"
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
	z, x, y, count := d.Header()
	if z != 0 || x != 0 || y != 1 || count != 64 {
		t.Errorf("got wrong header: z:%v(0) x:%v(0) y:%v(0) count:%v(64)", z, x, y, count)
	}

	// check entries table
	e := d.Entries()
	validEntries := []struct {
		index int
		entry metaEntry
	}{
		{0, metaEntry{532, 25093}},
		{1, metaEntry{25625, 11330}},
		{2, metaEntry{36955, 0}},
		{8, metaEntry{36955, 26298}},
		{9, metaEntry{63253, 10439}},
		{10, metaEntry{73692, 0}},
		{20, metaEntry{73692, 0}},
		{40, metaEntry{73692, 0}},
		{63, metaEntry{73692, 0}},
	}

	for _, tt := range validEntries {
		if tt.entry != e[tt.index] {
			t.Errorf("got wrong metaEntry value")
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
		fmt.Println(k, len(v))
	}

	// Output:
	// 0 25093
	// 1 11330
	// 2 26298
	// 3 10439
}
