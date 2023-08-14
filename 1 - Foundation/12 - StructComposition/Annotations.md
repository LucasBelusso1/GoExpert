### Composição de Structs

Dentro de uma struct, é possível adicionar outra struct como uma propriedade ou atributo do tipo de outra struct.
Exemplo:

```GO
type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
	Address Endereco
}

func main() {
	wesley := Cliente{
		Nome:  "Wesley",
		Idade: 30,
		Ativo: true,
	}

	wesley.Endereco.Cidade = "Imbiuba"
	wesley.Cidade = "Ubatuba " // Neste caso funciona pois não existe um campo chamado "Cidade" dentro da struct Cliente

	wesley.Address.Cidade = "Imbituba"
}
```
