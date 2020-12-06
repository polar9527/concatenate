package main

import (
	"flag"
	"strconv"
)

// var (
// 	var rootPath = "data"
// 	var outImageName = "wavs.img"
// 	var imageMapName = "wavs.map"
// 	var base = 99999
// )

var ()

func init() {

	outImageName
	imageMapName

}

func main() {
	var dataPath string
	var outImageName string
	var imageMapName = string
	var base16 string

	flag.StringVar(&dataPath, "dataPath", "data", "data path")
	flag.StringVar(&outImageName, "out", "wavs.img", "wavs.img file")
	flag.StringVar(&imageMapName, "index", "wavs.map", "wavs.map file")
	flag.StringVar(&base16, "base", "0x0000", "base address")

	flag.Parse()
	base10 := strconv.ParseInt(base16, 16, 64)

	strconv.FormatInt(i, 16)
	merge(rootPath, outImageName, imageMapName, base10)
}
