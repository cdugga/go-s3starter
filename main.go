package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"log"
	"os"
)

func createSession() *session.Session{
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	return sess
}

func uploadFile(){

	filename := os.Args[1]
	file, err := os.Open(filename)

	if err != nil {
		fmt.Printf("Unable to open file %q, %v", err)
	}
	defer file.Close()

	sess := createSession()
	uploader := s3manager.NewUploader(sess)

	upParams := &s3manager.UploadInput{
		Bucket: aws.String("cloudstarter.org"),
		Key:    aws.String(filename),
		Body:   file,
	}

	uploader.Upload(upParams)

	if err != nil {
		// Print the error and exit.
		fmt.Printf("Unable to upload %q to %q, %v", filename, err)
	}

	fmt.Printf("Successfully uploaded ", filename)
}

func createBucket(){
	sess := createSession()
	svc := s3.New(sess)

	input := &s3.CreateBucketInput{
		Bucket: aws.String("cloudstarter.org"),
	}

	result, err := svc.CreateBucket(input)
	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			//Get error details
			log.Println("Error:", awsErr.Code(), awsErr.Message())

			// Prints out full error message, including original error if there was one.
			log.Println("Error:", awsErr.Error())

			// Get original error
			if origErr := awsErr.OrigErr(); origErr != nil {
				// operate on original error.
			}
		} else {
			fmt.Println(err.Error())
		}
	}

	fmt.Println(result)
}

func main() {
	// 1. create
	createBucket()
	// 2. upload file
	uploadFile()
	// 3. list objects in bucket
}

