package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	//if len(args) < 2 {
	//	fmt.Println("Usage: mvold <filename-pattern> <directory> <age>")
	//	return
	//}
	filePattern := "/Users/tcrone/temp/*" //args[1]
	files, _ := filepath.Glob(filePattern)

	for i, filename := range files {
		file, _ := os.Open(filename)
		fileInfo, _ := file.Stat()

		if fileInfo.IsDir() {
			fmt.Println("Dir", i, filename)
		} else {
			fmt.Println("File", i, filename)
		}
	}
}
