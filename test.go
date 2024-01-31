package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

// func function_name( [parameter list] ) [return_types] {
//      function body
//      return [value]
// }

func add(x int, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

// naked return
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func for_loop() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
		fmt.Println(sum, i)
	}
	fmt.Println(sum)
}

func while_loop() {
	var sum int = 1
	for sum < 10 {
		sum += 1
		fmt.Println(sum)
	}
	// fmt.Println(sum)
}

func if_else(x int) {
	if x == 0 {
		fmt.Println(false)
	} else if x == 1 {
		fmt.Println(true)
	} else {
		fmt.Println(x)
	}
}

// func runtime() {
// 	fmt.Print("Go runs on ")
// 	switch os := runtime.GOOS; os {
// 	case "darwin":
// 		fmt.Println("OS X.")
// 	case "linux":
// 		fmt.Println("Linux.")
// 	default:
// 		// freebsd, openbsd,
// 		// plan9, windows...
// 		fmt.Printf("%s.\n", os)
// 	}
// }

// pointers
func pointer_sample() {
	// declare variables stored in memory
	var i, j int = 42, 2701

	// store the memory address to another variable
	// point to i
	p := &i

	// read i through the pointer
	fmt.Println(*p)

	// set i through the pointer
	*p = 21
	fmt.Println(i) // see the new value of i

	// point to j
	p = &j

	// set j through the pointer
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}

// slice
func print_slice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

// function closures
func adder() func(int) int {
	// sum is declared outside the function
	sum := 0
	// putting the function inside the function
	// the function closes over the variable sum to form a closure
	// a closure is a function value that references variables from outside its body
	return func(x int) int {
		// sum is accessible inside the function
		sum += x
		return sum
	}
}

// structs - collection of fields
type Vertex struct {
	X int
	Y int
}

type AnotherVertex struct {
	X, Y float64
}

// Methods - functions with a special receiver argument
// Adding a method to AnotherVertex struct
// methods with value receivers take either a value or a pointer as the receiver
func (v AnotherVertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Notes on methods:
// all methods on a give ntype should have either value or pointer receivers, but not a mixture of both

// reasons to use a pointer receiver for methods:
// 1. to avoid copying the value on each method call
// 2. so that the method can modify the value that its receiver points to

// Point receivers
// Methods with pointer receivers can modify the value to which the receiver points
// Since methods often need to modify their receiver, pointer receivers are more common than value receivers
// Without using * pointer receiver, you are only able to get the values from the class instance and unable to modify these values
// With * pointer receiver, you are able to get and modify the values from the class instance
func (v *AnotherVertex) Scale(f float64) {
	// methods with pointer receivers take either a value or a pointer as the receiver
	v.X = v.X * f
	v.Y = v.Y * f
}

// function instead of method
// error will occur if you try to use this function without pointer receiver
// functions with a pointer agument must take a pointer
func ScaleFunc(v *AnotherVertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func method_sample() {
	v := AnotherVertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
	ScaleFunc(&v, 10)

	// pointer receiver
	// for functions instead of methods, you need to explicitly pass the pointer
	p := &AnotherVertex{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)

}

// declaring methods on non-struct types
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func float_test() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}

// Interfaces
// interface type is defined as a set of method signatures
// a value of interface type can hold any value that implements those methods
// under the hood, interface values can be thought of as a tuple of a value and a concrete type
// (value, type) -> holds a value of a specific underlying concrete type
type Abser interface {
	Abs() float64
}

func interface_test() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := AnotherVertex{3, 4}

	a = f  // a MyFloat implements Abser
	a = &v // a *AnotherVertex implements Abser

	// In the following line, v is a AnotherVertex (not *AnotherVertex)
	// and does NOT implement Abser
	// a = v
	// AnotherVertex type (not a pointer) does not implement Abser because Abs() is defined only on *AnotherVertex (pointer)
	// it will depend on how you define the method

	fmt.Println(a.Abs())
}

// Implicit declaration of interface
// no need to explicitly declare that it implements the interface
// calling a method on a nil interface is a run-time error because there is no type inside the interface tuple to indicate which concrete method to call
// an empty interface may hold values of any type
// var i interface{} - empty interface -> (<nil>, <nil>)
// i = 42 -> (42, int)
// i = "hello" -> ("hello", string)

