package main

import "fmt"

func main() {
	// criando uma váriável do tiplo slice de string já atribuindo valores a mesma
	oláMundo := []string{
		"Olá Mundo!",
		"Hello World!",
		"Bonjour monde",
		"こんにちは世界",
	}

	for _, hello := range oláMundo {
		fmt.Println(hello)
	}

	fmt.Println("")

	// sim, não é mentira, coloquei acento na variável... Go aceita qualquer caracter UTF-8
	// mais um exemplo:
	こんにちは世界 := "Olá mundo, com a variável em japonês!"
	fmt.Println(こんにちは世界)

	fmt.Println("")

	// criando uma váriável do tiplo map de string/string já atribuindo valores a mesma
	oláMundoMap := map[string]string{
		"pt": "Olá Mundo!",
		"en": "Hello World!",
		"fr": "Bonjour monde",
		"jp": "こんにちは世界",
	}

	for lang, hello := range oláMundoMap {
		fmt.Println(lang + ": " + hello)
	}

}
