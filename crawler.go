package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	var root string
	if len(os.Args) == 1 {
		log.Fatal("No path given, Please specify path.")
		return
	}
	if root = os.Args[1]; root == "" {
		log.Fatal("No path given, Please specify path.")
		return
	}
	// filepath.Walk
	fmt.Println("Scanning path: ", root)
	files, err := FilePathWalkDir(root)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		fmt.Println(file)
	}
	fmt.Println("Total files scanned:", len(files))
}
func FilePathWalkDir(root string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return err

	})
	return files, err
}
