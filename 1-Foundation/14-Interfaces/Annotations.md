### Interfaces

Em GO é possível criar interfaces semelhantes ao que existe em outras linguagens, definindo métodos que farão parte
desta interface e todos os structs que implementarem estas funções farão parte desta determinada interface não sendo
necessário específicar a implementação de uma interface dentro de uma struct.

Veja como criar uma interface abaixo:

```GO
type Pessoa interface {
	Desativar()
}
```

No caso do exemplo acima, a interface pessoa possui o método `Desativar()`, sendo assim, caso uma struct implemente
uma função chamada `Desativar`, automaticamente ela já fará parte da interface `Pessoa`. Sendo assim, a struct Cliente
definida na aula anterior faria parte da interface pessoa, sendo possível executar o seguinte código:

```GO
func Desativacao(pessoa Pessoa) {
	pessoa.Desativar()
}
```

No exemplo acima, o método `Desativacao()` espera uma interface do tipo Pessoa e executa o método `Desativar()`. Desta
forma, é possível executar o seguinte código:

```GO
func main() {
	wesley := Cliente{
		Nome:  "Wesley",
		Idade: 30,
		Ativo: true,
	}

	Desativacao(wesley)
}
```