package main

func main() {

	msg := make(chan string) //Definimos msg como um canal de strings.

	/*
		Como e um canal sem Bufferizacao para enviar itens para o canal precisamos ler o que esta la.
	*/

	go func() {
		msg <- "FLAMENGO" //Atraves de uma goroutine enviamos uma string para o canal
	}()

	//message := <-msg //Vamos pegar (ler) a string do canal e armazenar em uma variaval

	//fmt.Println(message)
}
