### Server mux

Serve mux é uma forma de instanciar servidores com mais flexibilidade, podendo passar structs inteiras para atuar
como handlers e também atribuir rotas a diferentes portas.
É possível criar um mux através da função `NewServeMux()` do pacote `http`, passando-o como parâmetro para a função
`ListenAndServe()` do pacote `http`.
