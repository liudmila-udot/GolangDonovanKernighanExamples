package main

import (
	"awesomeProject/ch5_functions/prettyhtml"
	"golang.org/x/net/html"
	"io"
	"log"
	"os"
)

type simpleReader struct {
	s string
}

// Read reads up to len(p) bytes into p. It returns the number of bytes
func (r *simpleReader) Read(p []byte) (n int, err error) {
	n = copy(p, r.s)
	r.s = r.s[n:]
	if len(r.s) == 0 {
		err = io.EOF
	}
	return
}

func NewReader(s string) io.Reader {
	return &simpleReader{s}
}

func main() {
	doc, err := html.Parse(NewReader("<p>Hello, world!</p>"))
	if err != nil {
		log.Fatalf("ch07/ex04: %v", err)
	}
	prettyhtml.WriteHTML(os.Stdout, doc)
}
