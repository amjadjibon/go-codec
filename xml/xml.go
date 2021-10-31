package xml

import (
	"encoding/xml"

	"github.com/amjadjibon/go-codec/constant"
	"github.com/amjadjibon/go-codec/registry"
)

type CodecXML struct{}

func (c CodecXML) Name() string {
	return constant.XML
}

func (c CodecXML) Version() string {
	return constant.Version
}

func (c CodecXML) Marshal(v interface{}) ([]byte, error) {
	return xml.Marshal(v)
}

func (c CodecXML) Unmarshal(b []byte, v interface{}) error {
	return xml.Unmarshal(b, v)
}

func init() {
	registry.CodecRegistry().AddCodec(constant.XML, CodecXML{})
}
