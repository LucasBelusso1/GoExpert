### Type assertation

Em GO, é possível converter valores de tipos diferentes em outros tipos. Para isso é utilizada a seguinte notação após
a variável: `variable.({type})`, e esta conversão retorna 2 valores, o valor da conversão e se a conversão ocorreu com
sucesso:

```GO
variable := "Hello World!"
res, ok := variable.(int)

if (!ok) { // Verifica se houve erro na conversão.
	// Retorna erro.
}
```
