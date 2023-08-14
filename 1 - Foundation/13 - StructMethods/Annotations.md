### Métodos em Structs

É possível atribuir uma função a uma determinada struct, especificando antes do nome da função, a struct a qual será
aplicada a função. Desta forma:

```GO
func (client Cliente) Desativar() {
	client.Ativo = false
	fmt.Printf("O cliente %s foi desativado", c.Nome)
}
```