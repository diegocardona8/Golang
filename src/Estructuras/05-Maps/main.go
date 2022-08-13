package main

import "fmt"

func main() {

	days := make(map[int]string)

	fmt.Println(days)

	//Add Data

	days[1] = "Monday"
	days[2] = "saturday"

	fmt.Println(days)

	//Modify value

	days[2] = "Tuesday"

	fmt.Println(days)

	//Delete value

	delete(days, 2)
}
