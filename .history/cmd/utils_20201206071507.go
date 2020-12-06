package cmd

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func tools() {
	fmt.Println("hello from tools.")
}

func merge(rootPath string) {
	// outFileName := "d:\\merge_result.txt"
	outFileName := "merge_result.wav"
	outFile, openErr := os.OpenFile(outFileName, os.O_CREATE|os.O_WRONLY, 0600)

	indexFileName := "index_result.txt"
	indexFile, openErr := os.OpenFile(indexFileName, os.O_CREATE|os.O_WRONLY, 0600)

	if openErr != nil {
		fmt.Printf("Can not open file %s or %s!", outFileName, indexFileName)
	}
	bWriter := bufio.NewWriter(outFile)

	index := 1
	offset := 0

	filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		fmt.Println("Processing:", path)
		//filter, only wav file
		if strings.HasSuffix(path, ".wav") {
			// open
			fp, fpOpenErr := os.Open(path)
			if fpOpenErr != nil {
				fmt.Printf("Can not open file %v", fpOpenErr)
				return fpOpenErr
			}

			// merge
			bReader := bufio.NewReader(fp)
			totalCount := 0
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
					totalCount += readCount
				}
			}

			// record index

			offest += totalCount

		}
		return err
	})
	bWriter.Flush()
}
