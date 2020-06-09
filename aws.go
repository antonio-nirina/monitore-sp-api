package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/s3manager"
	"github.com/joho/godotenv"
)

func mainaaw() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	/*cf := aws.Credentials{
		AccessKeyID:     os.Getenv("AWS_ACCESS_KEY_ID"),
		SecretAccessKey: os.Getenv("AWS_SECRET_ACCESS_KEY"),
	}*/
	bucket := os.Getenv("FILE_CONVERTER")
	filename := "test.docx"
	key := "test_converter.docx"
	cfg, err := external.LoadDefaultAWSConfig()

	downloader := s3manager.NewDownloader(cfg)
	// Create a file to write the S3 Object contents to.
	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}

	// Write the contents of S3 Object to the file
	n, err := downloader.Download(f, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("file downloaded, %d bytes\n", n)

}
