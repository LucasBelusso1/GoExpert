# Fundação

### 1 - Entendendo a primeira linha

- Todo arquivo dentro de uma mesma pasta deve declarar o mesmo "package", do contrário, ao rodar o script utilizando
`go run` será exibido um erro.
- Tudo que está dentro de uma mesmo "package" será compartilhado, mesmo que em arquivos diferentes.
- Por convenção, o nome do pacote deve ser o mesmo nome do diretório atual, com exceção do `main`, que é o ponto de
entrada do código e este deve ser nomeado como `main` e deve possui a função `main()`.

### 2 - Declaração e atribuição

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

### 3 - Criação de tipos

No Go é possível criar um tipo específico da sua preferência, da seguinte forma:

```GO
type ID int

var idContato ID

func main() {
	idContato = 2023141231234
}
```

### 4 - Importando pacote e tipagem

Para importar um pacote no Go, no início do arquivo é necessário utilizar a palavra reservada `import` juntamente como o
`package`, desta forma:

```GO
import "fmt"

import ( // Múltiplas importações.
	"fmt"
)
```

OBS.: Da mesma forma que as variáveis, não é possível importar um pacote e não utilizá-lo.

### 5 - Percorrendo arrays

Em Go, um array é uma estrutura de valores tipada e com quantidade pré definida de elementos, desta forma:

```GO
var myArray = [3]int

myArray[0] = 10
myArray[1] = 20
myArray[2] = 30
myArray[3] = 40 // Resulta em erro pois ultrapassa o limite pré definido do array
```

Neste caso o array acima tem 3 posições, começando do zero, que podem receber apenas valores do tipo int.

### 6 - Slices

Os slices funcionam mais ou menos como um array que pode aumentar de tamanho se necessário.
Declarando um slice:

```GO
var mySlice = []int{10, 20, 30, 50, 60, 70, 80, 90, 100}
```

Com a slice declarada, é possível fazer uma "Slice da Slice", informando entre colchetes `[]` o range que deseja
"extrair", desta forma:

```GO
mySlice[:0] // Remove todas as posições da Slice.
mySlice[:4] // Remove da primeira posição até a posição 4, sem inclir a posição 4. Resultando em [10 20 30 50]
mySlice[2:]) // Remove as duas primeiras posições da slice.
```

Também é possível adicionar posições a slice, utilizando a função `append()`:

```GO
mySlice = append(mySlice, 110)
```

Note que neste caso, a slice foi declarada com 9 posições, e ao adicionar mais uma posição, a capacidade da slice
aumenta para 18. Isso acontece pois ao ultrapassar o tamanho da slice, o GO dobra a capacidade.

### 7 - Maps

Maps são estruturas chave -> valor também préviamente tipadas que pode receber mais valores assim como uma slice.
Declaração de um map:

```GO
salarys := map[string]int{"Wesley": 1000, "João": 2000, "Maria": 3000}
```

Abaixo um exemplo de como adicionar uma posição a um map e um exemplo de como deletar uma posição:

```GO
delete(salarios, "Wesley") // Deleta a posição Wesley
salarios["Wes"] = 5000 // Adiciona a posição Wes com um valor de 5000
```

É possível percorrer um map da mesma forma que se percorre uma slice:

```GO
for nome, salario := range salarios {
	fmt.Printf("O salario de %s é %d\n", nome, salario)
}
```

Para ignorar um índice ou um valor, é possível utilizar o "blank identifier" `_`, desta forma:

```GO
for _, salario := range salarios {
	fmt.Printf("O salario é %d\n", salario)
}
```

### 8 - Funções

Em GO, as funções podem tanto receber quanto retornar múltiplos resultados.
Declaração de função:

```GO
func sum(a, b int) (int, error) {
	if a+b >= 50 {
		return 0, errors.New("A soma é maior que 50")
	}

	return a + b, nil
}
```

A função acima requer 2 parâmetros do tipo int e retornará um inteiro (resultado da soma) e um erro, ou `nil` caso não
haja erro.
Chamando a função:

```GO
func main() {
	result, error := sum(10, 20) // output: 30 - nill.
	result2, error2 := sum(10, 60) // output: 0 - Error.

	if (error2 != nil) {
		fmt.Println("Ocorreu um erro", error2)
	}

	result3, _ := sum(10, 60) // Ignorando caso de erro.
}
```

Exemplo de função que retorna somente um parâmetro:

```GO
func sum(a, b int) int {
	return a+bz
}
```

### 9 - Funções variádicas

Uma função variádica nada mais é do que uma função que pode receber inúmeros parâmetros. Entretanto, somente é possível
passar múltiplos valores do mesmo tipo. Por exemplo:

```GO
func sum(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}
```

No exemplo acima a função recebe inúmeros valores inteiros e retorna a soma destes valores. É possível ter mais
parâmetros de outros tipos, entretanto apenas um deles pode ser variádico e este deve ser declarado por último.

### 10 - Closures

Um closure nada mais é do que uma função anônima declarada dentro do escopo de outra função e que possui acesso as
variáveis da função "pai". Exemplo utilizando a função soma da aula anterior:

```GO
func main() {
	result := func() int {
		return sum(1, 3, 45, 6, 34, 654, 654, 7645, 534, 543, 543, 543) * 2
	}()

	fmt.Println(result)
}
```

### 11 - Iniciando com structs

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
### 12 - Composição de Structs

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

### 13 - Métodos em Structs

É possível atribuir uma função a uma determinada struct, especificando antes do nome da função, a struct a qual será
aplicada a função. Desta forma:

