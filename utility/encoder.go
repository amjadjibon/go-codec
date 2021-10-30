package utility

import (
	"encoding/gob"
	"github.com/amjadjibon/go-codec/base32"
	"github.com/amjadjibon/go-codec/base64"
	"github.com/amjadjibon/go-codec/hex"
	"github.com/amjadjibon/go-codec/iface"
	"github.com/amjadjibon/go-codec/registry"
)

func init() {
	gob.Register(base32.EncoderBASE32{})
	gob.Register(base64.EncoderBASE64{})
	gob.Register(hex.EncoderHEX{})
}

func GetEncoder(name string) iface.Encoder {
	var c = registry.EncoderRegistry().GetEncoder(name)
	if c == nil {
		return nil
	}

	v, ok := c.(iface.Encoder)
	if !ok {
		return nil
	}

	return v
}
