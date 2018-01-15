package metatile

import (
	"encoding/binary"
	"fmt"
	"io"
)

// decodeHeader reads metatile data from r and decodes header to metaLayout struct.
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
	if ml.Count > MaxCount {
		return nil, fmt.Errorf("Count > MaxCount (Count = %v)", ml.Count)
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

// DecodeTile reads metatile data from r, decode tile data with (x, y) coordinates and return data.
func DecodeTile(r io.ReadSeeker, x, y int) ([]byte, error) {
	ml, err := decodeHeader(r)
	if err != nil {
		return nil, fmt.Errorf("DecodeTile: %v", err)
	}

	i := ml.tileIndex(int32(x), int32(y))
	if i >= ml.Count {
		return nil, fmt.Errorf("DecodeTile: invalid index %v/%v", i, ml.Count)
	}

	entry := ml.Index[i]
	data, err := entry.decode(r)
	if err != nil {
		return nil, fmt.Errorf("DecodeTile: %v", err)
	}

	return data, nil
}

// DecodeTileTo reads metatile data from r, decode tile data with (x, y) coordinates and writes it to w.
// func DecodeTileTo(w io.Writer, r io.ReadSeeker, x, y int) error {
// 	data, err := DecodeTile(r, x, y)
// 	if err != nil {
// 		return err
// 	}

// 	io.Copy(w, bytes.NewReader(data))
// 	return nil
// }

// DecodeTiles reads metatile data from r and decodes all tiles data.
func DecodeTiles(r io.ReadSeeker) ([][]byte, error) {
	var tiles [][]byte

	ml, err := decodeHeader(r)
	if err != nil {
		return nil, fmt.Errorf("DecodeTiles: %v", err)
	}

	for i := range ml.Index {
		entry := ml.Index[i]
		data, err := entry.decode(r)
		if err != nil {
			return nil, fmt.Errorf("DecodeTiles: %v", err)
		}

		tiles = append(tiles, data)
	}

	return tiles, nil
}
