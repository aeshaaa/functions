package main

import "fmt"

func main() {
	aaa()
	sss("How are you?")

	a := sasa("great")
	fmt.Println(a)

	x1, y1 := adsa("Aesha", "preety")
	fmt.Println(x1)
	fmt.Println(y1)

	x := 23
	y := 21
	X := (x + y)
	fmt.Println(X)

	//varadic (add x2 value)
	x2 := sum(2, 4, 6, 7, 8, 9)
	fmt.Println("Output", x2)

}

// add function methode
func aaa() {
	fmt.Println("Hello")
}

func sss(s string) {
	fmt.Println("Hey", s)
}

func sasa(k string) string {
	return fmt.Sprint("i'm gud. ", k)
}

func adsa(j string, n string) (string, bool) {
	a := fmt.Sprint(j, " ", n, `, says "Hello"`)
	b := true
	return a, b

}

//varadic parameters  ( you can create a func which tackes in unlimeted arguments) ("T" is signified v.perameters also T = type)
func sum(x2 ...int) int {
	fmt.Println(x2)
	fmt.Println(len(x2))
	fmt.Println(cap(x2))

	//fmt.Printf("%T\n", x2)

	sum := 0
	for i, v := range x2 {
		sum += v
		fmt.Println("Item in index position", i, "we are now adding", v, "to the total which is now", sum)
	}
	fmt.Println("The total is", sum)
	return sum
}
