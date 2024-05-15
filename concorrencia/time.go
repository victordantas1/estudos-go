package main
/*
import  (
	"fmt"
	"time"
)

func executar(c chan <- bool) {
	time.Sleep(1 * time.Second)
	c <- true
}

func main() {
	c := make(chan bool)
	go executar(c)
	fmt.Println("Executando...")
	fim := false
	for !fim {
		select {
		case fim = <- c: // Caso a funcao executar atribua o valor ao canal antes do timeout ir imprimir "Fim"
			fmt.Println("Fim!")
		case <-time.After(2 * time.Second): // Apos dois segundos a funcao After() ira retornar um canal booleano e sera parada a execucao do laco
			fmt.Println("Timeout!")
			fim = true
		}
	}
}
*/