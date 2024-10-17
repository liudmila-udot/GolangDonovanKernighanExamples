package main

import "fmt"

func main() {
	s := []int{5, 6, 7, 8, 9}
	ret := removePreserveOrder(s, 2)
	fmt.Println(ret)

	s1 := []int{5, 6, 7, 8, 9}
	ret1 := removeNoOrder(s1, 2)
	fmt.Println(ret1)
}

func removePreserveOrder(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

func removeNoOrder(slice []int, i int) []int {
	slice[i] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}
