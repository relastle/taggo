package taggo

import (
	"flag"
	"log"
	"strconv"
	"strings"
)

var (
	// --Flag--
	// Tag
	tag            string
	tagColorString string
	tagDelimiter   string
	// Colors
	colorizeQuery string
	// Icons
	iconIndicesString string
	// Common(Colors, Icons)
	// this is used commonly between Colors and Icons
	delimiter string

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
		index, err := strconv.Atoi(elm)
		if err != nil {
			log.Fatal(err)
		}
		res = append(res, index)
	}
	return res
}

// FlagParse parse command line arguments
func FlagParse() {
	flag.StringVar(&tag, "tag", "", "The tag value. It is inserted in head of every line")
	flag.StringVar(&tagColorString, "tagColor", "", "Color that is applied to tag.")
	flag.StringVar(&tagDelimiter, "tagDelimiter", "\t", "Delimiter used to delimite tag")

	flag.StringVar(&colorizeQuery, "colorizeQuery", "", "It requires the comma-seperated query to colorize columns ('0:red,1:blue,2:green')")
	flag.StringVar(&iconIndicesString, "iconIndices", "", "Index list which will be applied icon automatically (0,2,3)")
	flag.StringVar(&delimiter, "delimiter", "\t", "Delimiter that is parse a whole line")
	flag.Parse()

	tagColor = checkColor(tagColorString)
	colorizer = parseColorizeQuery(colorizeQuery)
	iconIndices = parseIndices(iconIndicesString)
}
