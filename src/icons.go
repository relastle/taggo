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
	"ai":           {code: "\ue7b4", color: "yello"},
	"android":      {code: "\ue70e", color: ""},
	"apple":        {code: "\uf179", color: "magenda"},
	"audio":        {code: "\uf001", color: ""},
	"avro":         {code: "\ue60b", color: ""},
	"bash":         {code: "\uf489", color: ""},
	"c":            {code: "\ue61e", color: ""},
	"clj":          {code: "\ue768", color: ""},
	"coffee":       {code: "\uf0f4", color: ""},
	"conf":         {code: "\ue615", color: ""},
	"cpp":          {code: "\ue61d", color: ""},
	"css":          {code: "\ue749", color: ""},
	"d":            {code: "\ue7af", color: ""},
	"dart":         {code: "\ue798", color: ""},
	"db":           {code: "\uf1c0", color: ""},
	"diff":         {code: "\uf440", color: ""},
	"doc":          {code: "\uf1c2", color: ""},
	"docker":       {code: "\uf308", color: ""},
	"ebook":        {code: "\ue28b", color: ""},
	"env":          {code: "\uf462", color: ""},
	"epub":         {code: "\ue28a", color: ""},
	"erl":          {code: "\ue7b1", color: ""},
	"file":         {code: "\uf15b", color: ""},
	"fish":         {code: "\uf489", color: ""},
	"font":         {code: "\uf031", color: ""},
	"gform":        {code: "\uf298", color: ""},
	"git":          {code: "\uf1d3", color: ""},
	"go":           {code: "\ue626", color: "cyan"},
	"gruntfile.js": {code: "\ue74c", color: ""},
	"hs":           {code: "\ue777", color: ""},
	"html":         {code: "\uf13b", color: ""},
	"image":        {code: "\uf1c5", color: ""},
	"iml":          {code: "\ue7b5", color: ""},
	"java":         {code: "\ue204", color: ""},
	"js":           {code: "\ue74e", color: "yellow"},
	"json":         {code: "\ue60b", color: "yellow"},
	"jsx":          {code: "\ue7ba", color: "yellow"},
	"less":         {code: "\ue758", color: ""},
	"log":          {code: "\uf18d", color: ""},
	"lua":          {code: "\ue620", color: "yellow"},
	"md":           {code: "\uf48a", color: ""},
	"mustache":     {code: "\ue60f", color: ""},
	"npmignore":    {code: "\ue71e", color: ""},
	"pdf":          {code: "\uf1c1", color: ""},
	"php":          {code: "\ue73d", color: "red"},
	"pl":           {code: "\ue769", color: ""},
	"ppt":          {code: "\uf1c4", color: ""},
	"psd":          {code: "\ue7b8", color: ""},
	"py":           {code: "\ue606", color: "yellow"},
	"r":            {code: "\uf25d", color: ""},
	"rb":           {code: "\ue21e", color: "red"},
	"rdb":          {code: "\ue76d", color: ""},
	"rss":          {code: "\uf09e", color: ""},
	"rubydoc":      {code: "\ue73b", color: "red"},
	"sass":         {code: "\ue603", color: ""},
	"scala":        {code: "\ue737", color: ""},
	"sh":           {code: "\uf489", color: ""},
	"sqlite3":      {code: "\ue7c4", color: "blue"},
	"styl":         {code: "\ue600", color: ""},
	"tex":          {code: "\ue600", color: ""},
	"ts":           {code: "\ue628", color: "yellow"},
	"twig":         {code: "\ue61c", color: ""},
	"txt":          {code: "\uf15c", color: ""},
	"video":        {code: "\uf03d", color: ""},
	"vim":          {code: "\ue62b", color: "blue"},
	"windows":      {code: "\uf17a", color: ""},
	"xls":          {code: "\uf1c3", color: ""},
	"xml":          {code: "\ue619", color: ""},
	"yarn.lock":    {code: "\ue718", color: ""},
	"yml":          {code: "\uf481", color: ""},
	"zip":          {code: "\uf410", color: ""},
	"zsh":          {code: "\uf489", color: ""},
}

var basenameIconMap = map[string]Icon{
	"Dockerfile": {code: "\ue7b0", color: "cyan"},
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
	extension := base[dotIndex+1:]
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
	return target[delimIndex+len(iconDelimiter):]
}
