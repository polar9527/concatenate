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
)

func init(){
	rootPath = flag.StringVar()
	outImageName
	imageMapName
	
}

func main() {

	merge(rootPath, outImageName, imageMapName, base)
}
