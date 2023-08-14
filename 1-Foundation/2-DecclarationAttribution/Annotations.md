### Declaração e atribuição

- Formas de declarar uma variável:

1. Declarando uma constante:
```GO
const variable = "Hello World"
```
2. Declarando uma variável:
```GO
var variable bool
```

OBS.: Neste caso é necessário utilizar a palavra reservada `var` junto com o nome da variável e o tipo.

3. Declarando múltiplas variáveis ao mesmo tempo:
```GO
var (
	verifier bool // Default false
	number int // Default 0
	text string // Default ""
	fraction float64 // Default +0.000000e+000
)
```

Quando a variável é declarada fora do escopo de funções, ela recebe o escopo global, sendo assim é possível acessá-la
em qualquer arquivo dentro do pacote. Do contrário a variável tem escopo de bloco (`{}`)

Também é possível atribuir um valor na declaração da variável, desta forma:

```GO
var (
	verifier bool    = true
	number   int     = 90
	text     string  = "Hello World!"
	fraction float64 = 5.546
)
```

Go é fortemente tipado, sendo assim, a variável é declarada com um determinado tipo e somente vai poder receber valores
deste tipo até o fim do uso desta variável.

```GO
var text string
text = 100 // Resulta em erro.
```

Go também não permite a não utilização de uma variável, ou seja, caso a variável seja declara e nunca usada, o Go irá
informar um erro.

Dentro das declarações há também a opção de usar a "Shorthand" do esquilo ou "Gopher", desta forma:

```GO
text := "Hello World!" // String (Duck type)
```