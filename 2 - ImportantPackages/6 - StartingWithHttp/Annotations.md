### Iniciando com HTTP

Para criar um servidor em GO, basta utilizar a função `ListenAndServe()` to pacote `net/http`, a partir dela,
especifica-se uma porta e basta rodar o script que o servidor estará rodando.

Para criar rotas, utiliza-se a função `HandleFunc()` também to pacote `net/http` passando uma função que executará
aquela tarefa como parâmetro.