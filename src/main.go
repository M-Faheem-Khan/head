/*
 * Author: Faheem Kha
 * Date: 22-01-2021
 * Description: reads n number of bytes/lines from file/Stdin (Golang implementation of head program in linuxs)
 */

package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

/*
 * readBytes: Read n number of bytes from file or Stdin
 * input (io.Reader): file or Stdin reader
 * bytesToRead (int): number of bytes to read
 */
func readBytes(input io.Reader, bytesToRead int) {
	buffer := make([]byte, bytesToRead)
	_, err := io.ReadFull(input, buffer)
	if err != nil {
		fmt.Println("Error reading bytes - ", err)
	}
	fmt.Printf("%s\n", buffer)
}

/*
 * readLines: Read n number of lines from file or Stdin
 * input (io.Reader): file or Stdin reader
 * linesToRead (int): number of lines to read
 */
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
	// Creating Variables
	var input io.Reader
	var bytesToRead, linesToRead int

	// Settings Flags
	flag.IntVar(&bytesToRead, "b", 0, "number of bytes to read")
	flag.IntVar(&linesToRead, "l", 5, "number of lines to read")
	flag.Parse()

	// Opening File
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

	// Checking which flag was set
	if bytesToRead != 0 && linesToRead != 5 { // cannot set -b & -l flags together!
		fmt.Println("Error can't set bytes to read and lines to read flag together!")
		os.Exit(1)
	} else if bytesToRead == 0 { // bytes flag was not set - readLines
		readLines(input, linesToRead)
	} else { // lines flag was not set - readBytes
		readBytes(input, bytesToRead)
	}
}

// EOF
