package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	host, _ := os.Hostname()
	fmt.Println("Host       :", host)

	wd, _ := os.Getwd()
	fmt.Println("Working Dir:", wd)

	fmt.Println("Environment:")
	for _, env := range os.Environ() {
		fmt.Printf("  %s\n", env)
	}

	fmt.Println("Files:")
	files, _ := filepath.Glob("*")
	for _, file := range files {
		fi, err := os.Stat(filepath.Join(".", file))
		if err != nil {
			fmt.Printf("  ? %s: %v", file, err)
			continue
		}
		filetype := "F"
		if fi.IsDir() {
			filetype = "D"
		}

		fmt.Printf("  %s %s\n", filetype, file)
	}
}
