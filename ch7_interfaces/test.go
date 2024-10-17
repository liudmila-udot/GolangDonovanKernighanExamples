package main

import (
	"log"
	"os"
)

func main() {
	path := "sample"
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		log.Println(err)
	}
}
