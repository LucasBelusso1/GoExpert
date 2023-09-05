### Exportando objetos

Aqui fazendo uma associação com Orientação a Objetos, em GO definimos algo como sendo privado ou público a partir
da primeira letra, caso for **maiúscula** o dado é exportado (público), do contrário o dado não é exportado (privado).

```GO
package math

var X string = "Hello World" // Acessível
var x string = "Hello World" // Não acessível

type math struct { // Não acessível
	a int // Não acessível
	b int // Não acessível
	C int // Acessível
}

type Math1 struct { // Acessível
	d int // Não acessível
}

func NewMath(a, b int) math { // Acessível
	return math{a: a, b: b}
}

func (m math) Add() int { // Acessível
	return m.a + m.b
}
```

**Observação:** Tudo que estiver dentro do mesmo pacote pode ser acessado indepentende da letra inicial.