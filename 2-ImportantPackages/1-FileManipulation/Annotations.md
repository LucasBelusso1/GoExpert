### Manipulação de arquivos

Em Go é possível fazer algumas operações com arquivos utilizando o pacote `os`. Nele é possível chamar a função
`Create()` para criar um arquivo (que retorna uma instância no `os.File`), `ReadFile()` ou `Open` para ler o conteúdo de
um arquivo e `Remove()` para deletar o arquivo, entre outras funções...

Com a instância do `os.File` retornada ao criar ou ler um arquivo, também é possível executar algumas funções para
manipular o arquivo em questão, como `os.File.Write()` que recebe um slice de bytes para inserir no arquivo,
`os.File.WriteString()` para inserir somente strings, entre outros recursos.

Também é possível ler o arquivo gradualmente (caso ele seja grande demais e aloque muito espaço na memória), utilizando
o pacote `bufio`, criando um novo Reader com `bufio.NewReader()` passando o arquivo em questão e passando buffer que
deverá ser respeitado ao ler o arquivo a função `Read()`.