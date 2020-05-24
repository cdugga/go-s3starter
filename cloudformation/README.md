aws cloudformation create-stack --stack-name s3starter --template-body file://sample.yaml --capabilities CAPABILITY_IAM 

aws cloudformation update-stack --stack-name s3starter --template-body file://sample.yaml --capabilities CAPABILITY_IAM 

aws cloudformation delete-stack --stack-name s3starter

# SAM CLI create simple lambda and api endpoint

sudo sam init --runtime go1.x --name s3startermod

# test locally with SAM
sam local start-api

# call endpoint with AWS api key
curl -v -H "'X-Api-Key': 'somekey'" http://127.0.0.1:3000/hello


# create S3 bucket
aws s3 mb s3://s3starter-lambda

# package 
sam package --output-template-file packaged.yaml  --s3-bucket s3starter-lambda

# deploy 
sam deploy --template-file C:\Users\dugga\cdugga_code_resources\projects_go\s3starter\sam\s3startermod\packaged.yaml --stack-name s3starter-lambda --capabilities CAPABILITY_IAM 

# get endpoint
aws cloudformation describe-stacks --stack-name s3starter-lambda  --query 'Stacks[].Outputs'


# build go module ( using lambda utility for go)
set GO111MODULE=on
go.exe get -u github.com/aws/aws-lambda-go/cmd/build-lambda-zip

~\Go\Bin\build-lambda-zip.exe -o main.zip .\s3starter\s3starter