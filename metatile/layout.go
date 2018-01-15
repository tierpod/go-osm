package metatile

import (
	"fmt"
	"io"
	"math"
)

const (
	// MaxCount is the maximum count of tiles in metatile.
	MaxCount = 1000
	// MaxEntrySize is the maximum size of metatile entry in bytes.
	MaxEntrySize = 2000000
)

type metaEntry struct {
	Offset int32
	Size   int32
}

// decode tile data for this entry
func (e metaEntry) decode(r io.ReadSeeker) ([]byte, error) {
	if e.Size > MaxEntrySize {
		return nil, fmt.Errorf("metaEntry.decode: entry size (%v) > MaxEntrySize", e.Size)
	}

	_, err := r.Seek(int64(e.Offset), 0)
	if err != nil {
		return nil, fmt.Errorf("metaEntry.decode: %v", err)
	}

	buf := make([]byte, e.Size)
	n, err := r.Read(buf)
	if err != nil {
		return nil, fmt.Errorf("metaEntry.decode: %v", err)
	}

	if int32(n) != e.Size {
		return nil, fmt.Errorf("metaEntry.decode: invalid tile size: %v != %v", n, e.Size)
	}

	return buf, nil
}

type metaLayout struct {
	Magic   []byte
	Count   int32
	X, Y, Z int32
	Index   []metaEntry
}

func (m metaLayout) size() int32 {
	return int32(math.Sqrt(float64(m.Count)))
}

func (m metaLayout) tileIndex(x, y int32) int32 {
	return (x-m.X)*m.size() + (y - m.Y)
}
