### Alterando dados no banco

Para alterar dados é muito semelhante ao que foi feito para inserir, mudando somente a query:

```GO
stmt, err := db.Prepare("INSERT INTO products(id, name, price) VALUES (?, ?, ?)")
```