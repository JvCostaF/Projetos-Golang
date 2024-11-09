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
	consumers := 100   // valor qualquer
	producers := 100   // valor qualquer
	bufferSize := 1000 // valor qualquer
	stream := make(chan string, bufferSize)
	for i := 0; i < consumers; i++ {
		go consumidor(stream)
	}
	for i := 0; i < producers; i++ {
		go produtor(stream)
	}
	select {}
}
