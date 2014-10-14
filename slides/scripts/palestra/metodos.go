package main

import "fmt"

// START OMIT
// Definição do tipo
type Point struct {
	x, y int
}

// Definição de um método:
func (p Point) String() string {
	return fmt.Sprintf("(%d, %d)", p.x, p.y)
}

// Chamando o método
func main() {
	p := Point{2, 3}
	fmt.Println(p.String())
	fmt.Println(Point{3, 5}.String())
}

// END OMIT
