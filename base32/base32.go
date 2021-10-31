package base32

import (
	"encoding/base32"

	"github.com/amjadjibon/go-codec/constant"
	"github.com/amjadjibon/go-codec/registry"
)

type EncoderBASE32 struct{}

func (e EncoderBASE32) Name() string {
	return constant.BASE32
}

func (e EncoderBASE32) Version() string {
	return constant.Version
}

func (e EncoderBASE32) Encode(src []byte) ([]byte, error) {
	var encoding = base32.StdEncoding
	var dst = make([]byte, encoding.EncodedLen(len(src)))
	encoding.Encode(dst, src)
	return dst, nil
}

func (e EncoderBASE32) Decode(src []byte) ([]byte, error) {
	var encoding = base32.StdEncoding
	var dst = make([]byte, encoding.DecodedLen(len(src)))
	n, err := encoding.Decode(dst, src)
	return dst[:n], err
}

func init() {
	registry.EncoderRegistry().AddEncoder(constant.BASE32, EncoderBASE32{})
}
