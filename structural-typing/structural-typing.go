/*

	Structural Typing

	see:

		https://yourbasic.org/golang/interfaces-explained/
		https://medium.com/higher-order-functions/duck-typing-vs-structural-typing-vs-nominal-typing-e0881860bf10

	notes:

		duck typed			Python, Javascript
		nominal typed		C++, Java
		structural typed	Go

*/

package main

import "fmt"
import "strconv"

func main() {
	emptyInterface()
	stringerInterface()
}

// Temp represents a temperature in Celcius
type Temp int

func (t Temp) String() string {
	return strconv.Itoa(int(t)) + " °C"
}

// Point represents a simple 2 dimension point
type Point struct {
	x, y int
}

func (p *Point) String() string {
	return fmt.Sprintf("(%d,%d)", p.x, p.y)
}

func emptyInterface() {

	var x interface{}

	x = 1.2
	fmt.Println(x)

	x = &Point{3, 4}
	fmt.Println(x)
}

// MyStringer represents types that can return a string representation
type MyStringer interface {
	String() string
}

func stringerInterface() {

	var x MyStringer

	x = Temp(56)
	fmt.Println(x)

	x = &Point{7, 8}
	fmt.Println(x)
}
