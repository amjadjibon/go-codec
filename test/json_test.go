package test

import (
	"testing"

	"github.com/amjadjibon/go-codec/constant"
	"github.com/amjadjibon/go-codec/utility"
)

func TestCodecJSON(t *testing.T) {
	codec, err := utility.GetCodec(constant.JSON)
	if err != nil {
		t.Error(err)
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
