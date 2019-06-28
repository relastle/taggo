package taggo

import (
	fatih_color "github.com/fatih/color"
)

func asItIs(format string, a ...interface{}) string {
	return format
}

// ColorFuncMap TODO
var ColorFuncMap = map[Color](func(format string, a ...interface{}) string){
	"":        asItIs,
	"black":   fatih_color.BlackString,
	"red":     fatih_color.RedString,
	"green":   fatih_color.GreenString,
	"yellow":  fatih_color.YellowString,
	"blue":    fatih_color.BlueString,
	"magenda": fatih_color.MagentaString,
	"cyan":    fatih_color.CyanString,
	"white":   fatih_color.WhiteString,
}
