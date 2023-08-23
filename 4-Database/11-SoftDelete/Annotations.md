### Trabalhando com soft delete

Ao declararmos a propriedade `gorm.Model` dentro da nossa struct de products, o gorm automaticamente cria as colunas
`created_at`, `updated_at` e `deleted_at`, fazendo assim com que tenhamos mais controle quanto as operações em cima
deste registro. Além disso, ao declarar o `deleted_at` e executar um `Delete` em cima do registro, o registro não será
realmente deletado do banco, o que acontecerá nesse caso é que a coluna de `deleted_at` será preenchida com um datetime,
sendo considerado deletado, porém ainda persistindo no banco de dados, isso é chamado de `soft delete`.

```GO
func main() {
	// Create
	db.Create(&Product{
		Name:  "Notebook",
		Price: 1000.0,
	})

	// Update
	var p Product
	db.First(&p, 1)
	p.Name = "New Mouse"
	db.Save(p)

	// Delete
	var p2 Product
	db.First(&p2, 1)
	fmt.Println(p2.Name)
	db.Delete(p2)
}
```