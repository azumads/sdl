package main

import (
	"github.com/azumads/sdl"
)

func main() {
	sdl.Config.UserName = "SDK_TEST"
	sdl.Config.PassWord = "p@ssw0rd"
	sdl.GetProjectGroup(1094)
	sdl.GetProject(1120)
	sdl.GetProject(1121)
	sdl.GetTask(1125)
	sdl.GetTask(1126)
}
