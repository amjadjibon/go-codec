package utility

import (
	"encoding/gob"

	"github.com/amjadjibon/go-codec/errors"
	codecGOB "github.com/amjadjibon/go-codec/gob"
	"github.com/amjadjibon/go-codec/iface"
	"github.com/amjadjibon/go-codec/json"
	"github.com/amjadjibon/go-codec/msgpack"
	"github.com/amjadjibon/go-codec/proto"
	"github.com/amjadjibon/go-codec/registry"
	"github.com/amjadjibon/go-codec/xml"
	"github.com/amjadjibon/go-codec/yaml"
)

func init() {
	gob.Register(json.CodecJSON{})
	gob.Register(yaml.CodecYAML{})
	gob.Register(proto.CodecPROTO{})
	gob.Register(xml.CodecXML{})
	gob.Register(msgpack.CodecMSGPACK{})
	gob.Register(codecGOB.CodecGOB{})
}

func GetCodec(name string) (iface.Codec, error) {
	var c = registry.CodecRegistry().GetCodec(name)
	if c == nil {
		return nil, errors.ErrCodecNotFound
	}
	return c, nil
}
