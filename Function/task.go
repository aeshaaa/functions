// Task 1
/* create a func with the identifier foo that returns an int
   create a func with the identifier bar that returns an int and a string call both funcs
 print out their results*/
package main

import "fmt"

func main() {
	x := foo()
	y, s := bar()
	fmt.Println(x, y, s)

	xi := []int{2, 4, 6, 8, 10, 12} // (TAsk = 2)

	n := fooo(xi...)
	fmt.Println(n)

	yi := []int{1, 3, 5, 7, 9, 11, 13} // (task = 2)
	n1 := baar(yi)
	fmt.Println(n1)

	//() task = 3)  Defer means = Delay
	//(delay the execution of the function or method or an anonymous method until the nearby functions returns.)

	//fmt.Println("Hello")
	defer foc()
	fmt.Println("hey,Aesha")

}

func foo() int {
	return 7
}
func bar() (int, string) {
	return 20, "vadaj"
}

// Task 2
/*create a func with the identifier foo that
takes in a variadic parameter of type int
pass in a value of type []int into your func (unfurl the []int)
returns the sum of all values of type int passed in

create a func with the identifier bar that
takes in a parameter of type []int
returns the sum of all values of type int passed in*/

func fooo(x1 ...int) int {
	a := 0
	for _, v := range x1 {
		a += v
	}
	return a
}

func baar(x2 []int) int {
	a := 0
	for _, V := range x2 {
		a += V
	}
	return a
}

//Task = 3
/*Use the “defer” keyword to show that a deferred func runs after the func containing it exits.*/

func foc() {
	defer func() {
		fmt.Println("I'm gud...")
		fmt.Println("How about you, sirii?")
	}()
	fmt.Println("How are you?")
}
