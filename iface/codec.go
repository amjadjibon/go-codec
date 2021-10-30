package iface

type Codec interface {
	Name() string
	Version() string
	Marshal(v interface{}) ([]byte, error)
	Unmarshal(b []byte, v interface{}) error
}
