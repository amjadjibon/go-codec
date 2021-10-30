package proto

import (
	"github.com/amjadjibon/go-codec/constant"
	"github.com/amjadjibon/go-codec/registry"
	"google.golang.org/protobuf/proto"
)

type CodecPROTO struct {}

func (c CodecPROTO) Name() string {
	return constant.PROTO
}

func (c CodecPROTO) Version() string {
	return constant.Version
}

func (c CodecPROTO) Marshal(v interface{}) ([]byte, error) {
	return proto.Marshal(v.(proto.Message))
}

func (c CodecPROTO) Unmarshal(b []byte, v interface{}) error {
	return proto.Unmarshal(b, v.(proto.Message))
}

func init() {
	registry.CodecRegistry().AddCodec(constant.PROTO, CodecPROTO{})
}