type I interface {
	M()
}
type T struct {
	S string
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
// This is called implicit implementation.
func (t T) M() {
	fmt.Println(t.S)
}

func implicit_interface_test() {
	var i I = T{"hello"}
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

// Handling nil interface values
// nil interface values behave like nil pointers
// calling a method on a nil interface is a run-time error because there is no type inside the interface tuple to indicate which concrete method to call
func (t *T) M2() {
	// common way to handle nil receivers while in some languages, this would trigger a null pointer exception
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func nil_interface_test() {
	var i I  // nil interface
	var t *T // nil pointer
	i = t
	describe(i) // (<nil>, *main.T)
	i.M2()      // <nil>

	i = &T{"hello"} // non-nil interface
	describe(i)     // (&{hello}, *main.T)
	i.M2()          // hello
}

// Type assertions
// provides access to an interface value's underlying concrete value
// t := i.(T) -> asserts that the interface value i holds the concrete type T and assigns the underlying T value to the variable t
// t, ok := i.(T) -> checks whether the interface value i holds the concrete type T
func type_assertion_test() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s) // hello

	s, ok := i.(string)
	fmt.Println(s, ok) // hello true

	// f := i.(float64) // panic
	// fmt.Println(f) // panic: interface conversion: interface {} is string, not float64

	f, ok := i.(float64)
	fmt.Println(f, ok) // 0 false
}

// Type switches
// type switch is a construct that permits several type assertions in series
// a type switch is like a regular switch statement, but the cases in a type switch specify types (not values), and those values are compared against the type of the value held by the given interface value
func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("Twice", v*2)
	case string:
		fmt.Println(v, "is string")
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func type_switch_test() {
	do(21)      // Twice 42
	do("hello") // hello is string
	do(true)    // I don't know about type bool!
}

// Stringers -> same concept of __str__ in python
// fmt package looks for a String method to convert the value to a string
//
//	type Stringer interface {
//		String() string
//	}
type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func stringer_test() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z) // Arthur Dent (42 years) Zaphod Beeblebrox (9001 years)
}

// Errors
// error type is a built-in interface similar to fmt.Stringer
//
//	type error interface {
//		Error() string
//	}
//
// fmt package looks for a String method to convert the value to a string
// error handling
type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func error_test() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}

// another error example
type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	return math.Sqrt(x), nil
	// return 0, nil
}

func error_test2() {
	fmt.Println(Sqrt(2))  // 1.4142135623730951 <nil>
	fmt.Println(Sqrt(-2)) // 0 cannot sqrt negative number: -2
}

// how to handle errors in Go
// 1. return error as a value
// 2. use panic to abort if error is unrecoverable
// 3. use log package to log error and continue

// generics
// functions that take an interface value as an argument must rely on type assertions to access the underlying concrete value
// template for generics
// func some_func[T any](arg T) {}
// functions can be written to work with multiple types using type parameters
// type parameters are placed in square brackets before the function name
func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		if v == x {
			return i
		}
	}
	return -1
}

// comparable - built-in interface
// type T is comparable if values of type T may be compared using the operators == and !=
// all basic types are comparable
// struct type is comparable if all its fields are comparable

func generic_test() {
	si := []int{10, 20, 15, -10}
	fmt.Println(Index(si, 15)) // 2

	sf := []float64{10.5, 20.5, 15.5, -10.5}
	fmt.Println(Index(sf, 15.5)) // 2

	ss := []string{"hello", "world", "golang"}
	fmt.Println(Index(ss, "golang")) // 2
}

// goroutines
// a goroutine is a lightweight thread managed by the Go runtime
// go f(x, y, z) -> starts a new goroutine running f(x, y, z)

// channels
// channels are a typed conduit through which you can send and receive values with the channel operator, <-
// ch <- v // SEND v to channel ch
// v := <-ch // RECEIVE from ch, and assign value to v
// data flows in the direction of the arrow
// like maps and slices, channels must be created before use
// ch := make(chan int)
// by default, sends and receives block until the other side is ready
// this allows goroutines to synchronize without explicit locks or condition variables
func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c channel
}

func goroutine_test() {
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int) // create a channel
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c channel
	fmt.Println(x, y, x+y)
}

// buffered channels
// ch := make(chan int, 100) // channel can buffer up to 100 values

func buffered_channel_test() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

// range and close
// closed -> only to make sure that no more values will be sent on it; also to terminate a range loop
// to test whether a channel has been closed, use a second parameter to the receive expression
// v, ok := <-ch
// ok is false if there are no more values to receive and the channel is closed
// close(ch) // closes the channel
// only the sender should close a channel, never the receiver
// sending on a closed channel will cause a panic
// channels aren't like files; you don't usually need to close them
// closing is only necessary when the receiver must be told there are no more values coming, such as to terminate a range loop

