package main

import (
	"os"
	"path/filepath"
	"strings"
)

func main() {
	var files []string

	root := "."
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})

	if err != nil {
		panic(err)
	}

	for _, file := range files {
		if isImage(file) {
			os.Rename(file, addPrefix(file))
		}
	}
}

func addPrefix(file string) string {
	var prefix = "a_"
	return prefix + file
}

func removePrefix(file string) string {
	var prefix = "a_a-"
	return strings.Replace(file, prefix, "", -1)
}

func isImage(path string) bool {
	if strings.Contains(path, "jpg") || strings.Contains(path, "png") {
		return true
	}

	return false
}
