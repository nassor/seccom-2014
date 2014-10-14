package main

import (
	"fmt"
	"time"
)

// START OMIT
// Além da msg e o tempo de duração, um channel também é passado por parâmetro
func f(msg string, delay time.Duration, ch chan string) {
	for {
		ch <- msg
		time.Sleep(delay)
	}
}

// Chamando funções concorrentemente
func main() {
	ch := make(chan string)
	go f("A--", 300*time.Millisecond, ch)
	go f("-B-", 500*time.Millisecond, ch)
	go f("--C", 1100*time.Millisecond, ch)

	for i := 0; i < 20; i++ {
		fmt.Println(i, <-ch)
	}
}

// END OMIT
