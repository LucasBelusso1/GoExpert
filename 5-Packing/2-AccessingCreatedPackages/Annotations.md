### Acessando pacotes que foram criados

Para acessar algum outro pacote que tenha sido criado dentro do módulo, basta importá-lo utilizando o nome que demos
ao executar o comando `go mod init {{nome}}`.

Por exemplo, se tivessemos o pacote `math`, o import ficaria desta forma:

```GO
import ("github.com/LucasBelusso1/GoExpert/2-AccessingCreatedPackages/math")
```

Neste caso o nome do nosso módulo é **github.com/LucasBelusso1/GoExpert/2-AccessingCreatedPackages**.

**Observação:** É importante que o nome do módulo tenha a url do repositório em que o código está armazenado, para
evitar conflito entre nomes.