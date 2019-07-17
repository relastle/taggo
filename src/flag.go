package taggo

import (
	"log"
	"strconv"
	"strings"

	"github.com/integrii/flaggy"
)

var (
	// --Flag--
	// Tag
	tag            = ""
	tagColorString = ""
	tagDelimiter   = "\t"
	// Colors
	colorizeQuery = ""
	// Icons
	iconIndicesString = ""
	iconDelimiter     = ":"
	// base-named
	basenamedIndex       = -1
	basenamedDelimiter   = "|||"
	basenamedMaxLen      = 40
	basenamedLenInterval = 10
	// Common(Colors, Icons)
	// this is used commonly between Colors and Icons
	delimiter = "\t"
	// revertFlag contains boolean whether `taggo` is launched
	// to stream input lines, or to revert one line.
	revertFlag = false

	// --Global Vairiable--
	tagColor    Color
	colorizer   Colorizer
	iconIndices []int
)

// Color representing one color
type Color string

// Colorizer contains information for which column should be colorized
// with which color
type Colorizer map[int]Color

func checkColor(s string) Color {
	ks := []string{}
	for k := range ColorFuncMap {
		if s == string(k) {
			return Color(s)
		}
		ks = append(ks, string(k))
	}
	log.Fatalf("color must be any of %v\n", ks)
	return ""
}

func parseColorizeQuery(query string) Colorizer {
	res := Colorizer{}
	pairs := strings.Split(query, ",")
	for _, pair := range pairs {
		if pair == "" {
			continue
		}
		elms := strings.Split(pair, ":")
		if len(elms) != 2 {
			log.Fatal("query must be like '0:red,1:blue'")
		}
		indexString := elms[0]
		colorString := elms[1]
		index, err := strconv.Atoi(indexString)
		if err != nil {
			log.Fatal(err)
		}
		color := checkColor(colorString)
		res[index] = color
	}
	return res
}

func parseIndices(indicesString string) []int {
	res := []int{}
	elms := strings.Split(indicesString, ",")
	for _, elm := range elms {
		if elm == "" {
			continue
		}
		index, err := strconv.Atoi(elm)
		if err != nil {
			log.Fatal(err)
		}
		res = append(res, index)
	}
	return res
}

func checkDelimiters(delimiter string, iconDelimiter string) {
	if strings.Contains(iconDelimiter, delimiter) {
		log.Fatal("delimiter must not be substring of iconDelimiter")
	}
}

func parseDelimiter(delimiter string) string {
	// Special handling of \t
	delimiter = strings.Replace(delimiter, "\\t", "\t", -1)
	return delimiter
}

// FlagParse parse command line arguments
func FlagParse() {
	// Add a flag
	flaggy.String(&tag, "t", "tag", "The tag value. It is inserted in head of every line.")
	flaggy.String(&tagColorString, "c", "tag-color", "Color that is applied to tag.")
	flaggy.String(&tagDelimiter, "s", "tag-delimiter", "Delimiter used to delimite tag.")

	flaggy.String(&colorizeQuery, "q", "colorize-query", "It requires the comma-seperated query to colorize columns ('0:red,1:blue,2:green').")
	flaggy.String(&iconIndicesString, "i", "icon-indices", "Index list which will be applied icon automatically (0,2,3).")
	flaggy.String(&iconDelimiter, "p", "icon-delimiter", "Delimiter that follows icon(it can not be a substring of delimiter if you want to revert correctly).")
	flaggy.Int(&basenamedIndex, "b", "basenamed-index", "Index of filepath whose basename is inserted after tag.")
	flaggy.String(&basenamedDelimiter, "g", "basenamed-delimiter", "Delimiter used as suffix of basename.")
	flaggy.Int(&basenamedMaxLen, "", "basenamed-max-len", "Maximum length of basename that will be basis of aligning.")
	flaggy.Int(&basenamedLenInterval, "", "basenamed-len-interval", "Interval of length for those cannot be aligned by using basenamedMaxLen.")
	flaggy.String(&delimiter, "d", "delimiter", "Delimiter that is parse a whole line.")

	flaggy.Bool(&revertFlag, "r", "revert", "If specified, it revert decorated line(ANSI colors are assumbed to be removed).")
	flaggy.Parse()

	tagDelimiter = parseDelimiter(tagDelimiter)
	iconDelimiter = parseDelimiter(iconDelimiter)
	delimiter = parseDelimiter(delimiter)

	checkDelimiters(delimiter, iconDelimiter)
	tagColor = checkColor(tagColorString)
	colorizer = parseColorizeQuery(colorizeQuery)
	iconIndices = parseIndices(iconIndicesString)
}
