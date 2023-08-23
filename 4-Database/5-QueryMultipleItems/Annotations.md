### Selecionando múltiplos registros

Para selecionar múltiplos registros, o processo é praticamente o mesmo dos outros exemplos, com a exceção de que não
é necessário sanitizar a query, já que não há entrada do usuário, sendo assim é possível chamar a função `db.Query()`
passando a query como parâmetro.
Para recuperar os dados, é necessário criar um slice da struct e interar sobre `rows.Next()` e popular o slice, veja
o exemplo:

```GO
rows, err := db.Query("SELECT id, name, price FROM products")

for rows.Next() {
	var p Product
	err = rows.Scan(&p.ID, &p.Name, &p.Price)

	if err != nil {
		return nil, err
	}

	products = append(products, p)
}
```