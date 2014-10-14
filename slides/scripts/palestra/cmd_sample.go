package main

import (
	"fmt"

	"github.com/PostmonAPI/postmongo" // HLcmd
)

func main() {
	res, err := postmongo.BuscarCep("01419101") // HLcmd
	fmt.Println(res)
	fmt.Println(err)
}
