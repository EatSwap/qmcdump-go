package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Printf("Usage: %s <FILE/FOLDER>\n", args[0])
		os.Exit(1)
	}
	stat, err := os.Stat(args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if stat.IsDir() {
		filepath.Walk(args[1], func(path string, info os.FileInfo, err error) error {
			if err != nil || info.IsDir() {
				fmt.Println(err)
				return err
			}
			ok, _, err := Convert(path)
			if err != nil {
				fmt.Println("Conversion failed. Message:")
				fmt.Println(err)
			} else if !ok {
				fmt.Println("Conversion failed with unknown reason.")
			} else {
				fmt.Println("Conversion completed successfully.")
			}
			fmt.Println("")
			return nil
		})
	} else {
		ok, _, err := Convert(args[1])
		if err != nil {
			fmt.Println("Conversion failed. Message:")
			fmt.Println(err)
		} else if !ok {
			fmt.Println("Conversion failed with unknown reason.")
		} else {
			fmt.Println("Conversion completed successfully.")
		}
	}
	fmt.Println("Press <Enter> to exit...")
	fmt.Scanln()
}
