package main

// VICTOR DA SILVA DANTAS - 22351265
// GUILHERME TELES DA COSTA MOURA - 22353219
/*
import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func fat(num int, parada int, res *int) { // faz o fatorial de um valor
	for i := num; i > parada; i-- {
		*res *= i
		time.Sleep(time.Second) // para a execucao por 1 segundo (so para demonstrar a diferenca de cada algoritmo)
	}
	wg.Done()
}

func main() {
	runtime.GOMAXPROCS(2) // define o maximo de nucleos que serao usados
	wg.Add(2) // adiciona duas threads a variavel waitGroup
	var num int
	fmt.Println("Insira um valor para ver o seu fatorial: ")
	fmt.Scanln(&num)
	res := 1
	inicio := time.Now()
	go fat(num, num / 2, &res) // faz o fatorial de uma metade do valor 
	go fat(num / 2, 1, &res) // faz o fatorial da outra metade
	wg.Wait() // aguarda a execucao das threads
	fmt.Println(res)
	fmt.Println("Tempo de execucao =", time.Since(inicio))
}
*/