package main

import (
	"fmt"

	"github.com/primeiro/exemplos"
)



func main() {

	/*jsonData := `{"id": 1, "nome": "Produto 1", "preco": 10.0, "quantidade": 5}`	
	var produto Produto
	err :=json.Unmarshal([]byte(jsonData), &produto)
	if err != nil {
		fmt.Println(err)
		return
	}*/

	produto := exemplos.NewProduto(1, "Produto 1", 10.0, 5)

	fmt.Println("ID:", produto.ID)
	fmt.Println("Nome:", produto.Nome)	
	fmt.Println("Preço:", produto.Preco)
	fmt.Println("Quantidade:", produto.Quantidade)
}
