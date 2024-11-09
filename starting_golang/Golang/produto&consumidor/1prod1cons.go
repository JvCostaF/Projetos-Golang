package main

import (
	"fmt"
)

func produtor(stream chan string) {
	var str string
	for {
		fmt.Scanln(&str)
		stream <- str
	}
}

func consumidor(stream chan string) {
	for {
		str := <-stream
		fmt.Println("Valor recebido: ", str)
	}
}

func main() {
	stream := make(chan string, 1) //Canal compartilhado pelas goroutines
	go consumidor(stream)
	go produtor(stream)
	select {} //Impede o fim da execucao do fluxo principal apos a criacao das goroutines, com um select sem clausula.
}
