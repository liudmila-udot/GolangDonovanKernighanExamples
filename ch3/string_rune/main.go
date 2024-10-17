package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	const nihongo = "日本語"
	for index, runeValue := range nihongo {
		fmt.Printf("%#U starts at byte position %d\n", runeValue, index)
	}
	fmt.Println(len(nihongo))                    // 9
	fmt.Println(utf8.RuneCountInString(nihongo)) // 3
}
