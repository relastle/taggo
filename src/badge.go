package taggo

import (
	"fmt"
	"path"
	"strings"
)

func addBadge(target string) string {
	base := path.Base(target)
	// if base name mathes to something, return that.
	if badge, ok := basenameNerdFontMap[base]; ok {
		return fmt.Sprintf("%v:%v", badge, target)
	}

	// No extension is assumed as simple text file
	dotIndex := strings.LastIndex(base, ".")
	if dotIndex < 0 {
		return fmt.Sprintf("%v:%v", extensionNerdFontMap["txt"], target)
	}
	extension := base[dotIndex+1 : len(base)]
	if badge, ok := extensionNerdFontMap[extension]; ok {
		return fmt.Sprintf("%v:%v", badge, target)
	}
	return fmt.Sprintf("%v:%v", extensionNerdFontMap["txt"], target)
}
