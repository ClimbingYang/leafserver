package msg

import (
	"github.com/name5566/leaf/network"
	"github.com/name5566/leaf/network/json"
)

var Processor network.Processor

var ProcessorJson json.Processor

func init() {
	ProcessorJson = *json.NewProcessor()
	ProcessorJson.Register(&Hello{})
}

type Hello struct {
	Name string
}
