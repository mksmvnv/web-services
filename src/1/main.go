package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}

func dirTree(out io.Writer, path string, printFiles bool) error {
	items, err := os.ReadDir(path)
	if err != nil {
		return err
	}
	for _, item := range items {
		if printFiles {
			if item.IsDir() {
				fmt.Fprintln(out, item.Name())
				subDir := filepath.Join(path+"/", item.Name())
				dirTree(out, subDir, printFiles)
			} else {
				fileSize, err := os.Stat(item.Name())
				if err != nil {
					return err
				}
				fmt.Fprintf(out, "%s (%db)\n", item.Name(), fileSize.Size())
			}

		}
		if !printFiles && item.IsDir() {
			fmt.Fprintln(out, item.Name())
			subDir := filepath.Join(path+"/", item.Name())
			dirTree(out, subDir, printFiles)
		}
	}
	return nil
}
