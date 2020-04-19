package main

import (
	"github.com/hyc3z/Omaticaya/src/global"
)

func main() {
	global.InitLogger()
	defer global.Logger.Sync()
	global.InitInfo()
}
