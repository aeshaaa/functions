//Task = 5
/*create a type SQUARE
 create a type CIRCLE
 attach a method to each that calculates AREA and returns it
	 circle area= πr2
	 square area = L * W
 create a type SHAPE that defines an interface as anything that has the AREA method
 create a func INFO which takes type shape and then prints the area
 ● create a value of type square
● create a value of type circle
● use func info to print the area of square
● use func info to print the area of circle*/
package main

import (
	"fmt"
	"math"
)

type SQUARE struct {
	length float64
}

type CIRCLE struct {
	radius float64
}

func (S SQUARE) area() float64 {
	return S.length * S.length
}
func (C CIRCLE) area() float64 {
	return math.Pi * C.radius * C.radius
}

type shape interface {
	area() float64
}

func info(S shape) {
	fmt.Println(S.area())
}

var x int    // (Task = 7)
var d func() // (Task = 7)

func main() {
	square := SQUARE{
		length: 10,
	}
	circle := CIRCLE{
		radius: 13.75,
	}
	info(square)
	info(circle)

	//Task = 6
	//Build and use an anonymous func
	func() {
		for i := 0; i < 25; i++ {
			fmt.Println(i)
		}
	}()

	//Task = 7
	// Assign a func to a variable, then call that func
	f := func() {
		for j := 0; j <= 15; j++ {
			fmt.Println(j)
		}
	}
	f()
	fmt.Printf("%T\n", f)
	fmt.Printf("%T\n", x)
	d = f
	d()
	fmt.Printf("%T\n", d)

	//Task = 8
	/* Create a func which returns a func
	assign the returned func to a variable
	call the returned func*/
	//a := Aesha()
	//fmt.Println(a())

	//____________

	//Task = 9
	/* A “callback” is when we pass a func into a func as an argument. For this exercise,
	   pass a func into a func as an argument*/

	b := func(xi []int) int {
		if len(xi) == 0 {
			return 0
		}
		if len(xi) == 1 {
			return xi[0]
		}
		return xi[0] + xi[len(xi)-1]
	}
	x := foo(b, []int{10, 20, 30, 40, 50})
	fmt.Println(x)

	//Task = 10

	/*Closure is when we have “enclosed” the scope of a variable in some code block. For this
	  hands-on exercise, create a func which “encloses” the scope of a variable*/

	S := fooo()
	fmt.Println(S())
	m := fooo()
	fmt.Println(m())
	fmt.Println(m())
	fmt.Println(m())

}

func foo(f func(xi []int) int, ii []int) int {
	n := f(ii)
	n++
	return n
} // ( 10 +50 +(n++)(1)= 61)

//_____________

//func Aesha() func() int {         //( Task = 8)
// return func() int {
// 	return 7
// }
//}

// Task = 10
func fooo() func() int {
	x := 15
	return func() int {
		x++
		return x
	}
}
