package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Uso: conversor <valores> <unidade>")
		os.Exit(1)
	}

	unidade_origem := os.Args[len(os.Args) - 1]
	valores_origem := os.Args[1 : len(os.Args) - 1]

	var unidade_destino string

	if unidade_origem == "celsius" {
		unidade_destino = "fahrenheit"
	} else if unidade_origem == "quilometros" {
		unidade_destino = "milhas"
	} else {
		fmt.Printf("%s nao eh uma unidade conhecida!! \n", unidade_origem)
		os.Exit(1)
	}

	for i, v := range valores_origem {
		valor_origem, err := strconv.ParseFloat(v, 64)
		if err != nil {
			fmt.Printf("O valor %s na posicao %d nao eh um numero valido!\n", v, i)
			os.Exit(1)
		}
		var valor_destino float64

		if unidade_origem == "celsius" {
			valor_destino = valor_origem * 1.8 + 32
		} else {
			valor_destino = valor_origem / 1.60934
		}
		fmt.Printf("%.2f %s = %.2f %s \n", valor_origem, unidade_origem,valor_destino, unidade_destino)
	}

}