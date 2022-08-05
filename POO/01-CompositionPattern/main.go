//El termino de objetos y herencia aquí no existe
//Para solventar esto se tiene los struct

package main

import "fmt"

//functions

func (l *liquid) isColor() {
	l.color = "black"
	fmt.Println(l.color)
}

func (b *beer) isBeer() {
	fmt.Println("Mi ceverza es de marca:", b.name,
		b.alcohol, b.color, b.country)
}

func main() {

	l := liquid{color: "yellow"}

	br := brand{country: "Aleman"}

	b := beer{"aguila", 8.0, l, br}

	b.isBeer()
}

//1.STRUCT -OBJECT
type liquid struct {
	color string
}

type brand struct {
	country string
}

//2.COMPOSITION PATTERN - HERENCY
type beer struct {
	name    string
	alcohol float32
	liquid
	brand
}
