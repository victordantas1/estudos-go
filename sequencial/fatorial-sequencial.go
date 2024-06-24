package main

// VICTOR DA SILVA DANTAS - 22351265
// GUILHERME TELES DA COSTA MOURA - 22353219

import (
	"fmt"
	"time"
)

func fat(num int64, parada int64, res *int64) { // Faz o fatorial de um valor
	for i := num; i > parada; i-- {
		*res *= i 
		time.Sleep(time.Second) // para a execucao por 1 segundo (so para demonstrar a diferenca de cada algoritmo)
	}
}

func main() {
	var num, res int64 = 0, 1
	fmt.Println("Insira um valor para ver o seu fatorial: ")
	fmt.Scanln(&num) 
	inicio := time.Now() // Pega o tempo no instante de inicio da execucao do algoritmo
    fat(num, 1, &res) 
	fmt.Printf("Fatorial de %d = %d\n", num, res) 
	fmt.Println("Tempo de execucao =", time.Since(inicio)) // imprime o tempo de execucao
}