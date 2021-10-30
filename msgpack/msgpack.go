package msgpack

import (
	"github.com/amjadjibon/go-codec/constant"
	"github.com/amjadjibon/go-codec/registry"
	"github.com/vmihailenco/msgpack/v5"
)

type CodecMSGPACK struct {}

func (c CodecMSGPACK) Name() string {
	return constant.MSGPACK
}

func (c CodecMSGPACK) Version() string {
	return constant.Version
}

func (c CodecMSGPACK) Marshal(v interface{}) ([]byte, error) {
	return msgpack.Marshal(v)
}

func (c CodecMSGPACK) Unmarshal(b []byte, v interface{}) error {
	return msgpack.Unmarshal(b, v)
}

func init() {
	registry.CodecRegistry().AddCodec(constant.MSGPACK, CodecMSGPACK{})
}
