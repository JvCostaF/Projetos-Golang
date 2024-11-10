package main

import "fmt"

func main() {

	messages := make(chan string, 3) //Criamos um canal Bufferizado com capacidade igual a

	/*
		Como e um canal Bufferizado podemos enviar valores para o canal sem precisar de uma leitura simultanea
	*/

	go func() {
		messages <- "FLAMENGO da GoRoutine 1"
	}()

	go func() {
		messages <- "FLAMENGO da GoRoutine 2"
	}()

	// messages <- "FLAMENGO da Main"
	messages <- "FLAMENGO da Main"

	// fmt.Println(<-messages)
	// fmt.Println(<-messages)
	// fmt.Println(<-messages)

	for i := 0; i < 3; i++ {
		fmt.Println(<-messages)
	}
}
