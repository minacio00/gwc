package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func lineCounter(r io.Reader) int {
	count := 0
	separator := []byte{'\n'}
	buf := make([]byte, 32*1024)

	for {
		c, err := r.Read(buf)
		count += bytes.Count(buf[:c], separator)

		switch {
		case err == io.EOF:
			return count
		case err != nil:
			return -1
		}
	}
}

func wordCounter(r io.Reader) int {
	count := 0
	reader := bufio.NewReader(r)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		lineArr := strings.Fields(string(line))
		count += len(lineArr)
	}
	return count
}

func main() {
	cPtr := flag.Bool("c", false, "counts how many bytes in a file")
	lPtr := flag.Bool("l", false, "number of lines in the file")
	wPtr := flag.Bool("w", false, "counts the number of words in a file")
	filePath := os.Args[len(os.Args)-1]

	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	flag.Parse()

	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		panic(err)
	}

	switch {
	case *cPtr:
		fmt.Printf("%d %s", fileInfo.Size(), filePath)
	case *lPtr:
		fmt.Printf("%d %s", lineCounter(file), filePath)
	case *wPtr:
		fmt.Printf("%d %s", wordCounter(file), filePath)
	}
}
