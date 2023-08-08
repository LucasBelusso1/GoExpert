package main

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

func main() {
	wesley := Cliente{
		Nome:  "Wesley",
		Idade: 30,
		Ativo: true,
	}
	wesley.Ativo = false
}
