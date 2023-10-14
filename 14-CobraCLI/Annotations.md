### Inicializando o projeto Cobra

Primeiramente acessamos o [repositório do projeto](https://github.com/spf13/cobra) e fazemos a instalação do cobra
em nossa máquina:

```SHELL
go install github.com/spf13/cobra-cli@latest
```
Depois de instalado, rodamos o comando:

```SHELL
cobra-cli init
```

Isso fará com que alguns arquivos sejam criados em nosso projeto. Sendo eles:

`main.go`: Arquivo que inicializará o cmd.
`/cmd/root.go`: Este é o arquivo principal da nossa linha de comando, é nele que ficará "acoplado" todos os outros
comandos.

Se quisermos, já podemos testar o nosso programa desta forma:

```SHELL
go run main.go help
```

### Criando nossos primeiros comandos

Para criar algum comando, podemos rodar o seguinte:

```SHELL
cobra-cli add {{comando}}
```

Sendo assim, neste caso vamos criar o comando "teste". Este comando executado acima irá criar o arquivo `/cmd/teste.go`.
Nele terá a função `init()` onde vamos definir nossas **flags** e o objeto `testeCmd` com a propriedade `run` que
executará a lógica relacionada ao nosso comando.

Na função `init()` adicionaremos a flag `command` desta forma:

```GO
testeCmd.Flags().StringP("command", "c", "", `Escolha "ping" ou "pong"`)
```

Neste caso a função `StringP` informará ao cobra que o comando **teste** poderá receber a flag `--command`, ou `-c`,
não possui valor default e a descrição é **Escolha "ping" ou "pong"**.
Agora dentro função da propriedade `Run`, escreveremos o seguinte:

```GO
command, _ := cmd.Flags().GetString("command") // Capturamos o valor da flag "command" ou "c"

if command == "ping" {
	fmt.Println("Ping")
} else {
	fmt.Println("Pong")
}
```
Agora para testar podemos rodar o seguinte:

```SHELL
go run main.go teste comando #imprimirá "pong"
go run main.go teste --comand ping #imprimirá "ping"
go run main.go teste --comand pong #imprimirá "pong"
go run main.go teste -c ping #imprimirá "ping"
go run main.go teste -c pong #imprimirá "pong"
```

### Comandos encadeados

Agora vamos deletar os comandos que criamos de teste e vamos criar o comando `category`, executando o comando:

```SHELL
cobra-cli add category
```

Agora queremos que seja possível listar categorias ou criá-las. Para isso, precisamos criar comandos encadeados para que
possamos dar um comando como `go run main.go category create\list ...`. Para isso vamos executar os seguintes comandos:

```SHELL
cobra-cli add create -p 'categoryCmd'
```

No comando acima, a flag `-p` significa "parent", ou seja, pai, então estamos vinculando o comando "create" ao comando
"category", e passamos o "categoryCmd" pois esse é o nome do objeto cmd do nosso `category.go` dentro de `cmd`.

**OBS.:** O mesmo pode ser feito apenas criando um novo comando e, no código, na função `init()` trocar o `rootCmd` por
`categoryCmd`.

### Flags locais vs flags globais

**Flags locais**: São flags que são definidas apenas para aquele comando ou subcomando e não ficam disponíveis para
nenhum outro comando. Dentro do Cobra, podemos definir uma flag local desta maneira:

```GO
categoryCmd.Flags().StringP("name", "n", "", "Name of the category")
```

No exemplo acima foi criada a flag `name` para o comando `category`. Sendo assim, se rodarmos o comando
`go run main.go category`, será exibida a flag `name` como sendo uma das opções. Entretanto se rodarmos o comando
`go run main.go category create`, a flag não será exibida.

**Flags globais**: São flags definidas em um comando que ficam disponíveis para todos os sub-comandos globalmente.
Podemos definir uma flag como sendo global desta maneira:

```GO
categoryCmd.PersistentFlags().StringP("description", "d", "", "Description of the category")
```

No exemplo acima foi criada a flag `description` para o comando `category` de forma global. Sendo assim, se rodarmos o
comando `go run main.go category` ou `go run main.go category create`, a flag será exibida como sendo uma das opções.

### Manipulando flags

Há varias formas de declarar uma flag dentro do Cobra, podendo ela mudar de tipo e possui ou não shortcuts. Veja abaixo
algumas possibilidades:

```GO
categoryCmd.Flags().StringP("name", "n", "", "Name of the category")
```
Neste caso o comando pode ser chamado destas formas:

```SHELL
go run main.go category --name {{value}}
go run main.go category -n {{value}}
go run main.go category -n # Aqui será retornado o valor default ""
```

Dentro do category, podemos capturar o valor do comando desta forma:

```GO
name, _ := cmd.Flags().GetString("name")
fmt.Println("Value of flag name:", name)
```

Agora veja exemplos utilizando outros tipos:

```GO
categoryCmd.Flags().BoolP("exists", "e", false, "Category exists")
categoryCmd.Flags().Int16("id", 0, "ID of the category")
```
Como chamar:

```SHELL
go run main.go category -e # exists igual a true
go run main.go category -e true # exists igual a true
go run main.go category -e false # exists igual a false
go run main.go category # exists igual a false

go run main.go category --id 10 # id = 10
go run main.go category -i 10 # erro pois não definimos abreviação
```
Agora todos juntos:

```SHELL
go run main.go category -n any -e --id 10 # Categoria com o nome "any" que existe e o ID é 10.
```

Para capturar dentro da função `run()`:

```GO
exists, _ := cmd.Flags().GetBool("exists") // Captura um valor booleano
fmt.Println("Value of flag exists:", exists)
id, _ := cmd.Flags().GetInt16("id") // Captura um valor Int16
fmt.Println("Value of flag id:", id)
```

### Flags mudando valor por referência

É possível definir uma flag que alterará o valor de uma variável por referência ao invés de termos que capturar o valor
dela na função `run()`. Veja o exemplo dentro de `category.go`:

```GO
var categoryName string

func init() {
	rootCmd.AddCommand(categoryCmd)
	categoryCmd.Flags().StringVarP(&categoryName, "name", "n", "", "Name of the category")
}
```
função `run()`:

```GO
Run: func(cmd *cobra.Command, args []string) {
	fmt.Println("Value of flag name:", categoryName)
}
```

### Entendendo hooks

Dentro de cada comando temos algumas opções além da função `run()` que podemos definir, são elas:

`RunE`: Mesma função do `run` porém pode retornar erro.
`PreRun`: Executa antes do run.
`PreRunE`: Executa antes do run e pode retornar um erro.
`PostRun`: Executa depois do run.
`PostRunE`: Executa depois do run e pode retornar um erro.

Veja o exemplo:

```GO
Run: func(cmd *cobra.Command, args []string) {
	fmt.Println("Value of flag name:", categoryName)
	exists, _ := cmd.Flags().GetBool("exists")
	fmt.Println("Value of flag exists:", exists)
	id, _ := cmd.Flags().GetInt16("id")
	fmt.Println("Value of flag id:", id)
},
RunE: func(cmd *cobra.Command, args []string) error {
	return fmt.Errorf("roda o comando e pode retornar erro")
},
PreRun: func(cmd *cobra.Command, args []string) {
	fmt.Println("Executa antes do run")
},
PostRun: func(cmd *cobra.Command, args []string) {
	fmt.Println("Executa depois do run")
},
PreRunE: func(cmd *cobra.Command, args []string) error {
	return fmt.Errorf("executa antes do run e pode retornar erro")
},
PostRunE: func(cmd *cobra.Command, args []string) error {
	return fmt.Errorf("Executa depois do run e pode retornar erro")
},
```

Existem estas e outras opções que podem ser vistas na documentação do cobra.

### Trabalhando com banco de dados

Aqui vamos criar, dentro do `root.go`, o seguinte:

```GO
type RunEFunc func(cmd *cobra.Command, args []string) error

func GetDb() *sql.DB {
	db, err := sql.Open("sqlite3", "./database.db")

	if err != nil {
		panic(err)
	}

	return db
}

func GetCategoryDb(db *sql.DB) database.Category {
	return *database.NewCategory(db)
}
```

E dentro de `create.go` vamos inserir o seguinte dentro da função `run()`:

```GO
db := GetDb()
category := GetCategoryDb(db)

name, _ := cmd.Flags().GetString("name") // Criar flag no init()
description, _ := cmd.Flags().GetString("description") // Criar flag no init()

category.Create(name, description)
```

Desta forma, conseguimos fazer com que seja criada uma nova categoria executando o comando:

```SHELL
go run main.go category create -n cat -d desc
```

Entretanto, nesta solução, a obtenção do banco de dados está sendo feito dentro da própria regra de negócio, o que torna
o código ruim para testes.

### Inversão de controle ao executar comandos

Para resolver o problema de código, podemos fazer o seguinte dentro de `create.go`:

```GO
func NewCreateCmd(categoryDb database.Category) *cobra.Command { // Injeção o db
	return &cobra.Command{
		Use:   "create",
		Short: "Create a new category",
		Long:  `Create a new category based on the flags "name" (-n or --name) and "description" (-d or --description)`,
		RunE:  RunCreate(categoryDb),
	}
}

func RunCreate(categoryDb database.Category) RunEFunc { // Retorna o tipo que criamos na root.go
	return func(cmd *cobra.Command, args []string) error {
		name, _ := cmd.Flags().GetString("name")
		description, _ := cmd.Flags().GetString("description")

		_, err := categoryDb.Create(name, description)

		if err != nil {
			return err
		}

		return nil
	}
}

func init() {
	createCmd := NewCreateCmd(GetCategoryDb(GetDb()))
	categoryCmd.AddCommand(createCmd)
	createCmd.Flags().StringP("name", "n", "", "Name of the category")
	createCmd.Flags().StringP("description", "d", "", "Description of the category")
}
```

Desta forma, temos mais controle e podemos injetar nossas entidades do pacote `database` para podermos fazer operações.