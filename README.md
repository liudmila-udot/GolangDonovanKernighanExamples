Examples from the book: https://github.com/adonovan/gopl.io

In Go, := is for declaration + assignment, whereas = is for assignment only.

For example, var foo int = 10 is the same as foo := 10.

```
go run main.go
```
You can generally "trigger EOF" in a program running in a terminal with a CTRL + D keystroke right after the last input flush.

```
// reduce scope of err variable
if err := r.ParseForm(); err != nil {
	log.Print(err)
}
```
Search for existing packages:
[https://pkg.go.dev/](https://pkg.go.dev/)

## Program structure

### Names

The case of the first letter of a name determines its visibility. 
If the name begins with upper case - means it is *exported*: visible and accessible outside its own package.

Acronyms and initialism HTML or ASCII - rendered in the same case: htmlEscape or escapeHTML.

### Declarations

**var**

**const**

**type**

**func**

### Variables

```
var name type = expression
```

Either the type or =expression may be omitted.

Declare and initialize a set of variables

```
var i, j, k int
var b, s, f = true, 2.3, "true"
```

```
var f, err = os.Open(name)
```

#### Short variable declaration

```
name := expression
```

```
anim := gif.GIF{LoopCpunt: nframes}
t := 0.0
```

```
i, j := 1, 0
```

Short declaration doesn't necessarily declare all variables

```
in, err := os.Open(infile)
// ...
out, err := os.Create(outfile)
```

Must declare at least one new variable, code will not compile:

```
f, err := os.Open(infile)
// ...
f, err := os.Create(outfile) // compile error: no new variables 
```

#### Pointers

Pointer is the address of a variable. With a pointer we can read or update value *indirectly*,
without using or even knowing the name of the variable.

A pointer is a variable that references a certain location in memory 
where a value is stored, so you take a variable, 
place it on memory then instead of reading the actual value, 
you just get the address of where that value is; 

When you pass it around the code, 
since it’s just a “number” it is cheaper to move than the actual value, 
which could be way longer than just the address.

p = &x - address of x
*p - value of pointer p to x

var ptr *int - means pointer to int

Example: Below is a pointer of type string which can store only the memory addresses of string variables.

```
var s *string
```

```
x := 1
p := &x
fmp.Println(*p) // "1"
*p = 2
fmp.Println(x) // "2"
```

```
func incr (p *int) int {
 *p++ // increments what p points to; doesn't change p directly
 return *p
}

v: = 1
incr(&v)                // side effect: v is now 2
fmp.Println(incr(&v))   // "3" (and v is 3 now)
fmp.Println(incr(&v))   // "3" (and v is 3 now)
```

```
func main() {
    // Declare an int variable
    value := 42

    // Declare a pointer to int
    var ptr *int

    // Assign the address of value to ptr
    ptr = &value

    // Print the value and the pointer
    fmt.Println("Value:", value)        // Output: Value: 42
    fmt.Println("Pointer:", ptr)        // Output: Pointer: 0xc0000120a0 (example address)
    fmt.Println("Dereferenced:", *ptr) // Output: Dereferenced: 42

    // Modify the value using the pointer
    *ptr = 100
    fmt.Println("Modified Value:", value) // Output: Modified Value: 100
}
```
#### The new Function

```
p := new(int) // p, of type *int, points to an unnamed int variable
fmp.Println(*p) // "0"
*p = 2 // sets the unnamed int to 2
fmp.Println(*p) // "2"
```

#### Tuple Assignments

Swapping the values of 2 variables:

```
x, y = y, x
```

We can assign unwanted values to blank identifier

```
_, err = io.Copy(dst, src)
```
Deal with error in if block

```
f, err := os.Open()
if err != nil {
    return err
}
f.Stat()
f.Close()   
```

Compilation error

```
var cwd string;

func init() {
cwd, err := os.Getwd() // NOTE: wrong! unused local variable
if err != nil {
        log.Fatalf("os.Getwd failed: %v", err)
    }
}
```

Fix:

```
var cwd string;

func init() {
var err error
cwd, err = os.Getwd()
if err != nil {
        log.Fatalf("os.Getwd failed: %v", err)
    }
}
```

## Composite Types

* arrays
* slices
* maps
* structs

### Array:

Arrays are fixed-length sequences of items of the same type. Arrays in Go can be created using the following syntax:
```
[N]Type
[N]Type{value1, value2, ..., valueN}
[...]Type{value1, value2, ..., valueN}

var intArray = [5] int {11, 22, 33, 44, 55}
var intArray = [...] int {11, 22, 33, 44, 55}
```

Unlike in C/C++ (where arrays act like pointers) and Java (where arrays are object references),
arrays in Go are values. 

This has a couple of important implications: 
(1) assigning one array to another copies all the elements 
(2) if you pass an array to a function, it will receive a copy of the array (not a pointer or reference to it).

As you might imagine, this can be very expensive, especially when you are working with arrays that have a large number of elements.

### Slice:

Slices, on the other hand, are much more flexible, powerful, and convenient than arrays.
Unlike arrays, slices can be resized using the built-in append function. 
Further, slices are reference types, meaning that they are cheap to assign and can be passed to other functions without having to create a new copy of its underlying array. 
Lastly, the functions in Go’s standard library all use slices rather than arrays in their public APIs.

https://go.dev/blog/slices-intro

Arrays have their place, but they’re a bit inflexible, so you don’t see them too often in Go code. 
Slices, though, are everywhere. They build on arrays to provide great power and convenience.

The type specification for a slice is []T, where T is the type of the elements of the slice. 
Unlike an array type, a slice type has no specified length.

```
make([]Type, length, capacity)
make([]Type, length)
[]Type{}
[]Type{value1, value2, ..., valueN}
```

The make function takes a type, a length, and an optional capacity.
When called, make allocates an array and returns a slice that refers to that array.

```
var slice1 []int = intArray[2:5]  // index 2 to 4, size = 3
s := make([]int, 10)

var s []byte
s = make([]byte, 5, 5)
// s == []byte{0, 0, 0, 0, 0}
```

The length and capacity of a slice can be inspected using the built-in len and cap functions.

```
len(s) == 5
cap(s) == 5
```

The zero value of a slice is nil. The len and cap functions will both return 0 for a nil slice.

A slice can also be formed by “slicing” an existing slice or array. Slicing is done by specifying a half-open range with two indices separated by a colon. For example, the expression b[1:4] creates a slice including elements 1 through 3 of b (the indices of the resulting slice will be 0 through 2).

```
b := []byte{'g', 'o', 'l', 'a', 'n', 'g'}
// b[1:4] == []byte{'o', 'l', 'a'}, sharing the same storage as b
```

This is also the syntax to create a slice given an array:

```
x := [3]string{"Лайка", "Белка", "Стрелка"}
s := x[:] // a slice referencing the storage of x
```

## Type declarations

A type declaration defines a new named type that has the same underlying type as an existing type.

```
type name underlying-type
```

```
type Celcius float64
```

## Type Conversions

https://go.dev/ref/spec#Conversions

An explicit conversion is an expression of the form T(x) where T is a type and x is an expression that can be converted to type T.

```
v := typeName(otherTypeValue)
```

Example:
```
package main
import "fmt"

func main() {
  
  var intValue int = 2

  // type conversion from int to float
  var floatValue float32 = float32(intValue)
 
  
  fmt.Printf("Integer Value is %d\n", intValue)
  fmt.Printf("Float Value is %f", floatValue)

}
```

If the type starts with the operator * or <-, or if the type starts with the keyword func and has no result list,
it must be parenthesized when necessary to avoid ambiguity:

```
*Point(p)        // same as *(Point(p))
(*Point)(p)      // p is converted to *Point
<-chan int(c)    // same as <-(chan int(c))
(<-chan int)(c)  // c is converted to <-chan int
func()(x)        // function signature func() x
(func())(x)      // x is converted to func()
(func() int)(x)  // x is converted to func() int
func() int(x)    // x is converted to func() int (unambiguous)
```

A [constant](https://go.dev/ref/spec#Constants) value x can be converted to type T if x is [representable](https://go.dev/ref/spec#Representability) by a value of T. 
As a special case, an integer constant x can be explicitly converted to a [string type](https://go.dev/ref/spec#String_types) using the [same rule](https://go.dev/ref/spec#Conversions_to_and_from_a_string_type) as for non-constant x.

Converting a constant to a type that is not a [type parameter](https://go.dev/ref/spec#Type_parameter_declarations) yields a typed constant.

```
uint(iota)               // iota value of type uint
float32(2.718281828)     // 2.718281828 of type float32
complex128(1)            // 1.0 + 0.0i of type complex128
float32(0.49999999)      // 0.5 of type float32
float64(-1e-1000)        // 0.0 of type float64
string('x')              // "x" of type string
string(0x266c)           // "♬" of type string
myString("foo" + "bar")  // "foobar" of type myString
string([]byte{'a'})      // not a constant: []byte{'a'} is not a constant
(*int)(nil)              // not a constant: nil is not a constant, *int is not a boolean, numeric, or string type
int(1.2)                 // illegal: 1.2 cannot be represented as an int
string(65.0)             // illegal: 65.0 is not an integer constant
```

Converting a constant to a type parameter yields a non-constant value of that type, with the value represented as a value of the type argument that the type parameter is [instantiated](https://go.dev/ref/spec#Instantiations) with.
For example, given the function:

```
func f[P ~float32|~float64]() {
… P(1.1) …
}
```
the conversion P(1.1) results in a non-constant value of type P and the value 1.1 is represented as a float32 or a float64
depending on the type argument for f. 

Accordingly, if f is instantiated with a float32 type, the numeric value of the expression P(1.1) + 1.2 will be computed with the same precision 
as the corresponding non-constant float32 addition.

A non-constant value x can be converted to type T in any of these cases:

* x is [assignable](https://go.dev/ref/spec#Assignability) to T.
* ignoring struct tags (see below), x's type and T are not type parameters but have identical [underlying](https://go.dev/ref/spec#Underlying_types) types.
* ignoring struct tags (see below), x's type and T are pointer types that are not named types, and their pointer base types are not type parameters but have identical underlying types.
* x's type and T are both integer or floating point types.
* x's type and T are both [complex](https://golangdocs.com/complex-numbers-in-golang#:~:text=There%20are%20two%20complex%20types,The%20complex64%20and%20the%20complex128.) types.
* x is an integer or a slice of bytes or runes and T is a string type.
* x is a string and T is a slice of bytes or runes.
* x is a slice, T is an array or a pointer to an array, and the slice and array types have identical element types.

### Implicit type conversions
Unlike other languages, Go doesn't support implicit type conversion.
Although when dividing numbers, implicit conversions happen depending on the scenario. 
So we need to be very careful about what type to use where.

```
package main
import "fmt"

func main() {
  
  // initialize integer variable to a floating-point number
  var number int = 4.34

  fmt.Printf("Number is %g", number)
}

// Compilation error:

// '4.34' (type untyped float) cannot be represented by the type int
```

## Errors

Panic - sign of a bug in the calling code and should never happen in a well-written program.

## Error Handling Strategies

### Propagate Error

```
resp, err := http.Get(url)

if err != null {
    return nil, error
}
```

or with additional info

```
doc, err := html.Parse(resp.Body)

if err != null {
	return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
}
```

### Retry

ch5/wait

log.Fatal equals to

### Print error and stop the program

```
fmt.Printf
os.Exit(a)
```

### Log the error and then continue

### Ignore the error

##  Deferred functions

The function and argument expressions are evaluated when the statement is executed, but
the actual call is deferred until the function, that contains deferred statement has finished, 
whether normally with return or with panicking.

Any number of calls may be deferred: they are executed in the reverse order in which they were deferred.
LIFO order.

Defer can change returned value for a named return value.

## Panic

Go's type system catches many mistakes at compile time. But others like an out-of-bounds array access require checks at run time.
When the Go runtime detects these mistakes, it panics.

During a typical panic, normal execution stops, all deferred functions call in that goroutine are executed,
and the program crashes with a log message.

Not all panics come from runtime. The built-in panic function may be called directly.


```

switch s := suit(drawCard()); s {
  case "Spades": //
  case "Hearts": //
  case "Diamonds": //
  case "Clubs": //
  default:
    panic(fmt.Sprintf("invalid suit %q", s))
}
```
## Methods

Method is a function with a receiver.

Receiver - struct or pointer to struct.

From https://medium.com/globant/go-method-receiver-pointer-vs-value-ffc5ab7acdb:

A method is just a function with a receiver argument. It is declared with the same syntax with the addition of the receiver.

```
func (p *Person) isAdult bool {
return p.Age > 18
}
```

In the above method declarations, we declared the *isAdult* method on the **Person* type.

Now we will see the difference between the **Value receiver** and **Pointer receiver**.

**Value receiver** makes a copy of the type and pass it to the function. The function stack now holds an equal object but at a different location on memory. That means any changes done on the passed object will remain local to the method. The original object will remain unchanged.

**Pointer receiver** passes the address of a type to the function. The function stack has a reference to the original object. So any modifications on the passed object will modify the original object.

```
package main
import (
"fmt"
)
type Person struct {
Name string
Age  int
}
func ValueReceiver(p Person) {
p.Name = "John"
fmt.Println("Inside ValueReceiver : ", p.Name)
}
func PointerReceiver(p *Person) {
p.Age = 24
fmt.Println("Inside PointerReceiver model: ", p.Age)
}
func main() {
p := Person{"Tom", 28}
p1:= &Person{"Patric", 68}
ValueReceiver(p)
fmt.Println("Inside Main after value receiver : ", p.Name)
PointerReceiver(p1)
fmt.Println("Inside Main after value receiver : ", p1.Age)
}
```

Output:

Inside ValueReceiver :  John

Inside Main after value receiver :  Tom

Inside PointerReceiver :  24

Inside Main after pointer receiver :  24

This shows that the method with value receivers modifies a copy of an object.

And the original object remains unchanged. 

**So how to choose between Pointer vs Value receiver?**

If you want to change the state of the receiver in a method, manipulating the value of it, use a **Pointer receiver**. It’s not possible with a **Value receiver**, which copies by value. Any modification to a **Value receiver** is local to that copy. If you don’t need to manipulate the receiver value, use a **Value receiver**.

The **Pointer receiver** avoids copying the value on each method call. This can be more efficient if the receiver is a large struct,

**Value receivers** are concurrency safe, while **Pointer receivers** are not concurrency safe. Hence a programmer needs to take care of it.

Notes

1. Try to use the same receiver type for all your methods as much as possible.
2. If state modification needed, use **Pointer receiver** if not use **Value receiver**.

## Interfaces

### Interface types

An interface type specifies a set of methods that a concrete type must possess to be considered an instance of that interface

```
package main

import (
    "fmt"
    "math"
)

type geometry interface {
    area() float64
    perim() float64
}

type rect struct {
    width, height float64
}
type circle struct {
    radius float64
}

func (r rect) area() float64 {
    return r.width * r.height
}
func (r rect) perim() float64 {
    return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
    return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
    return 2 * math.Pi * c.radius
}

func measure(g geometry) {
    fmt.Println(g)
    fmt.Println(g.area())
    fmt.Println(g.perim())
}

func main() {
    r := rect{width: 3, height: 4}
    c := circle{radius: 5}

    measure(r)
    measure(c)
}
```

### Assignability

A type satisfies an interface as it possesses all the methods the interface requires.

Assignability rule: an expression may be assigned to an interface only if its type satisfies the interface.


```
var w io.Writer
w = new(bytes.Buffer) // OK: *bytes.Buffer has Write method
w = time.Second // compile error: time.Duration lacks Write method
```

The rule applies even when the right-hand side is itself an interface.

```
var rwc io.ReadWriteCloser
w = rwc // OK io.ReadWriteCloser has Write method
rwc = w // compile error: io.Writer lacks Close method
```

A value of type T can not define all the methods that a *T pointer does, and as a result it might satisfy fewer interfaces.

### Empty interface

Empty interface: type interface {}.

We can assign any value to empty interface.

```
var any interface{}
any = true
any = 12.34
any = "hello"
```

For example fmt.Println accepts

```
// any is an alias for interface{} and is equivalent to interface{} in all ways.
type any = interface{}
```

We need a way to get value back out again. We'll see how to do that using type assertion.

# Comparison: interfaces equality

Interface value has two components: dynamic type and a value of that type.

Dynamic Type: The specific type of the value that the interface is currently holding.
Dynamic Value: The actual value that matches the dynamic type.

**Dynamic Type**
The dynamic type is the actual type of the value stored in the interface at runtime.
When you assign a value to an interface, Go records the type of the value along with the value itself.
Example:
```
var i interface{}
i = 42 // Assign an integer

fmt.Printf("Dynamic Type: %T\n", i) // Output: Dynamic Type: int
Here, the dynamic type of i is int because an integer value (42) was assigned to it.
```

**Dynamic Value**
The dynamic value is the actual value stored inside the interface.
This value must match the interface's dynamic type.
Example:
```
var i interface{}
i = 42 // Assign an integer

fmt.Printf("Dynamic Value: %v\n", i) // Output: Dynamic Value: 42
Here, the dynamic value is 42.
```

```
func main() {
    var i interface{} // Declare an empty interface

    i = 42 // Assign an integer
    fmt.Printf("Type: %T, Value: %v\n", i, i) // Output: Type: int, Value: 42

    i = "hello" // Assign a string
    fmt.Printf("Type: %T, Value: %v\n", i, i) // Output: Type: string, Value: hello

    i = 3.14 // Assign a float
    fmt.Printf("Type: %T, Value: %v\n", i, i) // Output: Type: float64, Value: 3.14
}
```

Two interface values are equal if both are nil, or if their dynamic types are identical and their dynamic values are equal according
to the usual behaviour of == for that type.

However, if two interface are compared and have the same dynamic type, but that type is not comparable (a slice, for instance),
then the comparison fails with a panic:

```
var x interface{} = []int{1, 2, 3}
fmr.Println(x == x) // panic: comparing uncomparable type []int
```

In this respect, interface types are unusual. Other types are either safely comparable (like basic types and pointers)
or not comparable at all (like slices, maps and functions). A similar risk exists when using interfaces as map keys or 
switch operands.

Only compare interface values if you are certain that they contain dynamic values of comparable types. 

To debug dynamic type of interface dynamic value:
```
var w io.Writer
w = os.Stdout
fmt.Printf("%T\n, w") // "*os.File"
```

### Type Assertions

A type assertion is an operation applied to an interface value. 
```
x.(T)
```
x is an expression of interface type

T is a type, called the "asserted" type

A type assertion checks that the dynamic type of its operand matched the asserted type.

There are two possibilities:

1. Asserted type T is a concrete type. Then the type assertion checks whether x's dynamic type is T. Otherwise it panics.

```
var w io.Writer
w = os.Stdout
f := w.(*os.File) // success: f == os.Stdout
c: = w.(*bytes.Buffer) // panic: interface holds *os.File, not *bytes.Buffer
```
2. Asserted type T is an interface.

Type assertion checks whether x's dynamic type satisfies T.
Then dynamic value stays the same, the result is still an interface, but has type T.

```
var w io.Writer
w = os.Stdout
rw := w.(io.ReadWrite) // sucess: *os.File has both Read and Write

// ReadWriter is the interface that groups the basic Read and Write methods.
type ReadWriter interface {
	Reader
	Writer
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

w = new(ByteCounter)
rw = w.(io.ReadWriter) // panic: *ByteCounter has no Read method
```

A type assertion to a less restrictive interface type (one with a fewer methods) is rarely needed.

Type assertion with two results doesn't panic, but return boolean indication of success

```
var w io.Writer = os.Stdout
f, ok := w.(*os.File) // success: ok, f == os.Stdout
b, ok := w.(*bytes.Buffer) // failure: !ok, b == nil
```

The extended form of if:

```
if f, ok := w.(*os.File); ok {
 // ... use f ...
}
```

Or with shadowing

```
if w, ok := w.(*os.File); ok {
 // ... use w ...
}
```

Example

```
pod, ok := object.(*corev1.Pod)
if !ok {
    return false, errors.New("cannot assert Pod resource type to *corev1.Pod")
}
```

### Error types Assertions

```
if _, ok := err.(watch.InitValueError); ok {
	w.logger.Warn("invalid placement update, continue to watch for placement updates",
	zap.Error(err))
	return nil
}
```

### Type switches

```
switch x.(type) {
    case nil:
    case int, uint:
    default
    
// transformAndCollect is a global transform function that is called for every imported object.
func (imp *Importer) transformAndCollect(object client.Object) (skip bool) {
	switch obj := object.(type) {
	case *corev1.Node:
		// Save a copy of the original object, before any transformations are applied.
		node := obj.DeepCopy()
		if skip := imp.resultBuilder.checkNode(node); skip {
			return true
		}
	case *corev1.Pod:
		// Save a copy of the original object, before any transformations are applied.
		pod := obj.DeepCopy()
		if skip := imp.resultBuilder.checkPod(pod); skip {
			return true
		}
	}
	return false
}    
```

## Concurrency

### Goroutines

A goroutine is a lightweight thread managed by the Go runtime.

When program starts, it has a single goroutine, called the main goroutine.
New goroutines are created by the go statement.

A go statement causes the function to be called in a newly created goroutine.

```
f() // call f(); wait for it to return
go f() // create a new goroutine that calls f(); don't wait
```

When main function returns, all goroutines are terminated and the program exits.

Example:
```
func printNumbers(prefix string) {
	for i := 1; i <= 3; i++ {
		fmt.Printf("%s: %d\n", prefix, i)
		time.Sleep(500 * time.Millisecond) // Simulate work
	}
}

func main() {
	// Start a goroutine
	go printNumbers("Goroutine")

	// Main function continues running in parallel
	printNumbers("Main")

	fmt.Println("Done!")
}
```

Output:

```
Main: 1
Goroutine: 1
Main: 2
Goroutine: 2
Main: 3
Goroutine: 3
Done!
```

### Channels

If goroutines are the threads of Go, **channels** are the connections between them.
A channel is a typed conduit through which you can send and receive values with the channel operator, <-.

To create a channel, use the make function:

```
ch := make(chan int)
```

Channel is a reference type, like a map or a slice. When we pass a channel to a function, 
we are passing a reference to the channel.

Channel has two principal operations: send and receive. Both operations are written using the <- operator.

```
ch <- v // Send v to channel ch.
x = <-ch // Receive from ch, and assign value to x.
<-ch // Receive from ch, and discard the received value.
```

Channels support a third operation: close, which sets a flag indicating that no more values will be sent on this channel.
Subsequent attempts to send will panic.
Receive operations on a closed channel yield the values that have been sent until no more values are left,
after which any receive will yield the zero value of the channel's element type.

To close a channel, use the close function:
```
close(ch)
```

Channel unbuffered by default. If pass non zero capacity to make, it will be buffered.

```
ch:=make(chan int) // unbuffered channel
ch := make(chan int, 100) // buffered channel with capacity 100
```
### Unbuffered Channels
Unbuffered channels are synchronous. 

**Blocking Nature:**
Sending on an unbuffered channel blocks the sender until another goroutine receives the data.
Receiving from an unbuffered channel blocks the receiver until another goroutine sends data.

**No Buffer:**
Data is directly passed from the sender to the receiver without being stored.

Only one value at a time can be in transit through the unbuffered channel.

When a value is sent on an unbuffered channel, the sender blocks until another goroutine receives the value.
The receipt of the value *happens before* the receiver calls the reawakening of the sender.

In concurrency, when we say *x happens before y*, it means that all x effects, such as updates to variables, are guaranteed to be observed by y.

#### Pipelines
Channels can be used to connect goroutines together so that the output of one is the input of another.
It is called *pipelines*.

```
func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	go func() {
		for x := 0; ; x++ {
			naturals <- x
		}
	}()

	// Squarer
	go func() {
		for {
			x := <-naturals
			squares <- x * x
		}
	}()

	// Printer (in main goroutine)
	for {
		fmt.Println(<-squares)
	}
}
```

There is no way to test directly whether a channel has been closed,
but there is a variant of the receive operation that produces two results: 
the received value, and a boolean value that reports whether the channel is open or closed.

```
//Squarer
go func() {
    for {
        x, ok := <-naturals
        if !ok {
            break // channel was closed and drained
        }
        squares <- x * x
    }
    close(squares)
}()
```

Syntax above is clumsy, so Go provides a *range clause* for channels.
This is more convenient syntax to receive values from a channel and terminating the loop after the last one.

```
func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	go func() {
		for x := 0; x < 100; x++ {
			naturals <- x
		}
		close(naturals)
	}()

	// Squarer
	go func() {
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()

	// Printer (in main goroutine)
	for x := range squares {
		fmt.Println(x)
	}
}
```

#### Unidirectional Channel Types

We can break up previous example into smaller pieces. We'll define 3 functions instead of 3 go routines local variables.

```
func counter(out chan<- int)
func squarer(out chan<- int, in <-chan int)
func printer(in <-chan int)
```


Go provides unidirectional channel types that restrict the direction of data flow.
Violations are detected at compile time.

```
func counter(out chan<- int) {
	for x := 0; x < 100; x++ {
		out <- x
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for v := range in {
		out <- v * v
	}
	close(out)
}

func printer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go counter(naturals)
	go squarer(squares, naturals)
	printer(squares)
}
```