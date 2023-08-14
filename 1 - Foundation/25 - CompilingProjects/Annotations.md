### Compilando projetos

Para compilar o projeto é possível utilizar o comando `go build {{nome ou * (para tudo)}}`, isso fará com que o seu
código seja compilado para rodar no sistema operacional atual.

Para compilar para outros sistemas operacionais e outras arquiteturas, é possível especificar as variáveis de ambiente
antes do comando acima, desta forma:

```shell
GOOS={{sistema}} GOARCH={{arquitetura}} go build {{nome ou * (para tudo)}}
```

Por default, ao rodar um `go build`, o compilador levará em consideração o que estiver setado nas variáveis de ambiente
principais do GO. Para ver estas variáveis basta rodar o comando `go env`.

Comando para verificar todos os sistemas operacionais e arquiteturas possíveis: `go tool dist list`