package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

const (
	BUFFER_SIZE = 8192
)

func ConvertFile(in, out string) (bool, int, error) {
	// Trying to open input file
	inFile, err := os.Open(in)

	if err != nil {
		return false, -1, err
	}

	// Opening output file
	outFile, err := os.OpenFile(out, os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		return false, -1, err
	}

	// Creating buffer
	buffer := make([]byte, BUFFER_SIZE)

	// Reading input file
	reader := bufio.NewReader(inFile)
	writer := bufio.NewWriter(outFile)
	defer writer.Flush()

	offset := 0
	for {
		// Reading data from input file
		n, err := reader.Read(buffer)

		if err != nil {
			break
		}

		data := buffer[:n]
		Encrypt(&data, offset, n)

		offset += n

		// Writing data to output file
		_, err = writer.Write(data)
		if err != nil {
			fmt.Println("Warning: Unexpected error -> " + err.Error())
			return false, -1, err
		}
	}

	return true, offset, nil
}

func Convert(in string) (bool, int, error) {
	fmt.Printf("Converting %s ...\n", in)
	start := time.Now()
	defer func() {
		fmt.Println("Time elapsed: ", time.Since(start))
	}()
	if strings.HasSuffix(in, ".flac") || strings.HasSuffix(in, ".mp3") {
		fmt.Printf("%s has a suffix of .flac / .mp3, assuming already converted.\n", in)
		return true, 0, nil
	} else if strings.HasSuffix(in, ".qmcflac") {
		return ConvertFile(in, strings.TrimSuffix(in, ".qmcflac")+".flac")
	} else if strings.HasSuffix(in, ".qmc0") {
		return ConvertFile(in, strings.TrimSuffix(in, ".qmc0")+".mp3")
	} else if strings.HasSuffix(in, ".qmc3") {
		return ConvertFile(in, strings.TrimSuffix(in, ".qmc3")+".mp3")
	} else {
		fmt.Println("Unrecognised file suffix in ", in, ", ignored.")
		return true, 0, nil
	}
}
