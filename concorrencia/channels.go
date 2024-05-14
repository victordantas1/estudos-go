package main

import "fmt"

func main() {
	c := make(chan int, 3)
	go produzir_c(c)
	for {
		valor, ok := <- c
		if ok {
			fmt.Println(valor)
		} else {
			break
		}
	}

}

func produzir_c(c chan int) {
	c <- 33
	c <- 44
	c <- 56
	close(c)
}
