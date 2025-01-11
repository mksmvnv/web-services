package main

import (
	"fmt"
	"io"
	"os"
)

type Item struct {
	Name string
}

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
	var buf []Item
	items, err := os.ReadDir(path)
	if err != nil {
		return err
	}
	for _, item := range items {
		if printFiles {
			buf = append(buf, Item{item.Name()})
		} else {
			if item.IsDir() {
				buf = append(buf, Item{item.Name()})
			}
		}
	}
	fmt.Println(buf)
	return nil
}
