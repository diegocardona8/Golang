//Paquete principal , se corre la aplicación
package main

import "fmt"

func main() {

	//Value default = 0 , float = 0
	var age int

	age = 12

	//Value default = ""
	var name string

	name = "diego"

	//Value default = false
	var status bool = true

	//Simplified variable declaration , define and assign

	work := "Programming"

	//Const

	const pi = 31.14

	fmt.Println(age, name, status, work, pi)

}
