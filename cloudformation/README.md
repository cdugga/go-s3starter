aws cloudformation create-stack --stack-name s3starter --template-body file://sample.yaml --capabilities CAPABILITY_IAM 

aws cloudformation update-stack --stack-name s3starter --template-body file://sample.yaml --capabilities CAPABILITY_IAM 

aws cloudformation delete-stack --stack-name s3starter

sudo sam init --runtime go1.x --name s3startermod

sam local start-api

curl -v -H "'X-Api-Key': 'somekey'" http://127.0.0.1:3000/hello