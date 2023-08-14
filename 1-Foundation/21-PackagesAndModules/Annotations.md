### Pacotes e módulos

Em GO, não é possível utilizar um pacote (A não ser que você esteja dentro da pasta `src/` do GO), mesmo estando no
mesmo diretório quando não declaramos um arquivo de módulo.

Utilizando o comando `go mod init {{nome}}`, será criado um arquivo `.mod` que será responsável por gerenciar os pacotes
da aplicação. A ideia é que cada repositório tenha apenas um arquivo `go.mod` que será responsável por gerenciar as
dependências do projeto.

Em GO, a visibilidade das funções, variáveis, estruturas... dos pacotes é definida a partir da primeira letra utilizada.
Caso a primeira letra seja maíscula aquele recurso é visível para outros pacotes, do contrário não.
Vale lembrar que tudo cabe nesta regra, incluíndo propriedades de structs, funções de structs entre outras coisas...