// fibonnacci with channels - range and close
func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; n++ {
		c <- x
		x, y = y, x+y
	}
	close(c) // command to close the channel; also the sender function
}

func range_and_close_test() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	// range iterates over values received from the channel repeatedly until it is closed
	for i := range c {
		fmt.Println(i)
	}
}

// select
// select statement lets a goroutine wait on multiple communication operations
// a select blocks until one of its cases can run, then it executes that case
// it chooses one at random if multiple are ready
// select lets you wait on multiple channel operations
func select_test() {
	tick := time.Tick(100 * time.Millisecond)  // tick channel
	boom := time.After(500 * time.Millisecond) // boom channel
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return // exit the program
		default:
			// default case happens immediately if none of the channels are ready
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

func select_test2() {

	// time go run select.go

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "two"
	}()

	for i := 0; i < 2; i++ {
		// select with two channels
		// select picks the first channel that is ready and receives from it (or sends to it)
		// if more than one of the channels are ready then it randomly picks which one to receive from
		// if none of the channels are ready, the statement blocks until one becomes available
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}

// default selection
// the default case in a select is run if no other case is ready
// use a default case to try a send or receive without blocking

// sync.Mutex
// mutex - mutual exclusion
// lock -> lock the mutex before accessing the shared variable
// unlock -> unlock the mutex after accessing the shared variable
// defer -> unlock will happen even if the function panics; to ensure the mutex wil be unlocks as the function returns

// SafeCounter is safe to use concurrently
type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

// Inc increments the counter for the given key
func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()   // lock so only one goroutine at a time can access the map c.v
	c.v[key]++    // value is accessed
	c.mu.Unlock() // unlock so other goroutines can access the map c.v
}

// Value returns the current value of the counter for the given key
func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()         // lock so only one goroutine at a time can access the map c.v
	defer c.mu.Unlock() // unlock will happen even if the function panics; to ensure the mutex will be unlocks as the function returns
	return c.v[key]     // value is accessed
}

func mutex_test() {
	c := SafeCounter{v: make(map[string]int)} // initialize SafeCounter
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey") // increment the counter for the key "somekey"
	}

	time.Sleep(time.Second)         // wait for the goroutines to finish
	fmt.Println(c.Value("somekey")) // print the current value of the counter
}

