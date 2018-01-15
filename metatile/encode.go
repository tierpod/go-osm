package metatile

import (
	"encoding/binary"
	"fmt"
	"io"
)

// encodeHeader encodes ml and writes it to w.
func encodeHeader(w io.Writer, ml *metaLayout) error {
	endian := binary.LittleEndian
	var err error
	if err = binary.Write(w, endian, ml.Magic); err != nil {
		return err
	}
	if err = binary.Write(w, endian, ml.Count); err != nil {
		return err
	}
	if err = binary.Write(w, endian, ml.X); err != nil {
		return err
	}
	if err = binary.Write(w, endian, ml.Y); err != nil {
		return err
	}
	if err = binary.Write(w, endian, ml.Z); err != nil {
		return err
	}
	for _, ent := range ml.Index {
		if err = binary.Write(w, endian, ent); err != nil {
			return err
		}
	}
	return nil
}

// Encode encodes tiles data to metatile layout and writes it to w.
func (m Metatile) Encode(w io.Writer, data Data) error {
	mSize := MaxSize * MaxSize

	if len(data) < Area {
		return fmt.Errorf("Metatile.Write: data size: %v < %v", len(data), mSize)
	}

	ml := &metaLayout{
		Magic: []byte{'M', 'E', 'T', 'A'},
		Count: int32(mSize),
		X:     int32(m.X),
		Y:     int32(m.Y),
		Z:     int32(m.Zoom),
	}
	offset := int32(20 + 8*mSize)

	// calculate offsets and sizes
	for i := 0; i < mSize; i++ {
		tile := data[i]
		s := int32(len(tile))
		if s > MaxEntrySize {
			return fmt.Errorf("Metatile.Write: entry size (%v) > MaxEntrySize", s)
		}

		ml.Index = append(ml.Index, metaEntry{
			Offset: offset,
			Size:   s,
		})
		offset += s
	}

	// fmt.Printf("%+v\n", ml)

	// encode and write headers
	if err := encodeHeader(w, ml); err != nil {
		return fmt.Errorf("encodeMetatile: %v", err)
	}

	// encode and write data
	for i := 0; i < len(data); i++ {
		tile := data[i]

		if _, err := w.Write(tile); err != nil {
			return fmt.Errorf("encodeMetatile: %v", err)
		}
	}

	return nil
}
