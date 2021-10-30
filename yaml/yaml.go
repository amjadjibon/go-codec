package yaml

import (
	"github.com/amjadjibon/go-codec/constant"
	"github.com/amjadjibon/go-codec/registry"
	"gopkg.in/yaml.v3"
)

type CodecYAML struct {}

func (c CodecYAML) Name() string {
	return constant.YAML
}

func (c CodecYAML) Version() string {
	return constant.Version
}

func (c CodecYAML) Marshal(v interface{}) ([]byte, error) {
	return yaml.Marshal(v)
}

func (c CodecYAML) Unmarshal(b []byte, v interface{}) error {
	return yaml.Unmarshal(b, v)
}

func init() {
	registry.CodecRegistry().AddCodec(constant.YAML, CodecYAML{})
}
