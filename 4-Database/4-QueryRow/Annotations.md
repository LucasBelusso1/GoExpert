### Trabalhando com QueryRow

Para executar uma query que busca apenas um registro, fazemos da seguinte forma:

```GO
stmt, err := db.Prepare("SELECT id, name, price FROM products WHERE id = ?")

if err != nil {
	return nil, err
}

defer stmt.Close()

var p Product
err = stmt.QueryRow(id).Scan(&p.ID, &p.Name, &p.Price)
```

No código acima, preparamos a query para buscar o produto pelo ID e para executar a query utilizamos o `QueryRow()`
passando o id e depois chamamos a função `Scan` para fazer o de-para das colunas do banco para nossa struct.
Também é possível utilizar a `QueryRowContext()` para caso a query seja muito pesada, seja possível interromper a
execução utilizando alguma regra.