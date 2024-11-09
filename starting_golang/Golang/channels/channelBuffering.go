package main

import "fmt"

func main() {

	messages := make(chan string, 2)

	/*
		Como e um canal Bufferizado podemos enviar valores para o canal sem precisar de uma leitura simultan
	*/

	messages <- "FLA"
	messages <- "MENGO"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
