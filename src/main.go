package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func readBytes(input io.Reader, bytesToRead int) {
	buffer := make([]byte, bytesToRead)
	_, err := io.ReadFull(input, buffer)
	if err != nil {
		fmt.Println("Error reading bytes - ", err)
	}
	fmt.Printf("%s\n", buffer)
}

func readLines(input io.Reader, linesToRead int) {
	buf := bufio.NewScanner(input)
	for i := 0; i < linesToRead; i++ {
		if !buf.Scan() {
			break
		}
		fmt.Println(buf.Text())
	}

	if err := buf.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading - ", err)
	}
}

func main() {
	// Flag
	var bytesToRead, linesToRead int
	flag.IntVar(&bytesToRead, "b", 0, "number of bytes to read")
	flag.IntVar(&linesToRead, "l", 5, "number of lines to read")
	flag.Parse()

	// Opening File
	var input io.Reader
	if fname := flag.Arg(0); fname != "" {
		f, err := os.Open(fname)
		if err != nil {
			fmt.Println("Error opening file - ", err)
			os.Exit(1)
		}
		defer f.Close()

		input = f
	} else {
		input = os.Stdin
	}

	if bytesToRead == 0 {
		readLines(input, linesToRead)
	} else {
		readBytes(input, bytesToRead)
	}
}

// EOF
