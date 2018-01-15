package metatile

import (
	"os"
	"testing"
)

func TestDecodeTile(t *testing.T) {
	f, err := os.Open("testdata/0.meta")
	if err != nil {
		t.Fatalf("got error %v", err)
	}
}
