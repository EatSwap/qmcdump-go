package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

const (
	ProgramName = "qmcdump-go"
	Version     = "0.1.0"
)

var (
	PrintChannel = make(chan string)
)

func main() {
	Welcome()
	go PrintSomething()
	args := os.Args
	if len(args) == 1 {
		PrintChannel <- fmt.Sprintf("Usage: %s <FILE/FOLDER>\n", args[0])
		os.Exit(1)
	}
	stat, err := os.Stat(args[1])
	if err != nil {
		PrintChannel <- fmt.Sprint(err)
		os.Exit(1)
	}

	start := time.Now()
	dataAmount := 0

	if stat.IsDir() {
		filepath.Walk(args[1], func(path string, info os.FileInfo, err error) error {
			if err != nil || info.IsDir() {
				PrintChannel <- fmt.Sprint(err)
				return err
			}
			ok, d, err := Convert(path)
			if err != nil {
				PrintChannel <- fmt.Sprintf("Conversion failed. Message:")
				PrintChannel <- fmt.Sprint(err)
			} else if !ok {
				PrintChannel <- fmt.Sprintf("Conversion failed with unknown reason.")
			} else {
				dataAmount += d
				PrintChannel <- fmt.Sprintf("Conversion completed successfully.")
			}
			PrintChannel <- fmt.Sprintf("")
			return nil
		})
	} else {
		ok, d, err := Convert(args[1])
		if err != nil {
			PrintChannel <- fmt.Sprintf("Conversion failed. Message:")
			PrintChannel <- fmt.Sprint(err)
		} else if !ok {
			PrintChannel <- fmt.Sprintf("Conversion failed with unknown reason.")
		} else {
			dataAmount += d
			PrintChannel <- fmt.Sprintf("Conversion completed successfully.")
		}
	}

	PrintChannel <- fmt.Sprintf("")
	PrintChannel <- fmt.Sprintf("Conversion completed in %s.", time.Since(start))
	PrintChannel <- fmt.Sprintf("Total data amount: %d bytes.", dataAmount)
	PrintChannel <- fmt.Sprintf("Conversion Rate: %.3f MB/s.", float64(dataAmount)/time.Since(start).Seconds()/1024/1024)
	PrintChannel <- fmt.Sprintf("Press <Enter> to exit...")
	fmt.Scanln()
}

func PrintSomething() {
	for {
		s, ok := <-PrintChannel
		if !ok {
			break
		}
		fmt.Println(s)
	}
}

func Welcome() {
	fmt.Printf("%s Version [%s]\n", ProgramName, Version)
	fmt.Printf("System information: %s\n", fmt.Sprintf("OS [ %s ], Architecture [ %s ], go runtime Version [ %s ]", runtime.GOOS, runtime.GOARCH, runtime.Version()))
}
