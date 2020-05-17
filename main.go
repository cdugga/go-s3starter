package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudfront"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"log"
	"math/rand"
	"os"
)

func createSession() *session.Session{
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	return sess
}

func createPolicy(b string) ([]byte) {

	publicPolicy := map[string]interface{}{
		"Version": "2012-10-17",
		"Statement": []map[string]interface{}{
			{
				"Sid":       "AddPerm",
				"Effect":    "Allow",
				"Principal": "*",
				"Action": []string{
					"s3:GetObject",
				},
				"Resource": []string{
					fmt.Sprintf("arn:aws:s3:::%s/*", b),
				},
			},
		},
	}

	// Marshal the policy into a JSON value so that it can be sent to S3.
	policy, err := json.Marshal(publicPolicy)
	if err != nil {
		fmt.Printf("Failed to marshal policy, %v", err)
	}

	return policy

}

func setBbucketPolicy(b string){
	sess := createSession()
	svc := s3.New(sess)

	policy := createPolicy(b)

	params := &s3.PutBucketPolicyInput{
		Bucket:                        aws.String(b),
		ConfirmRemoveSelfBucketAccess: aws.Bool(false),
		Policy:                        aws.String(string(policy)),
	}

	// Call S3 to put the policy for the bucket.
	_, err := svc.PutBucketPolicy(params)

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok && aerr.Code() == s3.ErrCodeNoSuchBucket {
			// Special error handling for the when the bucket doesn't
			// exists so we can give a more direct error message from the CLI.
			fmt.Printf("Bucket %q does not exist", b)
		}
		fmt.Printf(("Unable to set bucket %q policy, %v", b, err)
	}
}

func createCloudFrontDistribution(url string, ref string) {
	sess := createSession()

	svc := cloudfront.New(sess)

	params := &cloudfront.CreateDistributionInput{DistributionConfig: &cloudfront.DistributionConfig{
		//Aliases: &cloudfront.Aliases{
		//	Items:    []*string{aws.String("cloudstarter.example.com"),aws.String("*cloudstarter.example.com"), },
		//	Quantity: aws.Int64(2),
		//},
		Origins: &cloudfront.Origins{
			Items: []*cloudfront.Origin{
				{

					DomainName: aws.String(url),
					Id:         aws.String("s3-cloudstarter"),
					//OriginPath: aws.String("cloudstarter.example.com"),
					CustomOriginConfig: &cloudfront.CustomOriginConfig{
						HTTPPort:               aws.Int64(80),
						HTTPSPort:              aws.Int64(443),
						OriginProtocolPolicy:   aws.String("http-only"),
					},
				},
			},
			Quantity: aws.Int64(1),
		},
		//CacheBehaviors:       nil,
		Enabled: 			aws.Bool(true), // Required
		DefaultRootObject: 	aws.String("index.html"),
		CallerReference: aws.String(ref),
		Comment:           	aws.String("cloudstarter.org generated CloudFront Distribution"),
		//CustomErrorResponses: nil,
		DefaultCacheBehavior: &cloudfront.DefaultCacheBehavior{
			ForwardedValues: &cloudfront.ForwardedValues{
				Cookies: &cloudfront.CookiePreference{
					Forward: aws.String("all"),
				},
				QueryString: aws.Bool(false),
			},
			ViewerProtocolPolicy: aws.String("redirect-to-https"),
			MinTTL:               aws.Int64(42),
			TargetOriginId:       aws.String("s3-cloudstarter"),
			TrustedSigners: &cloudfront.TrustedSigners{
				Enabled:  aws.Bool(false),
				Quantity: aws.Int64(0),
			},
		},
		//HttpVersion:          nil,
		//IsIPV6Enabled:        nil,
		Logging:           &cloudfront.LoggingConfig{
			Bucket:         aws.String(""),
			Enabled:        aws.Bool(false),
			IncludeCookies: aws.Bool(false),
			Prefix:         aws.String(""),
		},
		//OriginGroups:         nil,
		PriceClass: 		aws.String("PriceClass_All"),
		//Restrictions:         nil,
		ViewerCertificate:    &cloudfront.ViewerCertificate{
			CloudFrontDefaultCertificate: aws.Bool(true),

		},
		//WebACLId:             nil,
	}}

	resp, err := svc.CreateDistribution(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())

		return
	}
	fmt.Println(resp)
}

func enableStaticHosting(b string){

	sess := createSession()
	svc := s3.New(sess)

	params := &s3.PutBucketWebsiteInput{
		Bucket: aws.String(b),
		WebsiteConfiguration: &s3.WebsiteConfiguration{
			IndexDocument: &s3.IndexDocument{
				Suffix: aws.String("index.html"),
			},
		},
	}

	c, err := svc.PutBucketWebsite(params)
	if err != nil {
		fmt.Printf("Unable to set bucket %q website configuration, %v",
			b, err)
	}

	fmt.Print(c)

	fmt.Printf("Successfully set bucket %q website configuration\n", b)
}

func listObjects(b string){
	sess := createSession()
	svc := s3.New(sess)

	input := &s3.ListObjectsInput{
		Bucket: aws.String(b),
		EncodingType: aws.String("url"),
	}

	objects, err := svc.ListObjects(input)
	if  err != nil{
		fmt.Printf("Failed to fetch bucket objects")
	}

	for _, r := range objects.Contents {
		fmt.Printf("Object name: %s \n", *r.Key)
		fmt.Printf("Object last modified: %s \n", r.LastModified)
	}

}

func uploadFile(b string){

	filename := os.Args[1]
	file, err := os.Open(filename)

	if err != nil {
		fmt.Printf("Unable to open file %q, %v", err)
	}
	defer file.Close()

	sess := createSession()
	uploader := s3manager.NewUploader(sess)

	upParams := &s3manager.UploadInput{
		Bucket: aws.String(b),
		Key:    aws.String(filename),
		Body:   file,
		ContentType: aws.String("text/html"),
	}

	uploader.Upload(upParams)

	if err != nil {
		// Print the error and exit.
		fmt.Printf("Unable to upload %q , %s", filename, err)
	}

	fmt.Printf("Successfully uploaded file: %s \n", filename)
}

func createBucket(b string){
	sess := createSession()
	svc := s3.New(sess)

	input := &s3.CreateBucketInput{
		Bucket: aws.String(b),
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

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func main() {

	var bucket = "cloudstarter.org"
	var bucketUrl = "cloudstarter.org.s3-website-eu-west-1.amazonaws.com"
	ref := randSeq(10)


	// 1. create
	createBucket(bucket)
	// 2. upload file
	uploadFile(bucket)
	// 3. list objects in bucket
	listObjects(bucket)
	// 4 enable static site hosting for s3 bucket
	enableStaticHosting(bucket)
	// 5. Enable permissions on bucket
	setBbucketPolicy(bucket)
	// 6. create CloudFront distribution
	createCloudFrontDistribution(bucketUrl, ref)

}

