# go-s3starter 

`go-s3starter` is a go module which accelerates the process of publishing and exposing HTML artifacts via AWS S3. This is 
most useful for ad-hoc experimentation or as part of a deployment pipeline

**Description**
--

`go-s3starter` automates the following:
 
1. The automation of S3 bucket provision and subsequent upload of artifacts to that S3 bucket.
2. Exposure of S3 bucket as a static website , using a supplied index.html as landing page
3. Provisioning of a CLoudFront Distribution using the newly created S3 bucket as a target origin
 

**Usage**
---

From the command line
```bash
#Build
$go build -o go-s3starter .

#Run
$./go-s3starter ./index.html
```

**Environment/Configuration**
--
1. Local setup
    + export AWS_REGION=eu-west-1 AWS_SDK_LOAD_CONFIG=true
    + export AWS_SDK_LOAD_CONFIG=true
    + export GOOGLE_APPLICATION_CREDENTIALS="/root/gcloud/creds.json"

## Authors

* **Colin Duggan** - *All dev work* - [LinkedIn](https://www.linkedin.com/in/colinduggan/?originalSubdomain=ie)


## Acknowledgments

* None yet!