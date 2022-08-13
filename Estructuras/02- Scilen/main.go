package main

import "fmt"

func main() {

	//1.Slicen
	//the are mutable ( )
	numberScilen := []int{

		1,
		2,
		3,
	}

	fmt.Println(numberScilen, len(numberScilen))

	//Add element to scilen

	numberScilen = append(numberScilen, 4)
	numberScilen = append(numberScilen, 5)

	fmt.Println(numberScilen)

	//Sub Slicen

	subSlicen := numberScilen[:2]

	numberScilen[0] = 0

	fmt.Println(subSlicen)
	fmt.Println(numberScilen)

	//2. Characteristics of data structures

	//punteros - pointers "where it starts where it ends"

	//Longitud - length
	//Capacidad - ability

	months := []string{

		"January",
		"February",
		"march",
		"April",
	}

	fmt.Printf("Len: %v, Cap: %v, %p \n", len(months), cap(months), months)

	months = append(months, "june")

	fmt.Printf("Len: %v, Cap: %v, %p \n", len(months), cap(months), months)

}
