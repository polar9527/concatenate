package main

import (
	"flag"
	"fmt"
	"strconv"
)

func main() {
	var dataPath string
	var outImageName string
	var imageMapName string
	var base16Str string
	var hexFlag bool

	flag.StringVar(&dataPath, "dataPath", "data", "data path")
	flag.StringVar(&outImageName, "out", "wavs.img", "wavs.img file")
	flag.StringVar(&imageMapName, "index", "wavs.map", "wavs.map file")
	flag.StringVar(&base16Str, "base", "0x800000", "base address")
	flag.BoolVar(&hexFlag, "hex", true, "whether use hex offset")

	flag.Parse()
	base10, err := strconv.ParseInt(base16Str, 0, 64)
	if err != nil {
		fmt.Errorf("base format error: %s", err.Error())
		return
	}

	merge(dataPath, outImageName, imageMapName, base10, hexFlag)
}
