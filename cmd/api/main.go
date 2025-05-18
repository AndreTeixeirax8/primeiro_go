package main

import "fmt"

const masculino = "masculino"
const feminino = "feminino"

const (
	maiorIdade = 18
	menorIdade = 17
)

func faixaEtaria (idade int) string {
		if idade < 18 {
			return "Menor de idade"
		} else if idade >= 18 && idade <= 65 {
			return "Adulto"
		} else {
			return "Idoso"
		}
		
	}

func main() {
	fmt.Println("Hello, World!")

	var nome string = "Lucas"
	idade := 20 //Aqui fica a variavel fica tipada automaticamente de acordo com o valor atribuido
	idade=36 // Aqui a variavel idade é reatribuida com um novo valor como um inteiro

	fmt.Println("Meu nome é", nome, "e tenho", idade, "anos.")

 if idade >= maiorIdade {
		fmt.Println("Você é maior de idade")
	}	 else {
		fmt.Println("Você é menor de idade")
	} 

	var filhos[3]string = [3]string{"Lucas", "Maria", "João"}

	for i := 0; i < len(filhos); i++ {
		fmt.Println("Filho", i+1, ":", filhos[i])
	}

	fmt.Println(faixaEtaria(50))


}
