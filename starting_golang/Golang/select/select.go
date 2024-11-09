package main

import (
	"fmt"
	"time"
)

func main() {

	c1 := make(chan string, 1) //canal bufferizado
	c2 := make(chan string, 1) //canalbufferizado

	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "fla"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "mengo"
	}()

	//for i := 0; i < 1; i++ {
	select {
	case msg1 := <-c1:
		fmt.Println("received", msg1)
	case msg2 := <-c2:
		fmt.Println("received", msg2)
	case <-time.After(2 * time.Second):
		fmt.Println("flamengo")
	}
	//}
}
