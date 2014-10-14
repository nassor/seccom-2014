package main

import "fmt"

// START OMIT
type Point struct {
	x, y int
}

// Definição de um método:
func (p Point) String() string {
	return fmt.Sprintf("(%d, %d)", p.x, p.y)
}

type Stringer interface {
	String() string
}

func main() {
	var x Stringer                // variáveis com tipo abstrato
	x = Point{2, 3}               // atribuindo um valor com um tipo não abtrato
	fmt.Println("A", x.String())  // invocando o método String() do tipo Point
	fmt.Println("C", Point{5, 1}) // fazendo a mesma coisa mas sem o método
}

// END OMIT
