package main

var (
	var rootPath = "data"
	var outImageName = "wavs.img"
	var imageMapName = "wavs.map"
	var base = 99999
)

func main() {

	merge(rootPath, outImageName, imageMapName, base)
}
