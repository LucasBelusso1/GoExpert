### Slices

Os slices funcionam mais ou menos como um array que pode aumentar de tamanho se necessário.
Declarando um slice:

```GO
var mySlice = []int{10, 20, 30, 50, 60, 70, 80, 90, 100}
```

Com a slice declarada, é possível fazer uma "Slice da Slice", informando entre colchetes `[]` o range que deseja
"extrair", desta forma:

```GO
mySlice[:0] // Remove todas as posições da Slice.
mySlice[:4] // Remove da primeira posição até a posição 4, sem inclir a posição 4. Resultando em [10 20 30 50]
mySlice[2:]) // Remove as duas primeiras posições da slice.
```

Também é possível adicionar posições a slice, utilizando a função `append()`:

```GO
mySlice = append(mySlice, 110)
```

Note que neste caso, a slice foi declarada com 9 posições, e ao adicionar mais uma posição, a capacidade da slice
aumenta para 18. Isso acontece pois ao ultrapassar o tamanho da slice, o GO dobra a capacidade.