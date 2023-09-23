### Race condition

Veja o código:

```GO
var number uint64 = 0

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		number++
		time.Sleep(300 * time.Millisecond)
		w.Write([]byte(fmt.Sprintf("Você é o visitante %d", number)))
	})
	http.ListenAndServe(":8000", nil)
}
```

Neste caso iniciamos um servidor que conta a quantidade de visitantes. Para testá-lo utilizamos o `Apache Bench` da
seguinte forma: `ab -n 10000 -c 100 http://localhost:8000/` que solicita 10000 requisições para nosso servidor em lotes
de 100 "pessoas" ao mesmo tempo. Ao final na execução do teste deveríamos ter o `number` igual a 10000, entretanto
o número que retorna é sempre variado e sempre inferior a 10000, isso porque por vezes, os usuários que acessam
simultâneamente acessam a mesma variável e alteram-na ao mesmo tempo, fazendo com que a soma fique incorreta.