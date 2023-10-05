### Trabalhando com repositórios privados

## GOPRIVATE

O GOPRIVATE é uma variável de ambiente do GO que por padrão é vazia. Nela podemos informar que iremos utilizar
repositórios privados passando a URL do repositório. Por exemplo o repositório
https://github.com/LucasBelusso1/GoExpert (caso fosse privado):

```SHELL
export GOPRIVATE=github.com/LucasBelusso1/GoExpert
```

E caso precisarmos utilizar mais de um repositório privado, podemos concatenar estes repositórios a variável GOPRIVATE
separando-os por vírgula:

```SHELL
export GOPRIVATE=github.com/LucasBelusso1/GoExpert,github.com/LucasBelusso1/MultithreadingChallange
```

**OBS.:** Neste cenário, apenas será possível utilizar repositórios privados referentes a sua conta.

### SSH

Caso não estejamos autenticados, precisaremos realizar a autenticação do github, e para fazer isso, editamos o arquivo
`.netrc` localizado na home do usuário do OS (`~/`) e adicionamos o seguinte:

```
machine github.com
login {{username}}
password {{token}}
```
Sendo `{{username}}` o user do github que vamos acessar e `{{token}}` o token gerado dentro de
`Settings > Developer settings > Personal access tokens` dentro da conta do github.
Vale ressaltar que esta autenticação será feita por HTTP, para fazer uma autenticação SSH precisaremos modificar o
arquivo `.git/config` (local) ou `~/.gitconfig` (global) adicionando o seguinte:

```
[url "ssh://git@github.com/"]
        insteadOf = https://github.com/
```

Isso informará ao git para sempre que for utilizar https, trocar para ssh.'

### Dependências

O GO possui um proxy responsável por centralizar os pacotes mais utilizados/baixados https://proxy.golang.org/.
Neste caso, o proxy consegue garantir de certa forma que pelo menos os pacotes mais baixados dos respositórios estejam
disponíveis para caso algum destes pacotes se tornar privado ou então o serviço de hospedagem do repositório cair, além
de tornar mais rápido o download dos pacotes.
Para garantir mais ainda que as dependências do seu projeto estejam sempre disponíveis, é possível executar o comando:

```SHELL
go mod vendor
```

Que criará a pasta `/vendor` contendo todas as dependências do projeto.