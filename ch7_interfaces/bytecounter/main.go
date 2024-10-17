// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 173.

// Bytecounter demonstrates an implementation of io.Writer that counts bytes.
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

//!+bytecounter

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p)) // convert int to ByteCounter
	return len(p), nil
}

type WordsCounter int

func (c *WordsCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		*c += 1
	}
	return len(p), nil
}

type CountingWriterWrapper struct {
	writer       io.Writer
	writtenBytes int64
}

func (c *CountingWriterWrapper) Write(p []byte) (int, error) {
	c.writtenBytes += int64(len(p)) // convert int to ByteCounter
	c.writer.Write(p)
	return len(p), nil
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	wrapper := CountingWriterWrapper{writer: w, writtenBytes: 0}
	return &wrapper, &wrapper.writtenBytes
}

//!-bytecounter

func main() {
	//!+main
	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c) // "5", = len("hello")

	c = 0 // reset the counter
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c) // "12", = len("hello, Dolly")
	//!-main

	var words WordsCounter = 0 // reset the counter
	var text = "Dolly Milla Dania Max Charles    "
	fmt.Fprintf(&words, "   hello, %s", text)
	fmt.Println(words) // "5", = len("hello, Dolly")
	//!-main

	f, err := os.Create("testFile.txt")
	if err != nil {
		log.Fatal(err)
	}
	writer, i := CountingWriter(f)

	writer.Write([]byte("Here is a string...."))

	writer.Write([]byte("Here is a string...."))
	fmt.Println(*i)

	getProjectFromZoneUrl("https://www.googleapis.com/compute/v1/projects/chronosphere-dev/zones/us-west1-c")
}

func getProjectFromZoneUrl(zoneUrl string) string {
	items := strings.Split(zoneUrl, "/")
	project := items[len(items)-3]
	return project
}
