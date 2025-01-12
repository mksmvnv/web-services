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
	// Read directory
	fileSystemObjects, err := os.ReadDir(path)
	if err != nil {
		return err
	}
	for _, fileSystemObject := range fileSystemObjects {
		// If we use -f argument
		if printFiles {
			if fileSystemObject.IsDir() {
				fmt.Fprintf(out, "%s\n", fileSystemObject.Name())
				subDir := filepath.Join(path, fileSystemObject.Name())
				dirTree(out, subDir, printFiles)
			} else {
				const emptySize int64 = 0
				fileSize, err := os.Stat(filepath.Join(path, fileSystemObject.Name()))
				if err != nil {
					return err
				}
				// If file is empty
				if fileSize.Size() == emptySize {
					fmt.Fprintf(out, "%s (empty)\n", fileSystemObject.Name())
					continue
				}
				fmt.Fprintf(out, "%s (%db)\n", fileSystemObject.Name(), fileSize.Size())
			}
		}
		// If we don't use -f argument
		if !printFiles && fileSystemObject.IsDir() {
			fmt.Fprintf(out, "%s\n", fileSystemObject.Name())
			subDir := filepath.Join(path, fileSystemObject.Name())
			dirTree(out, subDir, printFiles)
		}
	}
	return nil
}
