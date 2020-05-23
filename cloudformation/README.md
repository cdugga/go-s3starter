aws cloudformation create-stack --stack-name s3starter --template-body file://sample.yaml --capabilities CAPABILITY_IAM 

aws cloudformation update-stack --stack-name s3starter --template-body file://sample.yaml --capabilities CAPABILITY_IAM 

aws cloudformation delete-stack --stack-name s3starter