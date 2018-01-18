package metatile

import (
	"encoding/binary"
	"fmt"
	"io"
)

// Decoder is the metatile file decoder wrapper.
type Decoder struct {
	ml *metaLayout
	r  io.ReadSeeker
}

// NewDecoder reads metatile from r, parses layout and returns Decoder.
func NewDecoder(r io.ReadSeeker) (*Decoder, error) {
	ml, err := decodeHeader(r)
	if err != nil {
		return nil, err
	}

	return &Decoder{
		ml: ml,
		r:  r,
	}, nil
}

// Header returns metatile header: x, y, z coordinates and count of tiles.
func (m *Decoder) Header() (x, y, z, count int32) {
	return m.ml.X, m.ml.Y, m.ml.Z, m.ml.Count
}

// Entries returns metatile index table (offsets and sizes).
func (m *Decoder) Entries() []metaEntry {
	return m.ml.Index
}

// Tile reads data for tile with (x, y) coordinates.
func (m *Decoder) Tile(x, y int) ([]byte, error) {
	i, err := m.ml.tileIndex(int32(x), int32(y))
	if err != nil {
		return nil, err
	}

	entry := m.ml.Index[i]
	data, err := entry.decode(m.r)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// Tiles reads data for all tiles in metatile and returns none-empty data.
func (m *Decoder) Tiles() ([][]byte, error) {
	var tiles [][]byte

	for i := range m.ml.Index {
		entry := m.ml.Index[i]
		data, err := entry.decode(m.r)
		if err != nil {
			// skip tiles with empty data
			if err == ErrEmptyData {
				continue
			}
			return nil, err
		}

		tiles = append(tiles, data)
	}

	return tiles, nil
}

// decodeHeader reads metatile from r and decodes header to metaLayout struct.
func decodeHeader(r io.Reader) (*metaLayout, error) {
	endian := binary.LittleEndian
	ml := new(metaLayout)

	ml.Magic = make([]byte, 4)
	err := binary.Read(r, endian, &ml.Magic)
	if err != nil {
		return nil, err
	}
	if ml.Magic[0] != 'M' || ml.Magic[1] != 'E' || ml.Magic[2] != 'T' || ml.Magic[3] != 'A' {
		return nil, fmt.Errorf("invalid Magic field: %v", ml.Magic)
	}

	if err = binary.Read(r, endian, &ml.Count); err != nil {
		return nil, err
	}
	if err = binary.Read(r, endian, &ml.X); err != nil {
		return nil, err
	}
	if err = binary.Read(r, endian, &ml.Y); err != nil {
		return nil, err
	}
	if err = binary.Read(r, endian, &ml.Z); err != nil {
		return nil, err
	}

	for i := int32(0); i < ml.Count; i++ {
		var entry metaEntry
		if err = binary.Read(r, endian, &entry); err != nil {
			return nil, err
		}
		ml.Index = append(ml.Index, entry)
	}

	return ml, nil
}
