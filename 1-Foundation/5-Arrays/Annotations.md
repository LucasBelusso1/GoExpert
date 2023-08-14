### Percorrendo arrays

Em Go, um array é uma estrutura de valores tipada e com quantidade pré definida de elementos, desta forma:

```GO
var myArray = [3]int

myArray[0] = 10
myArray[1] = 20
myArray[2] = 30
myArray[3] = 40 // Resulta em erro pois ultrapassa o limite pré definido do array
```

Neste caso o array acima tem 3 posições, começando do zero, que podem receber apenas valores do tipo int.