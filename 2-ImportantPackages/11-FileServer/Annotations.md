### File server

É possível fazer um servidor arquivo utilizando mux, criando um `FileServer()` a partir do pacote `http` e passando ele
para um server mux no parâmetro da função `Handle`. A função `FileServer()` recebe um parâmetro, que neste caso passamos
um `http.Dir({caminho_da_pasta})`