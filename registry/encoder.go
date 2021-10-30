package registry

import (
	"github.com/amjadjibon/go-codec/iface"
	"sync"
)

var (
	encoderRegistryOnce sync.Once
	encoderRegistryIns  *encoderRegistry
)

type encoderRegistry struct {
	registry map[string]iface.Encoder
}

func (g *encoderRegistry) setup() {
	g.registry = make(map[string]iface.Encoder)
}

func (g *encoderRegistry) AddEncoder(name string, codec iface.Encoder) {
	g.registry[name] = codec
}

func (g *encoderRegistry) GetEncoder(name string) iface.Encoder {
	return g.registry[name]
}

func EncoderRegistry() *encoderRegistry {
	return encoderRegistryIns
}

func init() {
	encoderRegistryOnce.Do(func() {
		encoderRegistryIns = &encoderRegistry{}
		encoderRegistryIns.setup()
	})
}
