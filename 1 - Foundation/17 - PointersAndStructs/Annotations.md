### Ponteiros e structs

É possível fazer com que funções de uma struct altere os dados da própria struct (semelhante a usar o `this` em algumas
linguagens Orientadas a Objeto), desta forma:

```GO
type Conta struct {
	saldo int
}

func (c *Conta) simular(valor int) int {
	c.saldo += valor
	return c.saldo
}
```

No exemplo acima, caso criarmos uma variável do tipo Conta e chamarmos a função `conta.simular()`, o saldo da conta
mudará, do contrário, sem usar o caractere de asterísco `*`, o saldo seria alterado somente na função, e a variável
orinal não sofreria alteração.

Também é possível criar uma função que retorna a referência da memória para a struct:

```GO
func NewConta() *Conta {
	return &Conta{saldo: 0}
}
```