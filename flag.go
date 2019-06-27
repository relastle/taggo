package main

import (
	"flag"
	"log"
)

var (
	tag       string
	colorStr  string
	delimiter string
	index     int
	nerdIndex int
)

func checkColor() bool {
	ks := []string{}
	for k := range colorFuncMap {
		if colorStr == k {
			return true
		}
		ks = append(ks, k)
	}
	log.Fatalf("color must be any of %v\n", ks)
	return false
}

func parse() {
	flag.StringVar(&tag, "tag", "", "The tag value. It is inserted in head of every line")
	flag.IntVar(&index, "index", -1, "If this is set, `index` column value is use as tag(and colored)")
	flag.StringVar(&delimiter, "delimiter", "\t", "Delimiter")
	flag.StringVar(&colorStr, "color", "", "Color that is applied to tag.")
	flag.IntVar(&nerdIndex, "nerdIndex", -1, "If this is set, `nerdIndex` column is decorated by appropriate nerd font badges")
	flag.Parse()
	// Validation
	checkColor()
}
