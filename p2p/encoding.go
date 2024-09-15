package p2p

import (
	"encoding/gob"
	"io"
)

type Decoder interface {
	Decode(io.Reader, any) error
}

type GoBDecoder struct{}

func (dec GoBDecoder) Decode(r io.Reader, v any) error {
	return gob.NewDecoder(r).Decode(v)
}
