### Entendendo a primeira linha

- Todo arquivo dentro de uma mesma pasta deve declarar o mesmo "package", do contrário, ao rodar o script utilizando
`go run` será exibido um erro.
- Tudo que está dentro de uma mesmo "package" será compartilhado, mesmo que em arquivos diferentes.
- Por convenção, o nome do pacote deve ser o mesmo nome do diretório atual, com exceção do `main`, que é o ponto de
entrada do código e este deve ser nomeado como `main` e deve possui a função `main()`.