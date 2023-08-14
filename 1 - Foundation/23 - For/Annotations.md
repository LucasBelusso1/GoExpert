### for

Em, GO não existe while, foreach, do while... Existe apenas o `for`. Exemplos de for:

Convencional
```GO
	for i := 0; i < 10; i++ {
		println(i)
	}
```

Looping infinito
```GO
	for {
		println("Looping infinito!")
	}
```

Looping com condição de parada
```GO
	i := 10
	for i < 100{
		println("Looping infinito!")
		i += 10
	}
```

Percorrendo slices, maps, arrays...
```GO
	sliceOfInt := []int{1, 2, 3, 4, 5, 6}

	for _, value := range sliceOfInt {
		fmt.Println(value)
	}
}
```
