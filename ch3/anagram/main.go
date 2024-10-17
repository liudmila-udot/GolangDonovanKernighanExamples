package main

import (
	"fmt"
	"reflect"
)

func main() {
	var isAnagram = anagram("ddddf", "dddd")
	fmt.Println(isAnagram)
}

func anagram(s1 string, s2 string) bool {
	return reflect.DeepEqual(stringToMap(s1), stringToMap(s2))
}

func stringToMap(s string) map[int32]int {
	counts := make(map[int32]int)
	for _, char := range s {
		counts[char]++
	}
	return counts
}
