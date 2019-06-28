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
	"android":      Icon{code: "\ue70e", color: "black"},
	"apple":        Icon{code: "\uf179", color: "magenda"},
	"audio":        Icon{code: "\uf001", color: "black"},
	"avro":         Icon{code: "\ue60b", color: "black"},
	"bash":         Icon{code: "\uf489", color: "black"},
	"c":            Icon{code: "\ue61e", color: "black"},
	"clj":          Icon{code: "\ue768", color: "black"},
	"coffee":       Icon{code: "\uf0f4", color: "black"},
	"conf":         Icon{code: "\ue615", color: "black"},
	"cpp":          Icon{code: "\ue61d", color: "black"},
	"css":          Icon{code: "\ue749", color: "black"},
	"d":            Icon{code: "\ue7af", color: "black"},
	"dart":         Icon{code: "\ue798", color: "black"},
	"db":           Icon{code: "\uf1c0", color: "black"},
	"diff":         Icon{code: "\uf440", color: "black"},
	"doc":          Icon{code: "\uf1c2", color: "black"},
	"docker":       Icon{code: "\uf308", color: "black"},
	"ebook":        Icon{code: "\ue28b", color: "black"},
	"env":          Icon{code: "\uf462", color: "black"},
	"epub":         Icon{code: "\ue28a", color: "black"},
	"erl":          Icon{code: "\ue7b1", color: "black"},
	"file":         Icon{code: "\uf15b", color: "black"},
	"fish":         Icon{code: "\uf489", color: "black"},
	"font":         Icon{code: "\uf031", color: "black"},
	"gform":        Icon{code: "\uf298", color: "black"},
	"git":          Icon{code: "\uf1d3", color: "black"},
	"go":           Icon{code: "\ue626", color: "cyan"},
	"gruntfile.js": Icon{code: "\ue74c", color: "black"},
	"hs":           Icon{code: "\ue777", color: "black"},
	"html":         Icon{code: "\uf13b", color: "black"},
	"image":        Icon{code: "\uf1c5", color: "black"},
	"iml":          Icon{code: "\ue7b5", color: "black"},
	"java":         Icon{code: "\ue204", color: "black"},
	"js":           Icon{code: "\ue74e", color: "yellow"},
	"json":         Icon{code: "\ue60b", color: "yellow"},
	"jsx":          Icon{code: "\ue7ba", color: "yellow"},
	"less":         Icon{code: "\ue758", color: "black"},
	"log":          Icon{code: "\uf18d", color: "black"},
	"lua":          Icon{code: "\ue620", color: "yellow"},
	"md":           Icon{code: "\uf48a", color: "black"},
	"mustache":     Icon{code: "\ue60f", color: "black"},
	"npmignore":    Icon{code: "\ue71e", color: "black"},
	"pdf":          Icon{code: "\uf1c1", color: "black"},
	"php":          Icon{code: "\ue73d", color: "red"},
	"pl":           Icon{code: "\ue769", color: "black"},
	"ppt":          Icon{code: "\uf1c4", color: "black"},
	"psd":          Icon{code: "\ue7b8", color: "black"},
	"py":           Icon{code: "\ue606", color: "yellow"},
	"r":            Icon{code: "\uf25d", color: "black"},
	"rb":           Icon{code: "\ue21e", color: "red"},
	"rdb":          Icon{code: "\ue76d", color: "black"},
	"rss":          Icon{code: "\uf09e", color: "black"},
	"rubydoc":      Icon{code: "\ue73b", color: "red"},
	"sass":         Icon{code: "\ue603", color: "black"},
	"scala":        Icon{code: "\ue737", color: "black"},
	"sh":           Icon{code: "\uf489", color: "black"},
	"sqlite3":      Icon{code: "\ue7c4", color: "blue"},
	"styl":         Icon{code: "\ue600", color: "black"},
	"tex":          Icon{code: "\ue600", color: "black"},
	"ts":           Icon{code: "\ue628", color: "yellow"},
	"twig":         Icon{code: "\ue61c", color: "black"},
	"txt":          Icon{code: "\uf15c", color: "black"},
	"video":        Icon{code: "\uf03d", color: "black"},
	"vim":          Icon{code: "\ue62b", color: "blue"},
	"windows":      Icon{code: "\uf17a", color: "black"},
	"xls":          Icon{code: "\uf1c3", color: "black"},
	"xml":          Icon{code: "\ue619", color: "black"},
	"yarn.lock":    Icon{code: "\ue718", color: "black"},
	"yml":          Icon{code: "\uf481", color: "black"},
	"zip":          Icon{code: "\uf410", color: "black"},
	"zsh":          Icon{code: "\uf489", color: "black"},
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
	return fmt.Sprintf(ColorFuncMap[icon.color](icon.code))
}
