package main

import "fmt"

func main() {

	//Arrays

	var numeros [5]int
	fmt.Println(numeros)

	numeros[0] = 19
	numeros[1] = 20
	numeros[2] = 21

	fmt.Println(numeros)
	fmt.Println(numeros[4])

	//Array con valores

	names := [3]string{"alex", "pepo", "Juna"}
	fmt.Println(names)

	//Dynamic Array (the length depends on  the number of elements)

	colors := [...]string{

		"red",
		"yellow",
		"black",
		"white",
	}

	//know the length of an array
	fmt.Println(colors, len(colors))

	//DEFINED INDEXES

	coins := [...]string{

		0: "Dolar",
		2: "Soles",
		5: "oro",
	}
	fmt.Println(coins, len(coins))

	//Sub Array

	subCoins := coins[3:]
	fmt.Println(subCoins)
}
