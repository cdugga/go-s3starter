# go-s3starter 

`go-s3starter` is a go module which accelerates the process of publishing and exposing HTML artifacts via AWS S3. This is 
most useful for ad-hoc experimentation or as part of a deployment pipeline

**Description**
--

`go-s3starter` automates the following:
 
1. The automation of S3 bucket provision and subsequent upload of artifacts to that S3 bucket.
2. Updating of bucket Policy to requests matching "s3:GetObject"
3. Exposure of S3 bucket as a static website , using a supplied index.html as landing page
4. Provisioning of a CLoudFront Distribution using the newly created S3 bucket as a target origin
 

**Usage**
---

As a Go Module
Use go get to fetch module
```bash

$go get github.com/cdugga/go-s3starter

```

Fr
[P98;Lom source.
 
Run git clone git@github.com:cdugga/go-s3starter.git to view source and execute locally

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


## Authors

* **Colin Duggan** - *All dev work* - [LinkedIn](https://www.linkedin.com/in/colinduggan/?originalSubdomain=ie)


## Acknowledgments

* None yet!