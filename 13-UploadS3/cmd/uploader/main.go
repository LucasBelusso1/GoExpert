package main

import (
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var (
	s3Client *s3.S3
	s3Bucket string
	wg       sync.WaitGroup
)

func init() {
	sess, err := session.NewSession(
		&aws.Config{
			Region: aws.String("us-east-1"),
			Credentials: credentials.NewStaticCredentials(
				"", // Chave de acesso da AWS
				"", // Chave de acesso secreta da AWS
				"",
			),
		},
	)

	if err != nil {
		panic(err)
	}

	s3Client = s3.New(sess)
	s3Bucket = "goexpert-bucket-example-lucasbelusso1"
}

func main() {
	dir, err := os.Open("./tmp")

	if err != nil {
		panic(err)
	}

	defer dir.Close()
	uploadControl := make(chan struct{}, 100)
	errorFileChannel := make(chan string, 10)

	go func() {
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
		errorFileChannel <- fileName
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
		errorFileChannel <- fileName
		return
	}

	fmt.Printf("File %s uploaded successfully\n", fileName)
	<-uploadControl
}
