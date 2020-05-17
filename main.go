package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudfront"
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

func createCloudFrontDistribution() {
	sess := createSession()

	svc := cloudfront.New(sess)

	params := &cloudfront.CreateDistributionInput{DistributionConfig: &cloudfront.DistributionConfig{
		Aliases: &cloudfront.Aliases{
			Items:    []*string{aws.String("cloudstarter.org.s3-website-eu-west-1.amazonaws.com")},
			Quantity: aws.Int64(1),
		},
		//CacheBehaviors:       nil,
		CallerReference: aws.String("cloudstarter.org-callerReference"),
		Comment:           	aws.String("cloudstarter.org generated CloudFront Distribution"),
		//CustomErrorResponses: nil,
		DefaultCacheBehavior: &cloudfront.DefaultCacheBehavior{
			ForwardedValues: &cloudfront.ForwardedValues{
				//Cookies: &cloudfront.CookiePreference{
				//	Forward: aws.String("whitelist"),
				//},
				QueryString: aws.Bool(false),
			},
			ViewerProtocolPolicy: aws.String("redirect-to-https"),
			MinTTL:               aws.Int64(42),
			TargetOriginId:       aws.String("origin_1"),
			TrustedSigners: &cloudfront.TrustedSigners{
				Enabled:  aws.Bool(false),
				Quantity: aws.Int64(0),
			},
		},
		DefaultRootObject: 	aws.String("index.html"),
		Enabled: 			aws.Bool(true), // Required
		//HttpVersion:          nil,
		//IsIPV6Enabled:        nil,
		//Logging:              nil,
		//OriginGroups:         nil,
		Origins: &cloudfront.Origins{
			Items: []*cloudfront.Origin{
				{
					DomainName: aws.String("cloudstarter.org.s3-website-eu-west-1.amazonaws.com"),
					Id:         aws.String("origin_1"),
					OriginPath: aws.String("/"),
				},
			},
			Quantity: aws.Int64(1),
		},
		PriceClass: 		aws.String("PriceClass_All"),
		//Restrictions:         nil,
		//ViewerCertificate:    nil,
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

func main() {

	var bucket = "cloudstarter.org"

	// 1. create
	createBucket(bucket)
	// 2. upload file
	uploadFile(bucket)
	// 3. list objects in bucket
	listObjects(bucket)
	// 4 enable static site hosting for s3 bucket

	// 5. create CloudFront distribution
	//createCloudFrontDistribution()

}

