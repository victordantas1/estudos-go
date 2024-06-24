package main

// VICTOR DA SILVA DANTAS - 22351265
// GUILHERME TELES DA COSTA MOURA - 22353219

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func produtor(ch chan int) { // Le dois valores do usuario para que seja feita a exponenciacao
	var num1, num2 int
	fmt.Scanf("%d%d\n", &num1, &num2)
	ch <- num1
	ch <- num2
	wg.Done()
}

func consumidor(ch chan int) { // Faz o processamento dos dados enviados via canal pela funcao produtora e imprime o resultado
	res := mdc(<-ch, <-ch)
	fmt.Println("MDC:", res)
	wg.Done()
}

func mdc(num1, num2 int) int { // Faz a exponenciacao do valor
	div := 1
	mdc := 0
	menor := menor(num1, num2)
	for div <= menor {
		if num1 % div == 0 && num2 % div == 0 {
			mdc = div
		}
		div++
	}
	return mdc
}

func menor(num1, num2 int) int {
	if num1 < num2 {
		return num1
	} else {
		return num2
	}
}

func main() {
	wg.Add(2)
	ch := make(chan int, 2)
	fmt.Println("Insira dois valores para ver seu MDC: ")
	go produtor(ch)
	go consumidor(ch)
	wg.Wait()
}