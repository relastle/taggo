package main

import (
	"github.com/fatih/color"
)

func asItIs(format string, a ...interface{}) string {
	return format
}

var colorFuncMap = map[string](func(format string, a ...interface{}) string){
	"":        asItIs,
	"black":   color.BlackString,
	"red":     color.RedString,
	"green":   color.BlueString,
	"yellow":  color.YellowString,
	"blue":    color.CyanString,
	"magenda": color.MagentaString,
	"cyan":    color.CyanString,
	"white":   color.WhiteString,
}

func main() {
	parse()
	color.NoColor = false
	mainStream()
}
