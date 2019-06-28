package main

import (
	"fmt"
	"path"
	"strings"
)

func addBadge(target string) string {
	base := path.Base(target)
	dotIndex := strings.LastIndex(base, ".")
	if dotIndex < 0 {
		return "  " + target
	}
	extension := base[dotIndex+1 : len(base)]
	if badge, ok := nerdFontMap[extension]; ok {
		return fmt.Sprintf("%v:%v", badge, target)
	}
	return fmt.Sprintf("%v:%v", nerdFontMap["txt"], target)
}
