//Task = 4
/* Create a user defined struct with
	the identifier “person”
	the fields:
	first
	last
	age
 attach a method to type person with
	 the identifier “speak”
	 the method should have the person say their name and age
create a value of type person
call the method from the value of type person */
package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (a person) speak() {
	fmt.Println("I'm", a.first, a.last, "\n", "i'm", a.age, "years old")
}

func main() {
	p1 := person{
		first: "Aesha",
		last:  "Shah",
		age:   22,
	}
	p1.speak()

}

//