func main() {
	// fmt.Println("Hello, world!")
	// fmt.Println("The time is", time.Now())
	// fmt.Println("My favorite number is", rand.Intn(10))
	// fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
	// fmt.Println(math.Phi)
	fmt.Println(add(42, 13))

	// long declaration
	// var a, b string = "hello", "world"

	// short declaration
	a, b := "hello", "world"
	fmt.Println(swap(a, b))

	fmt.Println(split(20))

	// for_loop()
	// while_loop()

	// if_else(1)

	// runtime()

	pointer_sample()

	// struct
	fmt.Println(Vertex{1, 2})

	// struct fields - access using dot
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)

	// pointers to structs - struct fields can be accessed through a struct pointer
	// v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)

	// struct literals - it denotes a newly allocated struct value by listing the values of its fields
	var (
		v1 = Vertex{1, 2}  // has type Vertex
		v2 = Vertex{X: 1}  // Y:0 is implicit
		v3 = Vertex{}      // X:0 and Y:0
		p1 = &Vertex{1, 2} // has type *Vertex - special prefix & returns a pointer to the struct value
	)
	fmt.Println(v1, p1, v2, v3)

	// arrays - fixed length sequence of zero or more elements of a particular type
	var c [2]string
	c[0] = "Hello"
	c[1] = "World"
	fmt.Println(c[0], c[1])
	fmt.Println(c)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	// slices - dynamically-sized, flexible view into the elements of an array
	// unlike arrays, slices are typed only by the elements they contain
	// to create an empty slice with non-zero length, use the builtin make
	// make([]T, length, capacity)
	// s := make([]string, 3)
	// fmt.Println("emp:", s)
	// slice is formed by specifying two indices, low and high, separated by a colon
	// z[low : high]

	var s []int = primes[1:4]
	fmt.Println(s)

	// slices are like references to arrays - a slice does not store any data, it just describes a section of an underlying array
	// changing the elements of a slice modifies the corresponding elements of its underlying array
	// other slices that share the same underlying array will see those changes

	// slice literals - like array literals without the length
	// slice_literal := []bool{true, false, true}
	// creates an array below, then builds a slice that references it
	// [3]bool{true, false, true}

	// int slice declared and initialized
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	// boolean slice declared and initialized
	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	// struct slice
	sample_struct_slice := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(sample_struct_slice)

	// slice defaults - low bound defaults to 0, high bound defaults to length of slice
	// these slice expressions are equivalent
	// a[0:10]
	// a[:10]
	// a[0:]
	// a[:]

	// slice length and capacity
	// length - number of elements it contains
	// capacity - number of elements in the underlying array, counting from the first element in the slice. Original array length
	// length and capacity of a slice s can be obtained using the expressions len(s) and cap(s)
	slice_length_capacity := []int{2, 3, 5, 7, 11, 13}
	print_slice(slice_length_capacity)

	// slice the slice to give it zero length
	slice_length_capacity = slice_length_capacity[:0]
	print_slice(slice_length_capacity)

	// extend its length
	slice_length_capacity = slice_length_capacity[:4]
	print_slice(slice_length_capacity)

	// drop its first two values
	slice_length_capacity = slice_length_capacity[2:]
	print_slice(slice_length_capacity)

	// nil slices - zero value of a slice is nil
	// nil slice has a length and capacity of 0 and has no underlying array
	var nil_slice []int
	fmt.Println(nil_slice, len(nil_slice), cap(nil_slice))
	if nil_slice == nil {
		fmt.Println("nil!")
	}

	// creating a slice with make
	// make([]T, length, capacity)
	// s := make([]int, 5) // len(s) == 5, cap(s) == 5
	slice_make := make([]int, 5, 5)
	print_slice(slice_make)

	// slices of slices
	// a slice of a slice string
	board := [][]string{
		[]string{"1", "2", "3"},
		[]string{"4", "5", "6"},
	}
	fmt.Println(board)

	// appending to a slice
	// func append(s []T, vs ...T) []T
	board = append(board, []string{"7", "8", "9"})
	fmt.Println(board)

	// range - for loop to iterate over a slice
	// sample below skipped index
	for _, v := range board {
		fmt.Println(v)
	}

	// func Pic(dx, dy int) [][]uint8 {

	// 	// create an array of size dx,dy
	// 	// dx - horizontal length, dy - vertical height/length
	// 	a := make([][]uint8, dy)
	// 	// insert dx content
	// 	for i := 0; i < dy; i++ {
	// 		a[i] = make([]uint8, dx)
	// 	}

	// 	// Do something.
	// 	for i := 0; i < dy; i++ {
	// 		for j := 0; j < dx; j++ {
	// 			switch {
	// 			case j % 15 == 0:
	// 				a[i][j] = 240
	// 			case j % 3 == 0:
	// 				a[i][j] = 120
	// 			case j % 5 == 0:
	// 				a[i][j] = 150
	// 			default:
	// 				a[i][j] = 100
	// 			}
	// 		}
	// 	}
	// 	return a
	// }

	// func main() {
	// 	pic.Show(Pic)
	// }

	// maps - maps keys to values
	// map[key]value
	type Coordinates struct {
		Lat, Long float64
	}
	// var m map[string]Coordinates
	// m = make(map[string]Coordinates)
	// m["Bell Labs"] = Coordinates{
	// 	40.68433, -74.39967,
	// }
	// fmt.Println(m["Bell Labs"])

	// map literals - like struct literals, but the keys are required
	var m = map[string]Coordinates{
		"Bell Labs": Coordinates{
			40.68433, -74.39967,
		},
		"Google": Coordinates{
			37.42202, -122.08408,
		},
	}
	// or
	// var m = map[string]Coordinates{
	// 	"Bell Labs": {40.68433, -74.39967},
	// 	"Google":    {37.42202, -122.08408},
	// }
	fmt.Println(m)

	// mutating maps
	mutating_maps := make(map[string]int)
	// or
	// var mutating_maps2 = make(map[string]int)
	mutating_maps["Answer"] = 42    // insert element
	mutating_maps["Answer"] = 48    // update element
	delete(mutating_maps, "Answer") // delete element using key
	// elem, ok = maps[key]
	// v = element if present, otherwise zero value
	// ok = true if key is present
	vElem, ok := mutating_maps["Answer"] // test if key is present with two-value assignment
	fmt.Println("The value:", vElem, "Present?", ok)

	// function values - functions are values too
	// they can be passed around just like other values
	// function values may be used as function arguments and return values

	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))
	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))

	method_sample()
}
