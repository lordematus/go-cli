package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	// -mode=lines|words|bytes
	mode := flag.String("m", "lines", "Count mode: lines|words|bytes")
	flag.Parse()

	fmt.Println(count(os.Stdin, *mode))
}

func count(r io.Reader, countMode string) int {
	scanner := bufio.NewScanner(r)

	switch countMode {
		case "words":
			scanner.Split(bufio.ScanWords)
		case "bytes":
			scanner.Split(bufio.ScanBytes)
	}

	wc := 0
	for scanner.Scan() {
		wc++
	}

	return wc
}
