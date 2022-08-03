package main

import "fmt"

//Show value
func GetName(firstName string) {

	fmt.Println("Hi There", firstName)

}

//Return value
func sum(a, b int) int {

	return a + b
}

func main() {

	GetName("Dieguito")
	r := sum(1, 2)
	fmt.Printf("La suma es %d", r)

}
