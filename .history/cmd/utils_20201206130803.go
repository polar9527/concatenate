package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func merge(rootPath, outImageName, imageMapName string, base int64) {
	outFileName := outImageName
	outFile, openErr := os.OpenFile(outFileName, os.O_CREATE|os.O_WRONLY, 0600)

	indexFileName := imageMapName
	indexFile, openErr := os.OpenFile(indexFileName, os.O_CREATE|os.O_WRONLY, 0600)

	if openErr != nil {
		fmt.Printf("Can not open file %s or %s!", outFileName, indexFileName)
	}
	bWriter := bufio.NewWriter(outFile)
	iWriter := bufio.NewWriter(indexFile)
	index := 1
	offset := int64(0)

	filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		fmt.Println("Processing:", path)
		// filter, only wav file
		if strings.HasSuffix(path, ".wav") {
			// open
			fp, fpOpenErr := os.Open(path)
			if fpOpenErr != nil {
				fmt.Printf("Can not open file %v", fpOpenErr)
				return fpOpenErr
			}

			// merge
			bReader := bufio.NewReader(fp)
			totalCount := int64(0)
			for {
				buffer := make([]byte, 1024)
				readCount, readErr := bReader.Read(buffer)
				if readErr == io.EOF {
					if readCount != 0 {
						fmt.Printf("readCount %v is not equal zero!", readCount)
					}
					break
				} else {
					bWriter.Write(buffer[:readCount])
					totalCount += int64(readCount)
				}
			}

			// record index

			// i := fmt.Sprintf("%v,%v,%v\n", index, fp.Name(),strconv.FormatInt(base + offset, 16))
			// i := fmt.Sprintf("%v,%v,%v\n", index, fp.Name(), base+offset)
			_, err := iWriter.WriteString(i)
			if err != nil {
				outFile.Close()
				indexFile.Close()
				os.Remove(outFileName)
				os.Remove(indexFileName)
				panic("write index error")
			}
			index++
			// offset += info.Size()
			offset += totalCount
		}
		return err
	})
	bWriter.Flush()
	iWriter.Flush()
}
