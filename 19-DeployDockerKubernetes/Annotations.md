### Iniciando o projeto

Primeiramente vamos criar um servidor http extremamente simples da seguinte forma:

```GO
package main

import "net/http"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello World!"))
	})
	http.ListenAndServe(":8080", nil)
}
```

### Criando dockerfile

Agora criaremos um arquivo chamado `Dockerfile` com o seguinte conteúdo:

```dockerfile
FROM golang:latest # Versão do GO

WORKDIR /app # Diretório de trabalho

CMD ["tail", "-f", "/dev/null"] # Comando para manter o container de pé
```

### GO com docker em modo dev

Agora criaremos o arquivo `docker-compose.yaml` com o seguinte conteúdo:

```YAML
version: '3' # Versão

services: # Serviços que serão rodados
  goapp: # O nome do serviço
    build: . # Direciona o build da imagem para o dockerfile que criamos anteriormente
    ports:
      - "8080:8080" # Aponta a porta 8080 do computador para a porta 8080 do container
    volumes:
      - .:/app # Informa que irá compartilhar o volume /app com o computador.
```

### Entendendo processo de build

Agora criaremos o arquivo `Dockerfile.prod` com o seguinte conteúdo:

```dockerfile
FROM golang:latest # Versão do GO

WORKDIR /app # Diretório de trabalho

COPY . . # Copia os arquivos do diretório atual para o diretório do build

RUN GOOS=linux go build -o server . # Define o sistema operacional e informa que o binário de saída se chamará "server".
```

### Otimizando a geração do executável

`DWARF` -> Debugging With Arbitraty Record Format. Esta opção vem ativada por padrão ao buildar um projeto GO. O que
esta opção faz basicamente é inserir alguns recursos dentro do binário para que seja possível executar ferramentas de
debug e profiling. Sem esta opção o arquivo final fica mais leve, entretanto não será possível executar ferramentas de
debbuging e profiling.

Para remover as informações de debugging e profile, ao buildar o binário, podemos passar a flag `ldflags` com o valor
**"-w -s"**, ficando desta forma:

```dockerfile
RUN GOOS=linux go build -ldflags="-w -s" -o server .
```
E agora adicionamos o comando que será executado para rodar o servidor:

```dockerfile
CMD ["./server"]
```

### Gerando imagem otimizada

Primeiramente para gerar a imagem, rodamos o comando:

```SHELL
docker build -t {{nome}}/{{nomeRepositorio}}:latest -f Dockerfile.prod .
```

E em seguida pegamos a imagem gerada desta forma:

```SHELL
docker images | grep {{nomeRepositorio}}
```

Veremos que será gerada uma imagem de aproximadamente 850MB.
Para verificar se a imagem está funcionando, podemos rodar o seguinte comando:

```SHELL
docker run --rm -p 8080:8080 {{nome}}/{{nomeRepositorio}}:latest
```

Então para otimizar a imagem gerada, primeiro precisamos entender o que foi gerado na imagem anterior olhando para o
arquivo do `Dockerfile.prod`:

```Dockerfile
FROM golang:latest # Obtém o GO

WORKDIR /app # Define o diretório de trabalho

COPY . . # Copia o diretório de trabalho para o diretório atual

RUN GOOS=linux go build -ldflags="-w -s" -o server . # gera o binário de ~4MB

CMD ["./server"] # Roda o binário
```

Perceba que foi feita a obtenção do GO para gerar o binário, porém após a geração do binário o GO permaneceu dentro da
imagem, fazendo com que ela fique pesada.
Para otimizar a imagem final, reescrevemos o dockerfile para o seguinte:

```dockerfile
FROM golang:latest as builder # Cria um "passo" chamado "builder"

WORKDIR /app # Define o diretório de trabalho

COPY . . # Copia o diretório de trabalho para o diretório atual

RUN GOOS=linux go build -ldflags="-w -s" -o server . # gera o binário de ~4MB

FROM scratch # A partir de uma imagem "vazia" (mínimo possível)
COPY --from=builder /app/server . # Copia o arquivo server do passo "builder" criado anteriormente
CMD ["./server"] # Inicia o servidor
```

Neste caso, criamos um passo responsável por realizar o build e em seguida, a partir de uma imagem praticamente limpa,
geramos a nossa imagem, descartando o GO obtido para gerar o binário.

### C GO e seus impactos

Para casos de projetos que não tenha dependência de recursos em C, podemos informar para não utilizar o C GO (C com GO)
ao buildar o binário da seguinte forma:

```dockerfile
RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-w -s" -o server . # desabilita o C GO
```

### Criando cluster kubernetes com Kind

Primeiramente precisamos subir a imagem que criamos para a nossa conta do docker hub, desta forma:

```SHELL
docker push lucasbelusso1/19-deploydockerkubernetes:latest
```

