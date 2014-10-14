package main

import (
	"fmt"
	"time"
)

// START OMIT
// Função a ser executada
func f(msg string, delay time.Duration) {
	// while(true)
	for { // HL
		fmt.Println(msg)  // HL
		time.Sleep(delay) // HL
	} // HL
}

// Chamando funções concorrentemente
func main() {
	go f("A--", 300*time.Millisecond)  // HL
	go f("-B-", 500*time.Millisecond)  // HL
	go f("--C", 1100*time.Millisecond) // HL
	time.Sleep(5 * time.Second)        // irá executar o programa por 5 segundos
}

// END OMIT
