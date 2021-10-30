package test

import (
	"github.com/amjadjibon/go-codec/constant"
	"github.com/amjadjibon/go-codec/utility"
	"testing"
)

func TestCodecJSON(t *testing.T) {
	codec := utility.GetCodec(constant.JSON)

	if codec == nil {
		t.Error("codec should not be nil")
	}

	type example struct {
		Id   int
		Name string
	}

	var e1 = &example{
		Id:   1,
		Name: "Jon Doe",
	}

	m, err := codec.Marshal(e1)
	if err != nil {
		t.Error(err)
	}

	var e2 = new(example)

	err = codec.Unmarshal(m, e2)
	if err != nil {
		t.Error(err)
	}
}
