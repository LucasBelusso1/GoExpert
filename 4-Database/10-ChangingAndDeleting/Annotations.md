### Alterando e removendo registros

Para alterar um registro, basta primeiramente obte-lo do banco de dados, alterar diretamente a struct retornada
e após isso chamar a função `Save` do nosso db.

```GO
	var p Product
	db.First(&p, 1)
	p.Name = "New Mouse"
	db.Save(p)
```
E para deletar basta chamar a função `Delete` do nosso db passando a struct que foi retornada do banco.

```GO
	db.Delete(p)
```