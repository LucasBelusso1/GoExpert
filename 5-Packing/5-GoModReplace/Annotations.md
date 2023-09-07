### Trabalhando com go mod replace

Para trabalhar com pacotes que ainda não foram publicados em nenhum repositório remoto, mas que existem no seu
repositório local, é necessário informar ao GO que quando fizeremos o `get` deste pacote, ele deverá buscar dentro do
nosso repositório local, e não do remoto. Para isso, utiliza-se o seguinte comando:

```SHELL
go mod edit -replace {{repoRemoto}}={{caminhoRepoLocal}}
```

Exemplo:
```SHELL
go mod edit -replace github.com/LucasBelusso1/GoExpert/5-GoModReplace/math=../math
```
Feito isso, basta rodar o comando `go mod tidy` para fazer o require do pacote.

Entretanto, esta solução possui um problema: caso você queira disponibilizar o pacote que você está trabalhando, será
necessário publicar também o pacote que foi feito o replace (no caso acima o **math**) e alterar no `go.mod` para que
passe a utilizar o pacote remoto.