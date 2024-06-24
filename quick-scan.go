package main
/*
import (
	"fmt"
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

	go quicksort(menores)
	go quicksort(maiores)

	return append(append(maiores, pivo), menores...)
}

func main() {
	var tam int
	fmt.Println("Insira o tamanho do vetor")
	fmt.Scanln(&tam)
	var vet = make([]int, tam)
	for i := 0; i < tam; i++ {
		fmt.Scanf("%d", &vet[i])
	}
	array_ordenado := quicksort(vet)

	fmt.Println(array_ordenado)
}
*/