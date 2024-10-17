package main

import "fmt"

// f returns 1
func f() (result int) {
	defer func() {
		result++
	}() // why and how?
	return 0
}

func double(x int) (result int) {
	ret := 0
	// defer can change function result for a named result variable
	defer func() { fmt.Printf("double(%d) = %d\n", x, result); result = 89 }()
	ret = x + x
	//x = 6

	/*	Defer mechanism itself is simple and yet it is very powerful when used properly.
		When defer statement is called, it pushes the function call into a stack list and uses the LIFO principle.
		Hence, if a function states multiple defer,
		the last pushed defer will be called first as shown in the example below.*/
	i := 0
	for ; i < 5; i++ {
		defer fmt.Printf("%d ", i)
		//defer func() { fmt.Printf("%d ", i) }()
	}
	i = 7

	return ret
}

func triple(x int) (result int) {
	defer func() { result += x }()
	return double(x)
}

func main() {
	fmt.Println(double(2))
	fmt.Println(triple(4))
}
