package gob

import (
	"bytes"
	"encoding/gob"
	"github.com/amjadjibon/go-codec/constant"
	"github.com/amjadjibon/go-codec/registry"
)

type CodecGOB struct{}

func (c CodecGOB) Name() string {
	return constant.GOB
}

func (c CodecGOB) Version() string {
	return constant.Version
}

func (c CodecGOB) Marshal(v interface{}) ([]byte, error) {
	var buffer bytes.Buffer
	var encoder = gob.NewEncoder(&buffer)
	if err := encoder.Encode(v); err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func (c CodecGOB) Unmarshal(data []byte, v interface{}) error {
	var reader = bytes.NewReader(data)
	var decoder = gob.NewDecoder(reader)
	return decoder.Decode(v)
}

func init() {
	registry.CodecRegistry().AddCodec(constant.GOB, CodecGOB{})
}
