### Gerando arquivos de exemplo

Para gerar arquivos para que possamos realizar nossos testes, criamos dentro de `cmd/generator` um arquivo `main.go`
com o seguinte código:

```GO
i := 0

for i < 50000 {
	f, err := os.Create(fmt.Sprintf("./tmp/file%d.txt", i))

	if err != nil {
		panic(err)
	}

	defer f.Close()
	f.WriteString("Hello World!")
	i++
}
```

E antes de rodar criamos a pasta `tmp`.

### Configurando AWS session

Para configurar a sessão da AWS, primeiro criamos dentro de `/cmd` uma pasta chamada `/uploader` com um arquivo
`main.go`. Em seguida adicionamos o seguinte código:

```GO
package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var (
	s3Client *s3.S3
	s3Bucket string
)

func init() {
	sess, err := session.NewSession(
		&aws.Config{
			Region: aws.String("us-east-1"), // Região
			Credentials: credentials.NewStaticCredentials(
				"DJIASODJSAD", // ID obtido no painel da AWS
				"DASIOHDJSAI", // Senha obtida no painel da AWS
				"",
			),
		},
	)

	if err != nil {
		panic(err)
	}

	s3Client = s3.New(sess)
	s3Bucket = "goexpert-bucket-example" // Nome do bucket na aws
}
```

### Desenvolvendo função de upload

A função ficará desta maneira:

```GO
func uploadFile(fileName string) {
	completeFileName := fmt.Sprintf("./tmp/%s", fileName) // Completa filePath + fileName
	fmt.Printf("Uploading file %s to bucket %s\n", completeFileName, s3Bucket)
	f, err := os.Open(completeFileName) // Abre o arquivo

	if err != nil {
		fmt.Printf("Error opening file %s", fileName)
		return
	}
	defer f.Close() // Fecha o arquivo no final da execução

	_, err = s3Client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(s3Bucket),
		Key:    aws.String(fileName),
		Body:   f,
	}) // Realiza upload no s3

	if err != nil {
		fmt.Printf("Error uploading file %s\n", fileName)
		return
	}

	fmt.Printf("File %s uploaded successfully\n", fileName)
}
```

E na main fazemos da seguinte forma:

```GO
dir, err := os.Open("./tmp") // Abre o diretório

if err != nil {
	panic(err)
}

defer dir.Close() // Fecha o diretório

for {
	files, err := dir.Readdir(1) //

	if err != nil {
		if err == io.EOF {
			break
		}
		fmt.Printf("Error reading directory: %s\n", err)
		continue
	}
	uploadFile(files[0].Name())
}
```

### Realizando upload com go routines.

Neste cenário, precisamos alterar nosso código da seguinte forma:

```GO
var (
	s3Client *s3.S3
	s3Bucket string
	wg       sync.WaitGroup // Criação de um WaitGroup
)
```

Função `main()`:

```GO
		wg.Add(1) // Adiciona um "crédito" ao WaitGroup
		go uploadFile(files[0].Name())
	}
	wg.Wait() // Fecha WaitGroup
```

Função `uploadFile()`:

```GO
	defer wg.Done() // Retira um crédito do WaitGroup
```

Ao rodar este script, o que vai acontecer é que será criada uma goroutine nova para cada vez que a função `uploadFile()`
for chamada. Porém, como é uma quantidade muito grande de arquivos, muitas goroutines serão criadas e consequentemente
muito do processamento e memória do computador será exigido, fazendo com que comecem a aparecer erros.

### Limitando quantidade máxima de upload

Para contornar o problema de processamente e memória, podemos controlar a quantidade de threads usando `channels`. Desta
forma:

```GO
uploadControl := make(chan struct{}, 100) // Declaramos um canal com 100 posições (buffer de 100)

// OBS.: Definimos um canal de structs vazias pois é a unidade mais leve do GO.
```

Após declarar o canal, antes da criação de cada thread preenchemos o canal com uma posição:

```GO
uploadControl <- struct{}{}
go uploadFile(files[0].Name(), uploadControl)
```

Desta forma, quando o buffer atingir o limite de 100, não será possível criar mais threads. Somente será possível criar
mais threads quando o canal for esvaziado.
Agora dentro da função de upload, ao final de cada execução (seja return ou fim da função realmente) esvaziamos o canal.
Porém, neste cenário ainda temos um problema, que é o controle de retentativas caso haja algum erro por qualquer motivo
que seja.

### Fazendo retentativas de erro

Para resolver o problema do tratamento de erro, podemos adicionar um channel ficará responsável por receber até 10
strings, que neste caso serão os nomes dos arquivos que deram erro, e criaremos uma thread que ficará responsável por
"escutar" este canal e conforme for recebendo, criar novas threads para retentar.

Então nosso código ficará assim:

```GO
func main() {
	dir, err := os.Open("./tmp")

	if err != nil {
		panic(err)
	}

	defer dir.Close()
	uploadControl := make(chan struct{}, 100)
	errorFileChannel := make(chan string, 10) // Canal contendo o nome dos arquivos com erro

	go func() { // Thread para controle de erros
		for {
			select {
			case fileName := <-errorFileChannel:
				uploadControl <- struct{}{}
				go uploadFile(fileName, uploadControl, errorFileChannel)
			}
		}
	}()

	for {
		files, err := dir.Readdir(1)

		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Printf("Error reading directory: %s\n", err)
			continue
		}
		uploadControl <- struct{}{}
		go uploadFile(files[0].Name(), uploadControl, errorFileChannel)
	}
}

func uploadFile(fileName string, uploadControl <-chan struct{}, errorFileChannel chan<- string) {
	completeFileName := fmt.Sprintf("./tmp/%s", fileName)
	fmt.Printf("Uploading file %s to bucket %s\n", completeFileName, s3Bucket)
	f, err := os.Open(completeFileName)

	if err != nil {
		fmt.Printf("Error opening file %s", fileName)
		<-uploadControl
		errorFileChannel <- fileName // Adiciona nome do arquivo com erro no channel
		return
	}
	defer f.Close()

	_, err = s3Client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(s3Bucket),
		Key:    aws.String(fileName),
		Body:   f,
	})

	if err != nil {
		fmt.Printf("Error uploading file %s\n", fileName)
		<-uploadControl
		errorFileChannel <- fileName // Adiciona nome do arquivo com erro no channel
		return
	}

	fmt.Printf("File %s uploaded successfully\n", fileName)
	<-uploadControl
}
```