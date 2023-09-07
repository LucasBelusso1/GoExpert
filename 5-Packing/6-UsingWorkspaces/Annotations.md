### Usando workspaces

Esta é uma outra alternativa quando estiver trabalhando com multiplos repositórios (módulos) que não estão publicados.
Neste caso a solução para poder utilizar pacotes não publicados é utilizar workspaces passando o caminho dos pacotes
que serão utilizados, desta forma:

```SHELL
 go work init ./math ./system
```
Este comando gerará um arquivo chamado "go.work", e é a partir dele que o GO se guiará para buscar os pacotes da sua
aplicação.
A vantagem de utilizá-lo é que não será necessário definir replaces dos pacotes, entretanto também não será possível
importar pacotes remotos de terceiros utilizando o `go mod tidy`. O ideal neste caso é sempre manter todos os pacotes
publicados para que não tenha problema para quem for utilizar.

**Observação.:** O `go.work` é um arquivo que não é publicado, podendo ser inserido no `.gitignore`.