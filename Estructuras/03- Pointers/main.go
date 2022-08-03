package main

import "fmt"

//Without pointer

// func Increase(v int){
// 	v++
// }

//With pointer
func Increase(v *int) {

	*v++
}

func main() {

	//1. With variable of type pointer

	v := 19

	var p1 *int
	var p2 = new(int)

	p3 := &v

	//%T Allows to print the data type of the variable

	fmt.Printf("p1: %T \n", p1)
	fmt.Printf("p2: %T \n", p2)
	fmt.Printf("p3: %T \n", p3)

	// 1. Without pointer type variable

	//var b int = 19
	//Increase(b)
	//fmt.Println("el valor de v es: ", b)

	//1.1. Example with

	var b int = 10

	Increase(&b)

	fmt.Println("the number is", b)

}
