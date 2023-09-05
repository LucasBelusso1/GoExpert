### Introduzinho go mod

Em versões mais antigas, para criar um projeto em GO, era necessário criar a pasta do nosso projeto dentro do GOPATH
(constante do GO, que pode ser vista através do comando `go env`), e nas versões mais atuais é possível criar projetos
que não estejam dentro do GOPATH, a partir do sistema de módulo do GO.

Para iniciar um módulo em go, é necessário utilizar o comando `go mod {{nome}}`, e por conveniência, passamos a URL do
repositório como nome do módulo. Ex.: `go mod init github.com.br/LucasBelusso1/GoExpert`.
O módulo no GO se refere a um projeto, é este arquivo que vai gerenciar todos os pacotes necessários para rodar a sua
aplicação.