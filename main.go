package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	arguments := os.Args

	if len(arguments) < 2 {
		fmt.Println("ERROR: No Arguments were given, Please provide at least one argument!")
		return
	} else if len(arguments) > 2 {
		fmt.Println("WARNING: This command doesn't support multiple arguments yet, we will be adding this feature in a future release!")
		return
	}

	path := os.Getenv("PATH")
	pathList := filepath.SplitList(path)
	for _, directory := range pathList {
		fullPath := filepath.Join(directory, arguments[1])
		fileInfo, err := os.Stat(fullPath)
		if err == nil {
			mode := fileInfo.Mode()
			if mode.IsRegular() && mode&0111 != 0 {
				fmt.Println(fullPath)
				return
			}
		}
	}
}
