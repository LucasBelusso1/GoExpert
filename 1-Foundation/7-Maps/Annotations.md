### Maps

Maps são estruturas chave -> valor também préviamente tipadas que pode receber mais valores assim como uma slice.
Declaração de um map:

```GO
salarys := map[string]int{"Wesley": 1000, "João": 2000, "Maria": 3000}
```

Abaixo um exemplo de como adicionar uma posição a um map e um exemplo de como deletar uma posição:

```GO
delete(salarios, "Wesley") // Deleta a posição Wesley
salarios["Wes"] = 5000 // Adiciona a posição Wes com um valor de 5000
```

É possível percorrer um map da mesma forma que se percorre uma slice:

```GO
for nome, salario := range salarios {
	fmt.Printf("O salario de %s é %d\n", nome, salario)
}
```

Para ignorar um índice ou um valor, é possível utilizar o "blank identifier" `_`, desta forma:

```GO
for _, salario := range salarios {
	fmt.Printf("O salario é %d\n", salario)
}
```