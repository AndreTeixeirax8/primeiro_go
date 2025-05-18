package exemplos

type Produto struct {
	ID         int     `json:"id"`
	Nome       string  `json:"nome"`
	Preco      float64 `json:"preco"`
	Quantidade int     `json:"quantidade"`
}

func NewProduto(id int, nome string, preco float64, quantidade int) Produto {
	return Produto{
		ID:         id,
		Nome:       nome,
		Preco:      preco,
		Quantidade: quantidade,
	}
}

func (p *Produto) GetPrecoTotal() float64 {
	return p.Preco * float64(p.Quantidade)
}