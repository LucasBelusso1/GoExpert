### Google Cloud

1. Acessar o Google Cloud e criar uma conta para utilizar o free trial.

2. Instalar o Google Cloud CLI (gcloud). [Veja mais](https://cloud.google.com/sdk/docs/install).

3. Criar um Dockerfile com o seguinte conteúdo:

```DOCKERFILE
FROM golang:1.22 as build
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o cloudrun

FROM scratch
WORKDIR /app
COPY --from=build /app/cloudrun .
ENTRYPOINT ["./cloudrun"]
```

4. Criar um arquivo `main.go` com o seguinte conteúdo:

```GO
package main

import "net/http"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<h1>Hello, World!</h1>"))
	})

	http.ListenAndServe(":8080", nil)
}
```

5. Instalar a extensão do **[Google Cloud Code](https://marketplace.visualstudio.com/items?itemName=GoogleCloudTools.cloudcode)**.

6. Realizar deploy da aplicação.