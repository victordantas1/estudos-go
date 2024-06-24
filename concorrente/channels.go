package main

// VICTOR DA SILVA DANTAS - 22351265
// GUILHERME TELES DA COSTA MOURA - 22353219
/*
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
	res := pow(<-ch, <-ch)
	fmt.Println(res)
	wg.Done()
}

func limpa_canal(ch chan int) { // Faz a limpeza do canal caso ajam valores nao utilizados no final da execucao do programa
	fmt.Println("\nEsvaziando canal...")
	for valor := range ch {
		fmt.Println("Retirando:", valor)
	}
}

func preenche_canal(ch chan int, tam int) { // Faz o preenchimento total do canal para que seja bloqueada a inseracao
	valor := 1
	fmt.Println("Preenchendo canal...")
	for i := 0; i < tam; i++ {
		ch <- valor
		fmt.Printf("Inserido: %d\n", valor)
		valor += 1
	}
	close(ch)
}

func pow(num1, num2 int) int { // Faz a exponenciacao do valor
	res := num1
	for i := 0; i < num2 - 1; i++ {
		res *= num1
	}
	return res
}

func main() {
	wg.Add(2)
	ch := make(chan int, 2)
	fmt.Println("Insira uma base e seu expoente: ")
	go produtor(ch)
	go consumidor(ch)
	wg.Wait()
	fmt.Printf("\nUso das funcoes de preenchimento e limpeza do Canal: \n\n")
	preenche_canal(ch, 2)
	limpa_canal(ch)
}
*/