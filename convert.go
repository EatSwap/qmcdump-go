package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	BUFFER_SIZE = 8192
)

func ConvertFile(in, out string) (bool, error) {
	// Trying to open input file
	inFile, err := os.Open(in)

	if err != nil {
		return false, err
	}

	// Opening output file
	outFile, err := os.OpenFile(out, os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		return false, err
	}

	// Creating buffer
	buffer := make([]byte, BUFFER_SIZE)

	// Reading input file
	reader := bufio.NewReader(inFile)
	writer := bufio.NewWriter(outFile)
	defer writer.Flush()

	for offset := 0; ; offset += BUFFER_SIZE {
		// Reading data from input file
		n, err := reader.Read(buffer)

		if err != nil {
			break
		}

		data := buffer[:n]
		Encrypt(&data, offset, n)

		// Writing data to output file
		_, err = writer.Write(data)
		if err != nil {
			fmt.Println("Warning: Unexpected error -> " + err.Error())
			return false, err
		}
	}

	return true, nil
}
