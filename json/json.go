package json

import (
	"encoding/json"
	"github.com/amjadjibon/go-codec/constant"
	"github.com/amjadjibon/go-codec/registry"
)

type CodecJSON struct {}

func (c CodecJSON) Name() string {
	return constant.JSON
}

func (c CodecJSON) Version() string {
	return constant.Version
}

func (c CodecJSON) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func (c CodecJSON) Unmarshal(b []byte, v interface{}) error {
	return json.Unmarshal(b, v)
}

func init() {
	registry.CodecRegistry().AddCodec(constant.JSON, CodecJSON{})
}
