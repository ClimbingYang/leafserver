package gate

import (
	"leafserver/src/server/game"
	"leafserver/src/server/msg"
)

func init() {
	msg.ProcessorJson.SetRouter(&msg.Hello{}, game.ChanRPC)
}