```GO
func (client Cliente) Desativar() {
	client.Ativo = false
	fmt.Printf("O cliente %s foi desativado", c.Nome)
}
```

### 14 - Interfaces

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

### 15 - Ponteiros

`&` -> Quando utilizado o caractere de "E comercial", o GO buscará o endereço da memória daquela variável.
Exemplo:

```GO
func main() {
	number := 10
	// Resultará em algo assim: 0xc00004a768, que é o endereço da memória em que este valor está armazenado.
	println(&number)
}
```

`*` -> Quando usado o asterísco, o GO tentará buscar o valor de um endereço de memória.
Exemplo:

```GO
func main() {
	number := 10
	pointer := &number

	// Resultará em 10, pois ele busca o valor no endereço fornecido.
	println(*pointer)
}
```

### 16 - Quando usar ponteiros

Resumidamente, usa-se ponteiros quando é necessário alterar o valor de algo por referência dentro do escopo de outra
função ou processo. Exemplo:

```GO
func soma(a, b *int) int {
	*a = 50
	*b = 50
	return *a + *b
}
```

No exemplo acima, ao informarmos os parâmetros para a função `soma()`, os valores dos parâmetros informados serão
modificados para 50 no lugar em que a função `soma()` foi executada:

```GO
func main() {
	number1 := 10
	number2 := 20
	soma(number1, number2)

	println(number1) // 50
	println(number2) // 50
}
```

### 17 - Ponteiros e structs

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

### 18 - Interfaces vazias

A interface vazia `interface{}` é uma interface que está presente em todas as structs, variáveis, tipos no geral. Sendo
assim, é possível fazer com que uma função, por exemplo, possa receber mais de um tipo em uma mesma variável. Veja o
exemplo:

```GO
func main() {
	var number int = 10
	var text string = "Hello World!"

	PrintTypeAndValue(number)
	PrintTypeAndValue(text)
}

func PrintTypeAndValue(any interface{}) {
	fmt.Printf("O valor da variável é %v e o tipo da variável é %T\n", any, any)
}
```

No exemplo acima, o parâmetro `any` recebe tanto uma string quanto um int e não retorna erro. Entretanto nestes casos
é necessário realizar uma série de validações para não violar a tipagem do GO.

### 19 - Type assertation

Em GO, é possível converter valores de tipos diferentes em outros tipos. Para isso é utilizada a seguinte notação após
a variável: `variable.({type})`, e esta conversão retorna 2 valores, o valor da conversão e se a conversão ocorreu com
sucesso:

```GO
variable := "Hello World!"
res, ok := variable.(int)

if (!ok) { // Verifica se houve erro na conversão.
	// Retorna erro.
}
```

### 20 - Generics

Generics é uma forma de conseguir passar mais de um tipo para uma função sem precisar usar a interface vazia.
Exemplo de como implementar Generics em uma função:

```GO
func Soma[T int | float64](numbers []T) T {
	var soma T
	for _, number := range numbers {
		soma += number
	}
	return soma
}
```

No exemplo acima, declaramos o nome da função como `Soma` e informamos dentro de colchetes `[]` que a função pode
receber ou um slice de int ou um slice de float64.

Também é possível definir uma interface que será um tipo com mais de um tipo, e substituir nos colchetes:

```GO
type Number interface {
	int | float64
}

func Soma[T Number](numbers []T) T {
	var soma T
	for _, number := range numbers {
		soma += number
	}
	return soma
}
```

Caso haja um tipo de seja criado com um tipo subjacente (`type AnyNumber int`), o exemplo acima irá quebrar. Para que o
GO reconheça que o tipo executado possui um tipo subjacente válido para realizar a operação, é necessário informar a
interface que pode haver essa possibilidade. Para fazer isso basta adicionar o caractere `~` a esqueda dos tipos,
da seguinte forma:

```GO
type Number interface {
	~int | ~float64
}
```

### 21 - Pacotes e módulos

Em GO, não é possível utilizar um pacote (A não ser que você esteja dentro da pasta `src/` do GO), mesmo estando no
mesmo diretório quando não declaramos um arquivo de módulo.

Utilizando o comando `go mod init {{nome}}`, será criado um arquivo `.mod` que será responsável por gerenciar os pacotes
da aplicação. A ideia é que cada repositório tenha apenas um arquivo `go.mod` que será responsável por gerenciar as
dependências do projeto.

Em GO, a visibilidade das funções, variáveis, estruturas... dos pacotes é definida a partir da primeira letra utilizada.
Caso a primeira letra seja maíscula aquele recurso é visível para outros pacotes, do contrário não.
Vale lembrar que tudo cabe nesta regra, incluíndo propriedades de structs, funções de structs entre outras coisas...

### 22 - Instalando pacotes

Para a instalação dos pacotes é necessário ter dois comandos em mente:

`go get` => Usado para importar pacotes ao projeto.
`go mod tidy` => Utilizado para otimizar a `.mod`, removendo pacotes não utilizados ou adicionando pacotes ainda não
importados.

### 23 - for

Em, GO não existe while, foreach, do while... Existe apenas o `for`. Exemplos de for:

Convencional
```GO
	for i := 0; i < 10; i++ {
		println(i)
	}
```

Looping infinito
```GO
	for {
		println("Looping infinito!")
	}
```

Looping com condição de parada
```GO
	i := 10
	for i < 100{
		println("Looping infinito!")
		i += 10
	}
```

Percorrendo slices, maps, arrays...
```GO
	sliceOfInt := []int{1, 2, 3, 4, 5, 6}

	for _, value := range sliceOfInt {
		fmt.Println(value)
	}
}
```

### 23 - Condicionais

