package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type limitReader struct {
	r            io.Reader
	maxRemaining int64 // max bytes remaining
}

// Read reads up to len(p) bytes into p. It returns the number of bytes
func (r *limitReader) Read(p []byte) (n int, err error) {
	if r.maxRemaining <= 0 {
		return 0, io.EOF
	}
	if int64(len(p)) > r.maxRemaining {
		p = p[0:r.maxRemaining]
	}
	n, err = r.r.Read(p)
	r.maxRemaining -= int64(n)
	return
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &limitReader{r, n}
}

func main() {
	var w io.Writer
	w = os.Stdout
	rw := w.(io.ReadWriter)
	rw1 := w.(interface{})
	fmt.Println(rw)
	fmt.Println(rw1)
	reader := LimitReader(strings.NewReader("abcdef"), 1)
	buf := make([]byte, 1024)
	reader.Read(buf)
	_, err := reader.Read(buf)
	if err != nil {
		log.Fatal("Read error", err)
	} else {
		println(string(buf))
	}
}
