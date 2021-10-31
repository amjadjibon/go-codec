package registry

import (
	"sync"

	"github.com/amjadjibon/go-codec/iface"
)

var (
	codecRegistryOnce sync.Once
	codecRegistryIns  *codecRegistry
)

type codecRegistry struct {
	registry map[string]iface.Codec
}

func (g *codecRegistry) setup() {
	g.registry = make(map[string]iface.Codec)
}

func (g *codecRegistry) AddCodec(name string, codec iface.Codec) {
	g.registry[name] = codec
}

func (g *codecRegistry) GetCodec(name string) iface.Codec {
	return g.registry[name]
}

func CodecRegistry() *codecRegistry {
	return codecRegistryIns
}

func init() {
	codecRegistryOnce.Do(func() {
		codecRegistryIns = &codecRegistry{}
		codecRegistryIns.setup()
	})
}
