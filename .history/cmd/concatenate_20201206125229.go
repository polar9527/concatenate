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
	var base16 string

	flag.StringVar(&dataPath, "dataPath", "data", "data path")
	flag.StringVar(&outImageName, "out", "wavs.img", "wavs.img file")
	flag.StringVar(&imageMapName, "index", "wavs.map", "wavs.map file")
	flag.StringVar(&base16, "base", "0x10000", "base address")

	flag.Parse()
	base10, err := strconv.ParseInt(base16, 16, 64)
	if err != nil {
		fmt.Errorf("base format error: %s", err.Error())
		return
	}

	merge(dataPath, outImageName, imageMapName, base10)
}
