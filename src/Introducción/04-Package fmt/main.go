package main

import "fmt"

func main() {

	hola := "hola"
	mundo := "mundo"

	//each data is printed on the same line

	fmt.Print(hola)
	fmt.Print(mundo)

	//FORMAT INFORMATION (INTERPOLATION)

	//special characters ( we know the type of data to print)
	//first string : $s and second int $d

	firstName := "Diego"
	Age := 24

	fmt.Printf("Hola: %s y su edad es %d \n ", firstName, Age)

	//special characters ( we do not know the type of data to print)

	fmt.Printf("Hola: %v y su edad es %v \n ", firstName, Age)

	//Store a message in a variable (Sprintf)
	message := fmt.Sprintf("Hola: %v y su edad es %v", firstName, Age)

	fmt.Println(message)

	//Know the data type of a variable

	fmt.Printf("Type of variable: %T \n", firstName)

	//Enter data by  keyboard

	fmt.Print("Ingrese otro nombre : ")
	fmt.Scanln(&firstName)

	fmt.Println("El nuevo nombres es : ", firstName)
}
