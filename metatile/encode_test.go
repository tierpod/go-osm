package metatile

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"

	"github.com/tierpod/go-osm/tile"
)

func TestEncodeWrite(t *testing.T) {
	// create temporary file for writing
	tmpFile, err := ioutil.TempFile("testdata", "tmpfile")
	if err != nil {
		t.Fatalf("got error: %v", err)
	}
	defer func() {
		tmpFile.Close()
		os.Remove(tmpFile.Name())
	}()

	// read from valid metatile file
	f, err := os.Open("testdata/0.meta")
	if err != nil {
		t.Fatalf("got error: %v", err)
	}

	// write to temporary file
	decoder, err := NewDecoder(f)
	if err != nil {
		t.Fatalf("got error: %v", err)
	}

	data, err := decoder.Tiles()
	if err != nil {
		t.Fatalf("got error: %v", err)
	}

	tile := tile.New(1, 1, 1, "", "")
	metatile := NewFromTile(tile)
	metatile.EncodeWrite(tmpFile, data)

	// read both files and compare
	b1, err := ioutil.ReadFile("testdata/0.meta")
	if err != nil {
		t.Fatalf("got error: %v", err)
	}

	b2, err := ioutil.ReadFile(tmpFile.Name())
	if err != nil {
		t.Fatalf("got error: %v", err)
	}

	if !bytes.Equal(b1, b2) {
		t.Fatalf("files not equal")
	}
}
