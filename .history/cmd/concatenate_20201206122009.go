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

var (
	var rootPath string
	var outImageName string
	var imageMapName string
	var base int64
	var cliName = flag.String("dataPath", "f", "data", "data path")
)

func init(){
	rootPath = flag.String("", )
	outImageName
	imageMapName
	
}

func main() {

	merge(rootPath, outImageName, imageMapName, base)
}