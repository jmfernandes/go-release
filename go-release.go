package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func getFilePath(arg string) string {
	path, err := filepath.Abs(arg)
	check(err)
	removeFileName(&path)
	return path
}

func removeFileName(path *string) {
	fileInfo, err := os.Stat(*path)
	check(err)
	if fileInfo.Mode()&os.ModeType != os.ModeDir {
		*path = filepath.Dir(*path)
	}
}

func checkExists(path string) {
	_, err := os.Stat(path)
	check(err)
}

func main() {
	args := os.Args[1:]
	path := "."
	if length := len(args); length != 0 {
		path = args[0]
	}
	origin := getFilePath(path)
	checkExists(origin)
	fmt.Println(origin)
}
