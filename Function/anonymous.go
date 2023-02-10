// undefind value
package main

import "fmt"

func main() {
	aaa()
	// func expression
	//f1 := func(){ ... }f()

	func() {
		fmt.Println(" Here, Anonymous value?")
	}()

	//func expression (add f2 value)
	// f2 := func(x int){ ...}f2(7)

	func(x int) {
		fmt.Println("This is Anonymous value : ", x)
	}(7)
}

func aaa() {
	fmt.Println("hey!! run")
	
}
