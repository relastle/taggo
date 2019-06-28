package main

import (
	"github.com/fatih/color"
	taggo "github.com/relastle/taggo/src"
)

func main() {
	taggo.FlagParse()
	color.NoColor = false
	taggo.MainStream()
}
