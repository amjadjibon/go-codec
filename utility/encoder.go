package utility

import (
	"encoding/gob"

	"github.com/amjadjibon/go-codec/base32"
	"github.com/amjadjibon/go-codec/base64"
	"github.com/amjadjibon/go-codec/errors"
	"github.com/amjadjibon/go-codec/hex"
	"github.com/amjadjibon/go-codec/iface"
	"github.com/amjadjibon/go-codec/registry"
)

func init() {
	gob.Register(base32.EncoderBASE32{})
	gob.Register(base64.EncoderBASE64{})
	gob.Register(hex.EncoderHEX{})
}

func GetEncoder(name string) (iface.Encoder, error) {
	var e = registry.EncoderRegistry().GetEncoder(name)
	if e == nil {
		return nil, errors.ErrEncoderNotFound
	}
	return e, nil
}
