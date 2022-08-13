package main

import "fmt"

func main() {

	number := make([]int, 3, 3)

	number[0] = 100
	number[1] = 200
	number[2] = 800
	fmt.Println(number, len(number), cap(number))
}
