### Iniciando com structs

Uma struct no Go funciona como um tipo composto, semelhante a um objeto, que pode ter diversas propriedades e tipos.
Entretanto, vale lembrar que GO não é orientado a objetos. Exemplo de declaração de uma struct:

```GO
// Criação da struct
type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

// Utilizando a struct
func main() {
	wesley := Cliente{
		Nome:  "Wesley",
		Idade: 30,
		Ativo: true,
	}
	wesley.Ativo = false
}
```