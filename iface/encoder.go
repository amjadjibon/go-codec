package iface

type Encoder interface {
	Name() string
	Version() string
	Encode(src []byte) ([]byte, error)
	Decode(src []byte) ([]byte, error)
}

