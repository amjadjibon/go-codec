package main

import (
	_ "embed"
	"fmt"
	"github.com/amjadjibon/go-codec/constant"
	"github.com/amjadjibon/go-codec/utility"
)

//go:embed example.json
var jsonBytes []byte

type ExampleJSON struct {
	Id int
	Name string
}

func (e *ExampleJSON) Display() string {
	return fmt.Sprintf("ExampleJSON {id: %d, name: %s}", e.Id, e.Name)
}

func main() {
	codec := utility.GetCodec(constant.JSON)

	fmt.Printf("name: %s\n", codec.Name())
	fmt.Printf("version: %s\n", codec.Version())

	var ex = new(ExampleJSON)

	err := codec.Unmarshal(jsonBytes, ex)
	if err != nil {
		return
	}

	fmt.Println(ex.Display())

	marshal, err := codec.Marshal(ex)
	if err != nil {
		return
	}
	fmt.Println(string(marshal))
}
