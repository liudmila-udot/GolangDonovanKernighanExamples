// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 6.
//!+

// Echo2 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
	"time"
)

/*func main() {
	s, sep := "", ""
	for index, arg := range os.Args[0:] {
		s += sep + strconv.Itoa(index) + ":" + arg
		sep = " "
	}
	fmt.Println(s)
}*/

func main() {
	start := time.Now()
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	secs := time.Since(start)
	fmt.Println(secs)
}

//!-
