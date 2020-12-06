package main

import (
	"flag"
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
	var base int64
	flag.StringVar(&dataPath, "dataPath", "data", "data path")
	flag.StringVar(&outImageName, "out", "wavs.img", "wavs.img file")
	flag.StringVar(&imageMapName, "index", "wavs.map", "wavs.map file")

	flag.Parse()

	merge(rootPath, outImageName, imageMapName, base)
}
