package hex

import (
	"encoding/hex"

	"github.com/amjadjibon/go-codec/constant"
	"github.com/amjadjibon/go-codec/registry"
)

type EncoderHEX struct{}

func (e EncoderHEX) Name() string {
	return constant.HEX
}

func (e EncoderHEX) Version() string {
	return constant.Version
}

func (e EncoderHEX) Encode(src []byte) ([]byte, error) {
	var dst = make([]byte, hex.EncodedLen(len(src)))
	hex.Encode(dst, src)
	return dst, nil
}

func (e EncoderHEX) Decode(src []byte) ([]byte, error) {
	var dst = make([]byte, hex.EncodedLen(len(src)))
	n, err := hex.Decode(dst, src)
	return dst[:n], err
}

func init() {
	registry.EncoderRegistry().AddEncoder(constant.HEX, EncoderHEX{})
}
