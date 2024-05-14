package main
import ("fmt")

func imprimir_dados(nome string, idade int16) {
	fmt.Printf("Nome: %s, Idade: %d\n", nome, idade)
}

func soma(num1, num2 int) int {
	return num1 + num2
}

/*func main() {
	imprimir_dados("Victor", 19)
	fmt.Println("Soma = ", soma(1,2))
}*/