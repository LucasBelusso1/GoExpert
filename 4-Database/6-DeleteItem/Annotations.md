### Selecionando múltiplos registros

Para deletar é semelhante ao de realizar o update, sanitizando a entrada e executando a query:

```GO
stmt, err := db.Prepare("DELETE FROM products WHERE id = ?")
if err != nil {
	return err
}
defer stmt.Close()
_, err = stmt.Exec(id)
if err != nil {
	return err
}
return nil
```