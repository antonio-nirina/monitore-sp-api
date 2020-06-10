package main

import (
	"fmt"
	"log"

	// "log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/joho/godotenv"
)

func mainaaw() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	bucket := os.Getenv("FILE_CONVERTER")
	filename := "test"
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("eu-west-2"),
	}))
	f, err := os.Create(filename)

	if err != nil {
		fmt.Println("failed to create file", filename, err)
	}
	// svc := s3.New(sess)
	// res, err := svc.ListBuckets(nil)
	downloader := s3manager.NewDownloader(sess)
	n, err := downloader.Download(f, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		//Key:    aws.String(myString),
	})
	if err != nil {
		fmt.Println("failed to download file", err)
	}
	fmt.Printf("file downloaded, %d bytes\n", n)

}
