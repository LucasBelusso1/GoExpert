### Introdução a banco de dados

Nesta aula iniciamos um container com mysql e criamos uma tabela chamada `products` que possui as colunas ID, name e
price e criamos um script com uma struct com os campos da tabela e uma função que cria uma nova instância da struct:

```GO
type Product struct {
	ID    string
	Name  string
	Price float64
}

func newProduct(name string, price float64) *Product {
	return &Product{
		ID:    uuid.New().String(),
		Name:  name,
		Price: price,
	}
}
```

docker-compose.yaml:

```yaml
version: '3'

services:
  mysql:
    image: mysql:5.7
    container_name: mysql
    restart: always
    platform: linux/amd64
    environment:
      MYSQL_ROOT_PASSWORD: root
      MYSQL_DATABASE: goexpert
      MYSQL_PASSWORD: root
    ports:
      - 3306:3306
```