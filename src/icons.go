package taggo

import (
	"fmt"
	"path"
	"strings"
)

// Icon represents one icon
// which has nerdfont unicode
// and its color
type Icon struct {
	code  string
	color Color
}

// The code below is highly inspired by
// https://github.com/athityakumar/colorls/blob/master/lib/yaml/files.yaml
// Jun 27, 2019

// extensionNerdFontMap defines the `extension` -> `nerdfont`
// mapping.
var extensionIconMap = map[string]Icon{
	"ai":           Icon{code: "\ue7b4", color: "yello"},
	"android":      Icon{code: "\ue70e", color: ""},
	"apple":        Icon{code: "\uf179", color: "magenda"},
	"audio":        Icon{code: "\uf001", color: ""},
	"avro":         Icon{code: "\ue60b", color: ""},
	"bash":         Icon{code: "\uf489", color: ""},
	"c":            Icon{code: "\ue61e", color: ""},
	"clj":          Icon{code: "\ue768", color: ""},
	"coffee":       Icon{code: "\uf0f4", color: ""},
	"conf":         Icon{code: "\ue615", color: ""},
	"cpp":          Icon{code: "\ue61d", color: ""},
	"css":          Icon{code: "\ue749", color: ""},
	"d":            Icon{code: "\ue7af", color: ""},
	"dart":         Icon{code: "\ue798", color: ""},
	"db":           Icon{code: "\uf1c0", color: ""},
	"diff":         Icon{code: "\uf440", color: ""},
	"doc":          Icon{code: "\uf1c2", color: ""},
	"docker":       Icon{code: "\uf308", color: ""},
	"ebook":        Icon{code: "\ue28b", color: ""},
	"env":          Icon{code: "\uf462", color: ""},
	"epub":         Icon{code: "\ue28a", color: ""},
	"erl":          Icon{code: "\ue7b1", color: ""},
	"file":         Icon{code: "\uf15b", color: ""},
	"fish":         Icon{code: "\uf489", color: ""},
	"font":         Icon{code: "\uf031", color: ""},
	"gform":        Icon{code: "\uf298", color: ""},
	"git":          Icon{code: "\uf1d3", color: ""},
	"go":           Icon{code: "\ue626", color: "cyan"},
	"gruntfile.js": Icon{code: "\ue74c", color: ""},
	"hs":           Icon{code: "\ue777", color: ""},
	"html":         Icon{code: "\uf13b", color: ""},
	"image":        Icon{code: "\uf1c5", color: ""},
	"iml":          Icon{code: "\ue7b5", color: ""},
	"java":         Icon{code: "\ue204", color: ""},
	"js":           Icon{code: "\ue74e", color: "yellow"},
	"json":         Icon{code: "\ue60b", color: "yellow"},
	"jsx":          Icon{code: "\ue7ba", color: "yellow"},
	"less":         Icon{code: "\ue758", color: ""},
	"log":          Icon{code: "\uf18d", color: ""},
	"lua":          Icon{code: "\ue620", color: "yellow"},
	"md":           Icon{code: "\uf48a", color: ""},
	"mustache":     Icon{code: "\ue60f", color: ""},
	"npmignore":    Icon{code: "\ue71e", color: ""},
	"pdf":          Icon{code: "\uf1c1", color: ""},
	"php":          Icon{code: "\ue73d", color: "red"},
	"pl":           Icon{code: "\ue769", color: ""},
	"ppt":          Icon{code: "\uf1c4", color: ""},
	"psd":          Icon{code: "\ue7b8", color: ""},
	"py":           Icon{code: "\ue606", color: "yellow"},
	"r":            Icon{code: "\uf25d", color: ""},
	"rb":           Icon{code: "\ue21e", color: "red"},
	"rdb":          Icon{code: "\ue76d", color: ""},
	"rss":          Icon{code: "\uf09e", color: ""},
	"rubydoc":      Icon{code: "\ue73b", color: "red"},
	"sass":         Icon{code: "\ue603", color: ""},
	"scala":        Icon{code: "\ue737", color: ""},
	"sh":           Icon{code: "\uf489", color: ""},
	"sqlite3":      Icon{code: "\ue7c4", color: "blue"},
	"styl":         Icon{code: "\ue600", color: ""},
	"tex":          Icon{code: "\ue600", color: ""},
	"ts":           Icon{code: "\ue628", color: "yellow"},
	"twig":         Icon{code: "\ue61c", color: ""},
	"txt":          Icon{code: "\uf15c", color: ""},
	"video":        Icon{code: "\uf03d", color: ""},
	"vim":          Icon{code: "\ue62b", color: "blue"},
	"windows":      Icon{code: "\uf17a", color: ""},
	"xls":          Icon{code: "\uf1c3", color: ""},
	"xml":          Icon{code: "\ue619", color: ""},
	"yarn.lock":    Icon{code: "\ue718", color: ""},
	"yml":          Icon{code: "\uf481", color: ""},
	"zip":          Icon{code: "\uf410", color: ""},
	"zsh":          Icon{code: "\uf489", color: ""},
}

var basenameIconMap = map[string]Icon{
	"Dockerfile": Icon{code: "\ue7b0", color: "cyan"},
}

func getIcon(target string) Icon {
	base := path.Base(target)
	// if base name mathes to something, return that.
	if icon, ok := basenameIconMap[base]; ok {
		return icon
	}

	// No extension is assumed as simple text file
	dotIndex := strings.LastIndex(base, ".")
	if dotIndex < 0 {
		return extensionIconMap["txt"]
	}
	extension := base[dotIndex+1 : len(base)]
	if icon, ok := extensionIconMap[extension]; ok {
		return icon
	}
	return extensionIconMap["txt"]
}

func addIcon(target string) string {
	icon := getIcon(target)
	return fmt.Sprintf(
		"%s%s%s",
		ColorFuncMap[icon.color](icon.code),
		iconDelimiter,
		target,
	)
}

func removeIcon(target string) string {
	delimIndex := strings.Index(target, iconDelimiter)
	return target[delimIndex+len(iconDelimiter) : len(target)]
}
