package main

import (
	"fmt"
	"time"
)

// Além da msg e o tempo de duração, um channel também é passado por parâmetro
func f(msg string, delay time.Duration, ch chan string) {
	for {
		ch <- msg
		time.Sleep(delay)
	}
}

// START OMIT
// Chamando funções concorrentemente selecionando o
// destino das mesmas
func main() {
	chA := make(chan string)
	chB := make(chan string)
	chC := make(chan string)
	go f("A--", 300*time.Millisecond, chA)
	go f("-B-", 500*time.Millisecond, chB)
	go f("--C", 1100*time.Millisecond, chC)

	for i := 0; i < 20; i++ {
		select {
		case v := <-chA:
			fmt.Println("Channel A: ", v)
		case v := <-chB:
			fmt.Println("Channel B: ", v)
		case v := <-chC:
			fmt.Println("Channel C: ", v)
		}
	}
}

// END OMIT
