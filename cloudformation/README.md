aws cloudformation create-stack --stack-name s3starter --template-body file://sample.yaml

aws cloudformation update-stack --stack-name s3starter --template-body file://sample.yaml

aws cloudformation delete-stack --stack-name s3starter