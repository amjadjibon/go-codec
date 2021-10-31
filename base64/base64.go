package base64

import (
	"encoding/base64"

	"github.com/amjadjibon/go-codec/constant"
	"github.com/amjadjibon/go-codec/registry"
)

type EncoderBASE64 struct{}

func (e EncoderBASE64) Name() string {
	return constant.BASE64
}

func (e EncoderBASE64) Version() string {
	return constant.Version
}

func (e EncoderBASE64) Encode(src []byte) ([]byte, error) {
	var encoding = base64.StdEncoding
	var dst = make([]byte, encoding.EncodedLen(len(src)))
	encoding.Encode(dst, src)
	return dst, nil
}

func (e EncoderBASE64) Decode(src []byte) ([]byte, error) {
	var encoding = base64.StdEncoding
	var dst = make([]byte, encoding.DecodedLen(len(src)))
	n, err := encoding.Decode(dst, src)
	return dst[:n], err
}

func init() {
	registry.EncoderRegistry().AddEncoder(constant.BASE64, EncoderBASE64{})
}