Agora precisaremos instalar o [kind](https://kind.sigs.k8s.io/) e rodar o comando:

```SHELL
kind create cluste --name={{nomeDoCluster}}
```

Após instalado o kind, precisaremos instalar também o [kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl-linux/). e rodar o comando:

```SHELL
kubectl cluster-info --context kind-{{nomeDoCluster}}
```

### Criando o primeiro deployment

Agora criamos a pasta `/k8s` e criamos um arquivo `yaml` de configuração com o seguinte conteúdo:

```YAML
apiVersion: apps/v1
kind: Deployment
metadata:
  name: myapp
spec:
  selector:
    matchLabels:
      app: myapp
  template:
    metadata:
      labels:
        app: myapp
    spec:
      containers:
      - name: myapp
        image: lucasbelusso1/19-deploydockerkubernetes:latest
        resources:
          limits:
            memory: "32Mi"
            cpu: "100m"
        ports:
        - containerPort: 8080
```
Agora aplicamos nosso deployment:

```SHELL
kubectl apply -f k8s/deployment.yaml
```

Para verificar se a aplicação do deployment ocorreu com sucesso, executamos o seguinte comando:

```SHELL
kubectl get pods
```

### Criando service no k8s

É possível definir uma quantidade de réplicar da aplicação, passando dentro de `spec` a propriedade `replicas` no
arquivo `k8s/deployment.yaml`, assim ao rodar o comando `kubectl apply -f k8s/deployment.yaml`, 3 pods serão gerados.
É possível conferir os pods criados rodando o comando `kubectl get pods`.

Obs.: Ao definir uma quantidade de réplicas, sempre que uma réplica cair, o k8s tentará subir uma nova para manter
sempre 3 funcionando. É possível testar este comportamento executando o seguinte:

```SHELL
kubectl delete pod {{name_do_pod}} # deleta o pod.
kubectl get pods #lista os pods em execução.
```

Agora vamos criar o `service`, que fara o papel de load balancer da nossa aplicação.
Dentro de `/k8s` criaremos o arquivo `service.yaml` com o seguinte conteúdo:

```YAML
apiVersion: v1
kind: Service
metadata:
  name: serversvc
spec:
  type: LoadBalancer
  selector:
    app: server
  ports:
  - port: 8080
    targetPort: 8080
```

Agora para iniciar o nosso service, rodamos o seguinte comando:

```SHELL
kubectl apply -f k8s/service.yaml
```
Agora para obter os serviços em execução rodamos o seguinte comando:

```SHELL
kubectl get services
```

Agora para testar rodamos o comando:

```SHELL
kubectl port-forward svc/serversvc 8080:8080
```

E em seguida executamos um curl para `localhost:8080` que deve imprimir um "Hello World!"

### Probes

Probes são "verificações" que vão garantir que o seu container/pod já subiu, se o container que subiu está pronto para
receber requisições e para verificar se o container esta no ar.

Veja o arquivo `deployment.yaml`:

```YAML
apiVersion: apps/v1
kind: Deployment
metadata:
  name: myapp
spec:
  replicas: 3
  selector:
    matchLabels:
      app: myapp
  template:
    metadata:
      labels:
        app: myapp
    spec:
      containers:
      - name: myapp
        image: lucasbelusso1/19-deploydockerkubernetes:latest
        resources:
          limits:
            memory: "32Mi"
            cpu: "100m"

          # Startup probe
          startupProbe: # Verifica ao iniciar o container a primeira vez
            httpGet: # Realiza uma chamada http
              path: / # Para o path / (normalmente se cria uma rota /health que verifica uma séried de coisas, como banco de dados entre outras coisas).
              port: 8080 # Na porta 8080
            periodSeconds: 10 # A cada 10 segundos
            failureThreshold: 10 # Se após 10 tentativas houver falha, desiste.

          readinessProbe: # Verifica o tempo todo após a aplicação subir
            httpGet: # Realiza uma chamada http
              path: / # Para o path /
              port: 8080 # Na porta 8080
            periodSeconds: 10 # A cada 10 segundos
            failureThreshold: 2  # Se após 2 tentativas houver falha, o service parará de mandar tráfego para este pod.
            timeoutSeconds: 5 # Define um tempo de timeout para cada tentaiva

          livenessProbe: # Caso o pod esteja com problema, fará o restart dele dependendo das regras definidas abaixo.
            httpGet: # Realiza uma chamada http
              path: / # Para o path /
              port: 8080 # Na porta 8080
            periodSeconds: 10 # A cada 10 segundos
            failureThreshold: 3  # Se após 3 tentativas houver falha, o service parará de mandar tráfego para este pod.
            timeoutSeconds: 5 # Define um tempo de timeout para cada tentaiva
            successThreshold: 1 # Após um caso de sucesso, reinicia a contagem.
        ports:
        - containerPort: 8080
```