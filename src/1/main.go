package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
)

func dirTree(out io.Writer, path string, printFiles bool) error {
	// Create a nested function
	var walk func(string, string) error
	walk = func(path string, prefix string) error {
		fileSystemObjects, err := os.ReadDir(path)
		if err != nil {
			return err
		}

		// Sort by name
		sort.Slice(fileSystemObjects, func(i, j int) bool {
			return fileSystemObjects[i].Name() < fileSystemObjects[j].Name()
		})

		// Filter by object type
		if !printFiles {
			var dirsOnly []os.DirEntry
			for _, obj := range fileSystemObjects {
				if obj.IsDir() {
					dirsOnly = append(dirsOnly, obj)
				}
			}
			fileSystemObjects = dirsOnly
		}

		for index, fileSystemObject := range fileSystemObjects {
			isLast := index == len(fileSystemObjects)-1

			// Ignore .DS_Store MacOS
			if fileSystemObject.Name() == ".DS_Store" {
				continue
			}

			if fileSystemObject.IsDir() {
				printDir(out, fileSystemObject, prefix, isLast)
				subDir := filepath.Join(path, fileSystemObject.Name())
				// Add recursive call
				if isLast {
					if err := walk(subDir, prefix+"\t"); err != nil {
						return err
					}
				} else {
					if err := walk(subDir, prefix+"│\t"); err != nil {
						return err
					}
				}
			} else if printFiles {
				// Print files
				if err := printFile(out, fileSystemObject, prefix, isLast); err != nil {
					return err
				}
			}
		}

		return nil
	}

	return walk(path, "")
}

// Print directories
func printDir(
	out io.Writer,
	fileSystemObject os.DirEntry,
	prefix string, isLast bool,
) {
	if isLast {
		fmt.Fprintf(out, "%s└───%s\n", prefix, fileSystemObject.Name())
	} else {
		fmt.Fprintf(out, "%s├───%s\n", prefix, fileSystemObject.Name())
	}
}

// Print files
func printFile(
	out io.Writer,
	fileSystemObject os.DirEntry,
	prefix string,
	isLast bool,
) error {
	fileInfo, err := fileSystemObject.Info()

	if err != nil {
		return err
	}

	// If file is empty
	if fileInfo.Size() == 0 {
		if isLast {
			fmt.Fprintf(out, "%s└───%s (empty)\n", prefix, fileSystemObject.Name())
		} else {
			fmt.Fprintf(out, "%s├───%s (empty)\n", prefix, fileSystemObject.Name())
		}
		return nil
	}

	// If file is not empty
	if isLast {
		fmt.Fprintf(
			out, "%s└───%s (%db)\n", prefix, fileSystemObject.Name(), fileInfo.Size(),
		)
	} else {
		fmt.Fprintf(
			out, "%s├───%s (%db)\n", prefix, fileSystemObject.Name(), fileInfo.Size(),
		)
	}
	return nil
}

// Main
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
