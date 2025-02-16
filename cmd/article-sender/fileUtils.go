package main

import (
	"fmt"
	"os"
)

func printDirContents(path string) {
	// check for / at the start of path

	f, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	files, err := f.Readdir(0)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, v := range files {
		fmt.Println(v.Name())
	}
}
