package main

import (
	"fmt"
	"github.com/name5566/leaf"
	lconf "github.com/name5566/leaf/conf"
	"leafserver/src/server/conf"
	"leafserver/src/server/game"
	"leafserver/src/server/gate"
	"leafserver/src/server/login"
)

func main() {
	lconf.LogLevel = conf.Server.LogLevel
	lconf.LogPath = conf.Server.LogPath
	lconf.LogFlag = conf.LogFlag
	lconf.ConsolePort = conf.Server.ConsolePort
	lconf.ProfilePath = conf.Server.ProfilePath

	fmt.Println("aaaa")

	leaf.Run(
		game.Module,
		gate.Module,
		login.Module,
	)
}
