package main
/*
import "fmt"

func main() {
	i, p := make(chan int), make(chan int)
	pronto := make(chan bool)
	nums := []int{1, 23, 42, 5, 8, 6, 7, 4, 99, 100}

	go separar(nums, i, p, pronto)

	var impares, pares []int
	fim := false

	for !fim {
		select {
			case n := <- i: // Caso a funcao separar() atribua um valor para o canal i ele sera inserido na lista de impares
				impares = append(impares, n)
			case n := <- p: // Caso a funcao separar() atribua um valor para o canal p ele sera inserido na lista de pares
				pares = append(pares, n)
			case fim = <- pronto: // Caso a funcao separar() atribua um valor para o canal pronto o loop eh finalizado
		}
	}

	fmt.Printf("Impares: %v | Pares: %v\n", impares, pares)
}*/
func separar(nums []int, i, p chan <- int, pronto chan <- bool) {
	for _, n := range nums {
		if n % 2 == 0 {
			p <- n
		} else {
			i <- n
		}
	}
	pronto <- true
}