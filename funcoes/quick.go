package main
/*
import (
	"fmt"
	"os"
	"strconv"
)

func merge_array(array1, array2 []int) []int {
	new_array := append(array1, array2...)
	return new_array
}

func particionar(numeros []int, pivo int) (menores, maiores []int) {
	for _, n := range numeros {
		if n <= pivo {
			menores = append(menores, n)
		} else {
			maiores = append(maiores, n)
		}
	}
	return menores, maiores
}

func quicksort(numeros []int) []int {

	if len(numeros) <= 1 {
		return numeros
	}

	n := make([]int, len(numeros))
	copy(n, numeros)

	indice_pivo := len(n) / 2
	pivo := n[indice_pivo]

	n = append(n[:indice_pivo], n[indice_pivo + 1:]...)
	menores, maiores := particionar(n, pivo)

	return append(append(quicksort(menores), pivo), quicksort(maiores)...)
}

func main() {

	entrada := os.Args[1:]
	numeros := make([]int, len(entrada))

	for i, n := range entrada {
		numero, err := strconv.Atoi(n)
		if err != nil {
			fmt.Printf("%s nao eh um numero valido!\n", n)
			os.Exit(1)
		}
		numeros[i] = numero
	}

	fmt.Println(quicksort(numeros))

}
*/