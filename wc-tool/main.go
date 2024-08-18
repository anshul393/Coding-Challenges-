package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

/*
Requirements
 flags: Output
	-c: number of bytes
	-l: number of lines
	-w: number of words
	-m: number of chars

If no file specified, then read from stdin
*/

func main() {
	bytes := flag.Bool("c", false, "number of bytes counted")
	lines := flag.Bool("l", false, "number of lines counted")
	words := flag.Bool("w", false, "number of words counted")
	chars := flag.Bool("m", false, "number of chars counted")

	flag.Parse()

	var splitFunc bufio.SplitFunc
	var r io.Reader
	var prefix string

	switch {
	case len(flag.Args()) == 0:
		r = os.Stdin
	case len(flag.Args()) >= 1:
		// read from the specified file
		filePath := flag.Args()[0]
		f, err := os.Open(filePath)

		if err != nil {
			log.Print("provided filepath is wrong: ", err)
			os.Exit(1)
		}

		r = f
	}

	switch {
	case *bytes:
		splitFunc = bufio.ScanBytes
		prefix = "bytes: "
	case *lines:
		splitFunc = bufio.ScanLines
		prefix = "lines: "
	case *words:
		splitFunc = bufio.ScanWords
		prefix = "words: "
	case *chars:
		splitFunc = bufio.ScanRunes
		prefix = "chars: "
	default:
		//
	}

	s := bufio.NewScanner(r)
	s.Split(splitFunc)

	count := 0
	for s.Scan() {
		fmt.Println(s.Text())
		count++
	}

	fmt.Println(prefix, count)

}
