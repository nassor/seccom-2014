package main

import (
	"fmt"
	"time"
)

// START OMIT
func main() {

	// Cria um slice de slices que recebe qualquer tipo de dado.
	var aceitaQualquerCoisa [][]interface{}

	// Adicionando slices genérico dentro de um slice.
	aceitaQualquerCoisa = append(aceitaQualquerCoisa, []interface{}{time.Now(), "Nassor"})
	aceitaQualquerCoisa = append(aceitaQualquerCoisa, []interface{}{"Flavio", 100.0, time.Now().AddDate(-50, 0, 0)})

	fmt.Println(aceitaQualquerCoisa)
}

// END OMIT
