### Trabalhando com JSON

Em GO, fazemos o mapeamento de JSON's para `Structs`, utilizando o pacote `json`. Ou seja, para transformar um JSON
em um "objeto" no GO, precisamos ter uma `struct` com as propriedades daquele json, somente após ter a struct,
executamos a função `Unmarshal()` do pacote `json` passando uma struct vazia como parâmetro para que as informações
sejam prenchidas.

Já para transformar a nossa `struct` em um JSON, é necessário criar um objeto daquela struct e então chamar a função
`Marshal()` que retornará um binário que precisa ser convertido para string para obter o JSON.

Também dentro do pacote `json` é possível definir encoders que enviarão o JSON para algum output definido.