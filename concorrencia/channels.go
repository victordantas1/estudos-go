package main

import "fmt"


//	Um canal eh responsavel por conduzir informacoes de uma goroutine para outra
/*
func main() {
	c := make(chan int, 3)
	go produzir_c(c)
	consome_c(c)
}
*/
func produzir_c(c chan <- int) { // Controla o fluxo do canal para que ele so receba valores
	c <- 33
	c <- 44
	c <- 56
	close(c)
}

func consome_c(c <-chan int) { // Controla o fluxo do canal para que ele so possa ser consumido/lido
	for valor := range c{
		fmt.Println(valor)		
	}
}
